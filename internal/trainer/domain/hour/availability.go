package hour

import (
	"errors"

	pkg_errors "github.com/pkg/errors"
)

// Availability is enum.
//
// Using struct instead of `type Availability string` for enums allows us to ensure,
// that we have full control of what values are possible.
// With `type Availability string` you are able to create `Availability("i_can_put_anything_here")`
//
// TODO 但同时也有哨兵值问题，可选值虽然在编码期预定了，但值是 var 不是 const ，可能在运行时被改变
//
//	用函数供值并配合 IsXXX 函数的方式可以解决，但代码会又长又没营养。
type Availability struct {
	a string
}

// Every type in Go have zero value. In that case it's `Availability{}`.
// It's always a good idea to check if provided value is not zero!

func (h Availability) IsZero() bool {
	return h == Availability{}
}

var (
	Available         = Availability{"available"}
	NotAvailable      = Availability{"not_available"}
	TrainingScheduled = Availability{"training_scheduled"}
)

// TODO 错误哨兵值 -> 不透明错误
//
// 原版代码使用 pkg/errors 的 New 函数创建值。这里换成标准包的errors。
// 因为 pkg/errors New 带堆栈。只有业务代码里的堆栈信息才有用，这里加上堆栈的话实际只能是日志噪音。

var (
	ErrTrainingScheduled   = errors.New("unable to modify hour, because scheduled training")
	ErrNoTrainingScheduled = errors.New("training is not scheduled")
	ErrHourNotAvailable    = errors.New("hour is not available")
)

func (h Hour) IsAvailable() bool {
	return h.availability == Available
}

func (h Hour) HasTrainingScheduled() bool {
	return h.availability == TrainingScheduled
}

func (h *Hour) MakeNotAvailable() error {
	if h.HasTrainingScheduled() {
		return pkg_errors.WithStack(ErrTrainingScheduled)
	}

	h.availability = NotAvailable
	return nil
}

func (h *Hour) MakeAvailable() error {
	if h.HasTrainingScheduled() {
		return pkg_errors.WithStack(ErrTrainingScheduled)
	}

	h.availability = Available
	return nil
}

func (h *Hour) ScheduleTraining() error {
	if !h.IsAvailable() {
		return pkg_errors.WithStack(ErrHourNotAvailable)
	}

	h.availability = TrainingScheduled
	return nil
}

func (h *Hour) CancelTraining() error {
	if !h.HasTrainingScheduled() {
		return pkg_errors.WithStack(ErrNoTrainingScheduled)
	}

	h.availability = Available
	return nil
}
