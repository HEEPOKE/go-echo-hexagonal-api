package utils

import (
	"fmt"
	"time"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/app/hooks"
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/config"
	"github.com/go-co-op/gocron"
)

func ScheduleTaskDiscord(scheduTime time.Time, content string) error {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return err
	}

	s := gocron.NewScheduler(loc)
	_, err = s.At(scheduTime.Format("15:04")).Do(func() {
		hooks.SendDiscordMessage("HEEPOKE", content, config.Cfg.DISCORD_URL)
	})
	if err != nil {
		return err
	}
	fmt.Printf("Test:: %s\n", scheduTime)
	s.StartAsync()

	return nil
}
