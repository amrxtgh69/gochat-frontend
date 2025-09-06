package config

type Config struct {
	ServerAddress string
	ServerPort string
	MaxClients int
	ReadTimeout int
	WriteTimeout int
}

func DefaultConfig() *Config {
	return &Config{
		ServerAddress: "localhost",
		ServerPort: "8080",
		MaxClients: 100,
		ReadTimeout: 30,
		WriteTimeout: 30,
	}
}
func (c *Config) GetServerAddress() string {
	return c.ServerAddress + ":" + c.ServerPort
}
