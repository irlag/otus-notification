package processors

import (
	"otus-notification/app/api/responses"
)

type HealthcheckProcessor interface {
	Check() (responses.Status, error)
}

type healthcheckProcessor struct{}

func NewHealtcheckProcessor() HealthcheckProcessor {
	return &healthcheckProcessor{}
}

func (h *healthcheckProcessor) Check() (responses.Status, error) {
	status := responses.StatusOK

	return status, nil
}
