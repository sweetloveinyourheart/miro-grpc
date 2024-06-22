package configs

const (
	userServicePort int32 = 41001
)

// GRPCConfigs holds the gRPC configuration settings.
type GRPCConfigs struct {
	UserServicePort int32
}

// GetGRPCConfig returns the gRPC configuration settings.
func GetGRPCConfig() GRPCConfigs {
	return GRPCConfigs{
		UserServicePort: userServicePort,
	}
}
