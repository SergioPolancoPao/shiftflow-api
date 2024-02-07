package main

type CreateTeamRequestBody struct {
	Name string `json:"name" validate:"required"`
}
