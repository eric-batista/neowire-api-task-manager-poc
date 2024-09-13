package core

type Config struct {
	BasePath string
}

func SetEnvs() Config {
	return Config{
		BasePath: GetEnv("BASE_PATH", "/neowire/tasks-manager"),
	}
}
