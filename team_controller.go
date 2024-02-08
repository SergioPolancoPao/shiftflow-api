package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

func (tc *TeamController) GetTeams(c echo.Context) error {
	var qp ListTeamQueryParams

	echo.QueryParamsBinder(c).String("name", &qp.Name)

	if err := echo.QueryParamsBinder(c).String("name", &qp.Name).BindError(); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(qp); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	teams, err := tc.ts.GetTeams(qp.Name)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, teams)
}

func (tc *TeamController) GetTeam(c echo.Context) error {
	id := c.Param("id")

	team, err := tc.ts.GetTeam(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, team)
}

func (tc *TeamController) DeleteTeam(c echo.Context) error {
	id := c.Param("id")

	team, err := tc.ts.DeleteTeam(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, team)
}

func (tc *TeamController) addRoutes() {
	tg := tc.eServer.Group("/teams")
	tg.POST("", tc.CreateTeam)
	tg.GET("", tc.GetTeams)
	tg.GET("/:id", tc.GetTeam)
	tg.DELETE("/:id", tc.DeleteTeam)
}

func NewTeamController(server *echo.Echo, ts *TeamService) *TeamController {
	obj := &TeamController{
		ts:      ts,
		eServer: server,
	}
	obj.addRoutes()
	return obj
}
