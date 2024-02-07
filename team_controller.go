package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ITeamController interface {
	CreateTeam(c echo.Context) error
}

type TeamController struct {
	eServer *echo.Echo
	ts      *TeamService
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

func (tc *TeamController) addRoutes() {
	tc.eServer.POST("/team", tc.CreateTeam)
}

func NewTeamController(server *echo.Echo, ts *TeamService) *TeamController {
	obj := &TeamController{
		ts:      ts,
		eServer: server,
	}
	obj.addRoutes()
	return obj
}
