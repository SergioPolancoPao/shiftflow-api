package main

type CreateTeamRequestBody struct {
	Name string `json:"name" validate:"required"`
}

type GetTeamsQueryParams struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
