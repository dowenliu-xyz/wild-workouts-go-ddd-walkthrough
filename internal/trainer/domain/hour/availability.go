package hour

import (
	"errors"

	pkg_errors "github.com/pkg/errors"
)

var (
	available         = Availability{"available"}
	notAvailable      = Availability{"not_available"}
	trainingScheduled = Availability{"training_scheduled"}
)

func Available() Availability {
	return available
}

func IsAvailable(availability Availability) bool {
	return availability == available
}

func NotAvailable() Availability {
	return notAvailable
}

func IsNotAvailable(availability Availability) bool {
	return availability == notAvailable
}

func TrainingScheduled() Availability {
	return trainingScheduled
}

func IsTrainingScheduled(availability Availability) bool {
	return availability == trainingScheduled
}

var availabilityValues = []Availability{
	Available(),
	NotAvailable(),
	TrainingScheduled(),
}

// Availability is enum.
//
// Using struct instead of `type Availability string` for enums allows us to ensure,
// that we have full control of what values are possible.
// With `type Availability string` you are able to create `Availability("i_can_put_anything_here")`
type Availability struct {
	a string
}

func NewAvailabilityFromString(availabilityStr string) (Availability, error) {
	for _, availability := range availabilityValues {
		if availability.String() == availabilityStr {
			return availability, nil
		}
	}
	return Availability{}, pkg_errors.Errorf("unknown '%s' availability", availabilityStr)
}

// Every type in Go have zero value. In that case it's `Availability{}`.
// It's always a good idea to check if provided value is not zero!

func (h Availability) IsZero() bool {
	return h == Availability{}
}

func (h Availability) String() string {
	return h.a
}

var (
	errTrainingScheduled   = errors.New("unable to modify hour, because scheduled training")
	errNoTrainingScheduled = errors.New("training is not scheduled")
	errHourNotAvailable    = errors.New("hour is not available")
)

func NewErrTrainingScheduled() error {
	return pkg_errors.WithStack(errTrainingScheduled)
}

func IsErrTrainingScheduled(err error) bool {
	return errors.Is(err, errTrainingScheduled)
}

func NewErrNoTrainingScheduled() error {
	return pkg_errors.WithStack(errNoTrainingScheduled)
}

func IsErrNoTrainingScheduled(err error) bool {
	return errors.Is(err, errNoTrainingScheduled)
}

func NewErrHourNotAvailable() error {
	return pkg_errors.WithStack(errHourNotAvailable)
}

func IsErrHourNotAvailable(err error) bool {
	return errors.Is(err, errHourNotAvailable)
}

func (h Hour) Availability() Availability {
	return h.availability
}

func (h Hour) IsAvailable() bool {
	return IsAvailable(h.availability)
}

func (h Hour) HasTrainingScheduled() bool {
	return IsTrainingScheduled(h.availability)
}

func (h *Hour) MakeNotAvailable() error {
	if h.HasTrainingScheduled() {
		return NewErrTrainingScheduled()
	}

	h.availability = NotAvailable()
	return nil
}

func (h *Hour) MakeAvailable() error {
	if h.HasTrainingScheduled() {
		return NewErrTrainingScheduled()
	}

	h.availability = Available()
	return nil
}

func (h *Hour) ScheduleTraining() error {
	if !h.IsAvailable() {
		return NewErrHourNotAvailable()
	}

	h.availability = TrainingScheduled()
	return nil
}

func (h *Hour) CancelTraining() error {
	if !h.HasTrainingScheduled() {
		return NewErrNoTrainingScheduled()
	}

	h.availability = Available()
	return nil
}
