package event

import (
	"github.com/gorilla/mux"

	"otus-notification/app/api"
	"otus-notification/app/processors"
)

type Event struct {
	processors *processors.Processors
}

func NewEventApi(processors *processors.Processors) *Event {
	return &Event{
		processors: processors,
	}
}

func (r *Event) HandleMethods(router *mux.Router) {
	router.HandleFunc(api.AppRoutes["event_send"].Path, r.Send()).
		Methods(api.AppRoutes["event_send"].Method).
		Name(api.AppRoutes["event_send"].Name)
}
