package config

import "github.com/spf13/viper"

type Config struct {
	AuthServiceUrl    string `mapstructure:"AUTH_SVC_URL"`
	ProductServiceUrl string `mapstructure:"PRODUCT_SVC_URL"`
	CartServiceUrl    string `mapstructure:"CART_SVC_URL"`
	OrderServiceUrl   string `mapstructure:"ORDER_SVC_URL"`
	Port              string `mapstructure:"PORT"`
	JWT               string `mapstructure:"JWT_CODE"`
}

var envs = []string{"AUTH_SVC_URL",
	"PRODUCT_SVC_URL", "CART_SVC_URL", "ORDER_SVC_URL", "PORT", "JWT_CODE",
}
var config *Config

func LoadConfig() (*Config, error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return nil, err
		}
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return config, nil
}

func GetConfig() Config {
	return *config
}

// to get the secret code for jwt
// func GetJWTConfig() string {
// 	return config.JWT
// }
