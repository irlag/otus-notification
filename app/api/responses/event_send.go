package responses

import "net/http"

//go:generate easyjson

//easyjson:json
type EventSendOkResponse struct {
	Success bool `json:"success"`
}

func NewEventSendOkResponse() EventSendOkResponse {
	return EventSendOkResponse{
		Success: true,
	}
}

func (r *EventSendOkResponse) WriteResponse(rw http.ResponseWriter) {
	payload, _ := r.MarshalJSON()

	WriteJsonResponse(rw, http.StatusOK, payload)
}
