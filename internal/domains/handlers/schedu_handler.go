package handlers

import (
	"net/http"
	"time"

	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/core/utils"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/services"
	"github.com/labstack/echo/v4"
)

type ScheduHandler struct {
	scheduServices services.ScheduService
}

func NewScheduHandler(scheduServices services.ScheduService) *ScheduHandler {
	return &ScheduHandler{scheduServices: scheduServices}
}

func (sh *ScheduHandler) List(c echo.Context) error {
	shcedu, err := sh.scheduServices.List()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get schedu list")
	}

	return c.JSON(http.StatusOK, shcedu)
}

func (sh *ScheduHandler) CreateNoti(c echo.Context) error {
	var schedu models.Schedu

	err := c.Bind(&schedu)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to decode schedu data")
	}

	scheduTime, err := time.Parse("2006-01-02 15:04:05", schedu.ScheduTime)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid schedu_time format")
	}

	err = utils.ScheduleTaskDiscord(scheduTime, schedu.Content)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to schedule the task")
	}

	err = sh.scheduServices.CreateNoti(&schedu)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create schedu")
	}

	return c.NoContent(http.StatusCreated)
}
