package training

import (
	"fmt"

	pkg_errors "github.com/pkg/errors"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/errors"
)

// UserType is enum-like type.
// We are using struct instead of string, to ensure about immutability.
type UserType struct {
	s string
}

func (u UserType) IsZero() bool {
	return u == UserType{}
}

func (u UserType) String() string {
	return u.s
}

var (
	trainer  = UserType{"trainer"}
	attendee = UserType{"attendee"}
)

func UserTypeTrainer() UserType {
	return trainer
}

func IsUserTypeTrainer(u UserType) bool {
	return u == trainer
}

func UserTypeAttendee() UserType {
	return attendee
}

func IsUserTypeAttendee(u UserType) bool {
	return u == attendee
}

func NewUserTypeFromString(userType string) (UserType, error) {
	switch userType {
	case "trainer":
		return trainer, nil
	case "attendee":
		return attendee, nil
	}

	return UserType{}, errors.NewSlugError(
		fmt.Sprintf("invalid '%s' role", userType),
		"invalid-role",
	)
}

type User struct {
	userUUID string
	userType UserType
}

func (u User) UUID() string {
	return u.userUUID
}

func (u User) Type() UserType {
	return u.userType
}

func (u User) IsEmpty() bool {
	return u == User{}
}

func NewUser(userUUID string, userType UserType) (User, error) {
	if userUUID == "" {
		return User{}, pkg_errors.New("missing user UUID")
	}
	if userType.IsZero() {
		return User{}, pkg_errors.New("missing user type")
	}

	return User{userUUID: userUUID, userType: userType}, nil
}

func MustNewUser(userUUID string, userType UserType) User {
	u, err := NewUser(userUUID, userType)
	if err != nil {
		panic(err)
	}

	return u
}

type ForbiddenToSeeTrainingError struct {
	RequestingUserUUID string
	TrainingOwnerUUID  string
}

func (f ForbiddenToSeeTrainingError) Error() string {
	return fmt.Sprintf(
		"user '%s' can't see user '%s' training",
		f.RequestingUserUUID, f.TrainingOwnerUUID,
	)
}

func CanUserSeeTraining(user User, training Training) error {
	if IsUserTypeTrainer(user.Type()) {
		return nil
	}
	if user.UUID() == training.UserUUID() {
		return nil
	}

	return pkg_errors.WithStack(ForbiddenToSeeTrainingError{user.UUID(), training.UserUUID()})
}
