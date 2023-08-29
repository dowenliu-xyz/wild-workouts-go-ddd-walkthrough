package training

import (
	"errors"
	"fmt"
	"time"

	pkg_errors "github.com/pkg/errors"
)

func (t Training) MovedProposedBy() UserType {
	return t.moveProposedBy
}

func (t Training) ProposedNewTime() time.Time {
	return t.proposedNewTime
}

type CantRescheduleBeforeTimeError struct {
	TrainingTime time.Time
}

func (c CantRescheduleBeforeTimeError) Error() string {
	return fmt.Sprintf(
		"can't reschedule training, not enough time before, training time: %s",
		c.TrainingTime,
	)
}

func (t *Training) RescheduleTraining(newTime time.Time) error {
	if !t.CanBeCanceledForFree() {
		err := CantRescheduleBeforeTimeError{
			TrainingTime: t.Time(),
		}
		return pkg_errors.WithStack(err)
	}

	t.time = newTime

	return nil
}

func (t *Training) ProposeReschedule(newTime time.Time, proposerType UserType) {
	// TODO 感觉这里少了些检查。比如当前是否培训是否已过时、是否已取消
	t.moveProposedBy = proposerType
	t.proposedNewTime = newTime
}

func (t *Training) IsRescheduleProposed() bool {
	return !t.moveProposedBy.IsZero() && !t.proposedNewTime.IsZero()
}

var errNoRescheduleRequested = errors.New("no training reschedule was requested yet")

func NewErrNoRescheduleRequested() error {
	return pkg_errors.WithStack(errNoRescheduleRequested)
}

func IsErrNoRescheduleRequested(err error) bool {
	return errors.Is(err, errNoRescheduleRequested)
}

func (t *Training) ApproveReschedule(userType UserType) error {
	if !t.IsRescheduleProposed() {
		return NewErrNoRescheduleRequested()
	}

	if t.moveProposedBy == userType {
		return pkg_errors.Errorf(
			"trying to approve reschedule by the same user type which proposed reschedule (%s)",
			userType.String(),
		)
	}

	t.time = t.proposedNewTime

	t.proposedNewTime = time.Time{}
	t.moveProposedBy = UserType{}

	return nil
}

func (t *Training) RejectReschedule() error {
	if !t.IsRescheduleProposed() {
		return NewErrNoRescheduleRequested()
	}

	t.proposedNewTime = time.Time{}
	t.moveProposedBy = UserType{}

	return nil
}
