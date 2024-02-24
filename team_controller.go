package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TeamController struct {
	ts *TeamService
}

func (tc *TeamController) CreateTeam(c echo.Context) error {
	var body CreateTeamRequestBody

	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(body); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	team := Team{
		Name: body.Name,
	}

	if res := tc.ts.CreateTeam(&team); res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, team)
}

func (tc *TeamController) GetTeams(c echo.Context) error {
	var qp GetTeamsQueryParams

	if err := echo.QueryParamsBinder(c).
		String("name", &qp.Name).
		Uint("id", &qp.ID).BindError(); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(qp); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	var teams []Team
	err := tc.ts.GetTeams(qp, &teams)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, teams)
}

func (tc *TeamController) GetTeam(c echo.Context) error {
	id := c.Param("id")

	var team Team
	err := tc.ts.GetTeam(id, &team)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, team)
}

func (tc *TeamController) DeleteTeam(c echo.Context) error {
	id := c.Param("id")
	var team Team

	err := tc.ts.DeleteTeam(id, &team)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, team)
}

func NewTeamController(ts *TeamService) *TeamController {
	return &TeamController{
		ts: ts,
	}
}
