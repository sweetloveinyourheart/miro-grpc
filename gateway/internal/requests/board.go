package requests

type CreateBoardRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=256"`
	Description string `json:"description" validate:"required,min=3,max=256"`
}
