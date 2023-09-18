package tasks

import (
	"time"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/core/utils"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/services"
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/constants"
	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/enums"
)

type TaskSchedu struct {
	scheduServices services.ScheduService
}

func NewScheduTask(scheduServices services.ScheduService) *TaskSchedu {
	return &TaskSchedu{scheduServices: scheduServices}
}

func (sh *TaskSchedu) AutoTaskRunner() {
	scheduList, err := sh.scheduServices.List()
	if err != nil {
		utils.LoggingMessage(string(enums.DEBUG), constants.SCHEDUS_LIST_SERVICE, err.Error())
	}

	currentTime := time.Now()

	for _, schedu := range scheduList {
		scheduTime, err := time.Parse("2006-01-02 15:04:05", schedu.ScheduTime)
		if err != nil {
			utils.LoggingMessage(string(enums.DEBUG), constants.TASKAUTO_SERVICE, err.Error())
			continue
		}

		if scheduTime.After(currentTime) {
			err := utils.ScheduleTaskDiscord(scheduTime, schedu.Content)
			if err != nil {
				utils.LoggingMessage(string(enums.DEBUG), constants.TASKAUTO_SERVICE, err.Error())
				continue
			}
			utils.LoggingMessage(string(enums.INFO), constants.TASKAUTO_SERVICE, schedu.Content)
		}
	}
}
