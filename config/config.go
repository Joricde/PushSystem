package config

type config struct {
	AppConfig struct {
		Release bool `yaml:"release"`
		Port    int  `yaml:"port"`
	} `yaml:"appConfig"`

	Mysql struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		DB       string `yaml:"db"`
	} `yaml:"mysql"`
	Redis struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
	JwtSecret string `yaml:"jwtSecret"`
	LogPath   string `yaml:"logPath"`
}
