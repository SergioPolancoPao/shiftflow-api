package main

type CreateTeammateRequestBody struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type GetTeammatesQueryParams struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
