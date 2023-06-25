package event

import (
	"net/http"

	"otus-notification/app/api/parameters"
	"otus-notification/app/api/responses"
)

func (r *Event) Send() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		eventSendParams, err := parameters.NewEventSendParamsFromRequest(request)
		if err != nil {
			responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(writer)

			return
		}

		err = r.processors.EventProcessor.Send(request.Context(), eventSendParams)
		if err != nil {
			responses.NewErrorResponse(http.StatusInternalServerError, err).WriteErrorResponse(writer)

			return
		}

		eventSendOkResponse := responses.NewEventSendOkResponse()
		eventSendOkResponse.WriteResponse(writer)
	}
}
