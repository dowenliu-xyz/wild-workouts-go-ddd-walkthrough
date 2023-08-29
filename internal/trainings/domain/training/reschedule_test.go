package training_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainings/domain/training"
)

func TestTraining_RescheduleTraining(t *testing.T) {
	tr := newExampleTraining(t)

	oldTime := tr.Time()
	newTime := time.Now().AddDate(0, 0, 14).Round(time.Hour)

	// it's always a good idea to ensure about pre-conditions in the test ;-)
	assert.False(t, oldTime.Equal(newTime))

	err := tr.RescheduleTraining(newTime)
	assert.NoError(t, err)
	assert.True(t, tr.Time().Equal(newTime))
}

func TestTraining_RescheduleTraining_less_than_24h_before(t *testing.T) {
	originalTime := time.Now().Round(time.Hour)
	rescheduleRequestTime := originalTime.AddDate(0, 0, 5)

	tr := newExampleTrainingWithTime(t, originalTime)

	err := tr.RescheduleTraining(rescheduleRequestTime)

	assert.EqualError(t, err, training.CantRescheduleBeforeTimeError{
		TrainingTime: tr.Time(),
	}.Error())
}

func TestTraining_ProposeReschedule_by_attendee(t *testing.T) {
	testCases := []struct {
		Name     string
		Proposer training.UserType
		Approver training.UserType
	}{
		{
			Name:     "proposed_by_attendee",
			Proposer: training.UserTypeAttendee(),
			Approver: training.UserTypeTrainer(),
		},
		{
			Name:     "proposed_by_trainer",
			Proposer: training.UserTypeTrainer(),
			Approver: training.UserTypeAttendee(),
		},
	}

	for _, c := range testCases {
		t.Run(c.Name, func(t *testing.T) {
			originalTime := time.Now().Round(time.Hour)
			rescheduleRequestTime := originalTime.AddDate(0, 0, 5)
			tr := newExampleTrainingWithTime(t, originalTime)

			assert.False(t, tr.IsRescheduleProposed())

			tr.ProposeReschedule(rescheduleRequestTime, c.Proposer)

			assert.True(t, tr.IsRescheduleProposed())

			err := tr.ApproveReschedule(c.Approver)
			require.NoError(t, err)

			assert.True(t, tr.Time().Equal(rescheduleRequestTime))
			assert.False(t, tr.IsRescheduleProposed())
		})
	}
}

func TestTraining_ProposeReschedule_approve_by_proposer(t *testing.T) {
	testCases := []struct {
		Proposer training.UserType
	}{
		{
			Proposer: training.UserTypeAttendee(),
		},
		{
			Proposer: training.UserTypeTrainer(),
		},
	}

	for _, c := range testCases {
		t.Run(c.Proposer.String(), func(t *testing.T) {
			originalTime := time.Now().Round(time.Hour)
			rescheduleRequestTime := originalTime.AddDate(0, 0, 5)
			tr := newExampleTrainingWithTime(t, originalTime)

			tr.ProposeReschedule(rescheduleRequestTime, c.Proposer)

			err := tr.ApproveReschedule(c.Proposer)
			assert.Error(t, err)

			assert.True(t, tr.Time().Equal(originalTime))
			assert.True(t, tr.IsRescheduleProposed())
		})
	}
}

func TestTraining_ApproveReschedule_not_proposed(t *testing.T) {
	tr := newExampleTrainingWithTime(t, time.Now().Round(time.Hour))

	assert.EqualError(t, tr.ApproveReschedule(training.UserTypeTrainer()), training.NewErrNoRescheduleRequested().Error())
}

func TestTraining_RejectRescheduleTraining(t *testing.T) {
	originalTime := time.Now().Round(time.Hour)
	rescheduleRequestTime := originalTime.AddDate(0, 0, 5)
	tr := newExampleTrainingWithTime(t, originalTime)

	tr.ProposeReschedule(rescheduleRequestTime, training.UserTypeAttendee())

	err := tr.RejectReschedule()
	assert.NoError(t, err)

	assert.True(t, tr.Time().Equal(originalTime))
	assert.False(t, tr.IsRescheduleProposed())
}
