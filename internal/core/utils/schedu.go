package utils

import (
	"time"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/app/hooks"
)

func ScheduleTaskDiscord(scheduTime time.Time, content string) error {
	_, err := gocron.Every(1).Day().At(scheduTime.Format("15:04")).Do(hooks.SendDiscordMessage(), content)
	if err != nil {
		return err
	}

	return nil
}
