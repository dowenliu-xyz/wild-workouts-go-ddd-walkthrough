package ports

import (
	"net/http"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/go-chi/render"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/auth"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/server/httperr"
	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/trainer/app"
)

type HttpServer struct {
	service app.HourService
}

func NewHttpServer(service app.HourService) HttpServer {
	return HttpServer{
		service: service,
	}
}

func (h HttpServer) GetTrainerAvailableHours(w http.ResponseWriter, r *http.Request, queryParams GetTrainerAvailableHoursParams) {
	if queryParams.DateFrom.After(queryParams.DateTo) {
		httperr.BadRequest("date-from-after-date-to", nil, w, r)
		return
	}

	dateModels, err := h.service.GetTrainerAvailableHours(r.Context(), queryParams.DateFrom, queryParams.DateTo)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	dates := dateModelsToResponse(dateModels)
	render.Respond(w, r, dates)
}

func dateModelsToResponse(models []app.Date) []Date {
	var dates []Date
	for _, d := range models {
		var hours []Hour
		for _, h := range d.Hours {
			hours = append(hours, Hour{
				Available:            h.Available,
				HasTrainingScheduled: h.HasTrainingScheduled,
				Hour:                 h.Hour,
			})
		}

		dates = append(dates, Date{
			Date: openapi_types.Date{
				Time: d.Date,
			},
			HasFreeHours: d.HasFreeHours,
			Hours:        hours,
		})
	}

	return dates
}

func (h HttpServer) MakeHourAvailable(w http.ResponseWriter, r *http.Request) {
	user, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	if user.Role != "trainer" {
		httperr.Unauthorised("invalid-role", nil, w, r)
		return
	}

	hourUpdate := &HourUpdate{}
	if err := render.Decode(r, hourUpdate); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	err = h.service.MakeHoursAvailable(r.Context(), hourUpdate.Hours)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h HttpServer) MakeHourUnavailable(w http.ResponseWriter, r *http.Request) {
	user, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	if user.Role != "trainer" {
		httperr.Unauthorised("invalid-role", nil, w, r)
		return
	}

	hourUpdate := &HourUpdate{}
	if err := render.Decode(r, hourUpdate); err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	err = h.service.MakeHoursUnavailable(r.Context(), hourUpdate.Hours)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
