package module

type Config struct {
	Port int
}

func (c Config) init() {
	c.Port = 5000
}

func (c Config) SetConfig(port int) Config {
	c.Port = port

	return c
}

func (c Config) LoadConfig() Config {
	return c
}
