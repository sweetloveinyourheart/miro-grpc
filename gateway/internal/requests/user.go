package requests

type RegisterRequestData struct {
	FirstName string `json:"first_name" validate:"required,min=3,max=256"`
	LastName  string `json:"last_name" validate:"required,min=3,max=256"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6,max=123"`
}
