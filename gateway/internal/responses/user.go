package responses

type RegisterResponseData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type SignInResponseData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ErrorResponseData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type UserProfileResponseData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
