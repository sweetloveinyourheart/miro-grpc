package configs

const (
	userServicePort  int32 = 41001
	boardServicePort int32 = 42001
)

// GRPCConfigs holds the gRPC configuration settings.
type GRPCConfigs struct {
	UserServicePort  int32
	BoardServicePort int32
}

// GetGRPCConfig returns the gRPC configuration settings.
func GetGRPCConfig() GRPCConfigs {
	return GRPCConfigs{
		UserServicePort:  userServicePort,
		BoardServicePort: boardServicePort,
	}
}
