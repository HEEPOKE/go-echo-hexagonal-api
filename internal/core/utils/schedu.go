package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/procyon-projects/chrono"
)

func ScheduleTaskDiscord(scheduTime time.Time, content string) error {
	taskScheduler := chrono.NewDefaultTaskScheduler()

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return err
	}

	scheduledTime := time.Date(
		scheduTime.Year(),
		scheduTime.Month(),
		scheduTime.Day(),
		scheduTime.Hour(),
		scheduTime.Minute(),
		0,
		0,
		loc,
	)

	_, err = taskScheduler.Schedule(func(ctx context.Context) {
		fmt.Printf("TimeNow:: %s\n", time.Now())
	}, chrono.WithTime(scheduledTime))

	if err != nil {
		return err
	}

	return nil
}
