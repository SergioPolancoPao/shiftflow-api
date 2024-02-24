package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TeammateController struct {
	tms *TeammateService
}

func (tmc *TeammateController) CreateTeammate(c echo.Context) error {
	var body CreateTeammateRequestBody

	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(body); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	teammate := Teammate{
		Name:  body.Name,
		Email: body.Email,
	}

	if res := tmc.tms.CreateTeammate(&teammate); res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, teammate)
}

func (tmc *TeammateController) GetTeammates(c echo.Context) error {
	var qp GetTeammatesQueryParams

	if err := echo.QueryParamsBinder(c).
		String("name", &qp.Name).
		String("email", &qp.Email).
		Uint("id", &qp.ID).BindError(); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(qp); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	var teammates []Teammate
	err := tmc.tms.GetTeammates(qp, &teammates)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, teammates)
}

func (tmc *TeammateController) GetTeammate(c echo.Context) error {
	id := c.Param("id")

	var teammate Teammate
	err := tmc.tms.GetTeammate(id, &teammate)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, teammate)
}

func (tmc *TeammateController) DeleteTeammate(c echo.Context) error {
	id := c.Param("id")
	var teammate Teammate

	err := tmc.tms.DeleteTeam(id, &teammate)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, teammate)
}

func NewTeammateController(tms *TeammateService) *TeammateController {
	return &TeammateController{
		tms: tms,
	}
}
