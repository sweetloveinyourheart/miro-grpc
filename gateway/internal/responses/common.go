package responses

type AppResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
