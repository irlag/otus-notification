package services

import (
	"go.uber.org/zap"

	"otus-notification/app/config"
)

type Services struct {
}

func New(log *zap.Logger, config *config.Config) *Services {
	return &Services{}
}
