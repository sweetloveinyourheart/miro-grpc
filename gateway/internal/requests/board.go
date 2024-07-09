package requests

type CreateBoardRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
