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
	return c.String(http.StatusOK, tc.ts.CreateTeam())
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
