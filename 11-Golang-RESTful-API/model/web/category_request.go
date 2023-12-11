package web

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required,min=3,max=200"`
}

type CategoryUpdateRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required,min=3,max=200"`
}
