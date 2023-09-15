package handlers

import (
	"net/http"

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
