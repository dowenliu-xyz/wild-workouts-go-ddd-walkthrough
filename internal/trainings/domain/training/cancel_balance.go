package training

import "fmt"

// CancelBalanceDelta return trainings balance delta that should be adjusted after training cancellation.
func CancelBalanceDelta(tr Training, cancelingUserType UserType) int {
	if tr.CanBeCanceledForFree() {
		// just give training back
		return 1
	}

	switch {
	case IsUserTypeTrainer(cancelingUserType):
		// 1 for cancelled training +1 "fine" for cancelling by trainer less than 24h before training
		return 2
	case IsUserTypeAttendee(cancelingUserType):
		// "fine" for cancelling less than 24h before training
		return 0
	default:
		panic(fmt.Sprintf("not supported user type %s", cancelingUserType)) // TODO 开除预警
	}
}
