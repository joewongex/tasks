package module

import "github.com/spf13/viper"

type Config struct {
	Redis struct {
		Host string
		Port int
		DB   int
	}
	Notify struct {
		Bark struct {
			Url       string
			DeviceKey string `mapstructure:"device_key"`
		}
	}
	Log struct {
		Level  string
		File   string
		Stdout bool
	}
}

func NewConfig() (config *Config, err error) {
	viper.SetConfigFile("./config.json")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	config = new(Config)
	err = viper.Unmarshal(config)
	return
}
