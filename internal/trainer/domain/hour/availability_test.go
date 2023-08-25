package hour_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainer/domain/hour"
)

func TestHour_MakeNotAvailable(t *testing.T) {
	h, err := hour.NewAvailableHour(validTrainingHour())
	require.NoError(t, err)

	require.NoError(t, h.MakeNotAvailable())
	assert.False(t, h.IsAvailable())
}

func TestHour_MakeNotAvailable_with_scheduled_training(t *testing.T) {
	h := newHourWithScheduledTraining(t)

	assert.ErrorIs(t, h.MakeNotAvailable(), hour.ErrTrainingScheduled)
}

func TestHour_MakeAvailable(t *testing.T) {
	h, err := hour.NewAvailableHour(validTrainingHour())
	require.NoError(t, err)

	require.NoError(t, h.MakeNotAvailable())

	require.NoError(t, h.MakeAvailable())
	assert.True(t, h.IsAvailable())
}

func TestHour_MakeAvailable_with_scheduled_training(t *testing.T) {
	h := newHourWithScheduledTraining(t)

	assert.ErrorIs(t, h.MakeAvailable(), hour.ErrTrainingScheduled)
}

func TestHour_ScheduleTraining(t *testing.T) {
	h, err := hour.NewAvailableHour(validTrainingHour())
	require.NoError(t, err)

	require.NoError(t, h.ScheduleTraining())

	assert.True(t, h.HasTrainingScheduled())
	assert.False(t, h.IsAvailable())
}

func TestHour_ScheduleTraining_with_not_available(t *testing.T) {
	h := newNotAvailableHour(t)
	assert.ErrorIs(t, h.ScheduleTraining(), hour.ErrHourNotAvailable)
}

func TestHour_CancelTraining(t *testing.T) {
	h := newHourWithScheduledTraining(t)

	require.NoError(t, h.CancelTraining())

	assert.False(t, h.HasTrainingScheduled())
	assert.True(t, h.IsAvailable())
}

func TestHour_CancelTraining_no_training_scheduled(t *testing.T) {
	h, err := hour.NewAvailableHour(validTrainingHour())
	require.NoError(t, err)

	assert.ErrorIs(t, h.CancelTraining(), hour.ErrNoTrainingScheduled)
}

func newHourWithScheduledTraining(t *testing.T) *hour.Hour {
	h, err := hour.NewAvailableHour(validTrainingHour())
	require.NoError(t, err)

	require.NoError(t, h.ScheduleTraining())

	return h
}

func newNotAvailableHour(t *testing.T) *hour.Hour {
	h, err := hour.NewAvailableHour(validTrainingHour())
	require.NoError(t, err)

	require.NoError(t, h.MakeNotAvailable())

	return h
}
