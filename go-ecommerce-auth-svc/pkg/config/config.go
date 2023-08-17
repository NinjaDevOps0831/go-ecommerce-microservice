package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
	//DBUrl string `mapstructure:"DB_URL"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	JWT string `mapstructure:"JWT_CODE"`

	AUTHTOKEN  string `mapstructure:"TWILIO_AUTH_TOKEN"`
	ACCOUNTSID string `mapstructure:"TWILIO_ACCOUNT_SID"`
	SERVICESID string `mapstructure:"TWILIO_SERVICE_SID"`
}

var envs = []string{"PORT", "DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD",
	"JWT_CODE",
	"TWILIO_AUTH_TOKEN", "TWILIO_ACCOUNT_SID", "TWILIO_SERVICE_SID",
}
var config Config

func LoadConfig() (config *Config, err error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	for _, env := range envs {
		if err = viper.BindEnv(env); err != nil {
			fmt.Println("debug checkout  - error in config.go")
			return config, err
		}
	}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("debug checkout  - error 2 in config.go")
		return config, err
	}
	fmt.Println("Config is", config, "twilio authtoken is", config.AUTHTOKEN, "twilio acnt sid is", config.ACCOUNTSID, "twilio service sid is", config.SERVICESID)

	return config, nil
}

func GetConfig() Config {
	return config
}

// to get the secret code for jwt
func GetJWTConfig() string {
	return config.JWT
}
