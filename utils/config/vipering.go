package config

import "github.com/spf13/viper"

type Config struct {
	BackendHost string 	`mapstructure:"BACKEND_HOST"`
	BackendPort	string	`mapstructure:"BACKEND_PORT"`
	BackendScheme string `mapstructure:"BACKEND_SCHEME"`
	JwtSignUpEmailSecretKey string `mapstructure:"JWT_SIGN_UP_EMAIL_SECRET_KEY"`
	JwtSignUpEmailAudience string `mapstructure:"JWT_SIGN_UP_EMAIL_AUDIENCE"`
	SmtpServerHost string `mapstructure:"SMTP_SERVER_HOST"`
	SmtpServerPort string `mapstructure:"SMTP_SERVER_PORT"`
	SmtpServerUsername string `mapstructure:"SMTP_SERVER_USERNAME"`
	SmtpServerPassword string `mapstructure:"SMTP_SERVER_PASSWORD"`
	DbSource	string	`mapstructure:"DB_SOURCE"`
	DbName	string	`mapstructure:"DB_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("auth")
	// viper will use existing env variable
	//viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
