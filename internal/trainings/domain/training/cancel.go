package training

import (
	"errors"
	"time"

	pkg_errors "github.com/pkg/errors"
)

func (t Training) CanBeCanceledForFree() bool {
	return time.Until(t.time) >= time.Hour*24
}

var errTrainingAlreadyCanceled = errors.New("training is already canceled")

func NewErrTrainingAlreadyCanceled() error {
	return pkg_errors.WithStack(errTrainingAlreadyCanceled)
}

func IsErrTrainingAlreadyCanceled(err error) bool {
	return errors.Is(err, errTrainingAlreadyCanceled)
}

func (t *Training) Cancel() error {
	if t.IsCanceled() {
		return NewErrTrainingAlreadyCanceled()
	}

	t.canceled = true
	return nil
}

func (t Training) IsCanceled() bool {
	return t.canceled
}
