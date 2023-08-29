package query

import (
	"context"
	"time"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainer/domain/hour"
)

type HourAvailabilityHandler struct {
	hourRepo hour.Repository
}

func NewHourAvailabilityHandler(hourRepo hour.Repository) HourAvailabilityHandler {
	if hourRepo == nil {
		panic("nil hourRepo") // TODO 开除预警
	}

	return HourAvailabilityHandler{hourRepo: hourRepo}
}

func (h HourAvailabilityHandler) Handle(ctx context.Context, time time.Time) (bool, error) {
	domainHour, err := h.hourRepo.GetHour(ctx, time)
	if err != nil {
		return false, err
	}

	return domainHour.IsAvailable(), nil
}
