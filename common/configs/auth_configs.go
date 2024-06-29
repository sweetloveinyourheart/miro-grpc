package configs

const (
	jwtSecret string = "jwt_secret"
)

// AuthConfigs holds the Auth configuration settings.
type AuthConfigs struct {
	JwtSecret string
}

// GetAuthConfig returns the Auth configuration settings.
func GetAuthConfig() AuthConfigs {
	return AuthConfigs{
		JwtSecret: jwtSecret,
	}
}
