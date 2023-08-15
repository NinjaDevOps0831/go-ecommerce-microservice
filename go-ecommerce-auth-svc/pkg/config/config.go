package config

import "github.com/spf13/viper"

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
			return
		}
	}
	err = viper.Unmarshal(&config)

	return
}

func GetConfig() Config {
	return config
}

// to get the secret code for jwt
func GetJWTConfig() string {
	return config.JWT
}
