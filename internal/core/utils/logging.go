package utils

import (
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/enums"
	"github.com/rs/zerolog/log"
)

func LoggingMessage(level, services, message string) {
	switch level {
	case string(enums.INFO):
		log.Info().Str("service", services).Msg(message)
	case string(enums.DEBUG):
		log.Debug().Str("service", services).Msg(message)
	default:
		log.Info().Str("service", services).Msg(message)
	}
}
