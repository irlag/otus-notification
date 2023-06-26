package parameters

//go:generate easyjson

import (
	"io"
	"net/http"
)

//easyjson:json
type EventSendParams struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func NewEventSendParamsFromRequest(request *http.Request) (*EventSendParams, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	eventSendParams := &EventSendParams{}
	err = eventSendParams.UnmarshalJSON(body)
	if err != nil {
		return nil, err
	}

	return eventSendParams, nil
}
