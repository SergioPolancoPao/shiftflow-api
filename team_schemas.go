package main

type CreateTeamRequestBody struct {
	Name string `json:"name" validate:"required"`
}

type ListTeamQueryParams struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
