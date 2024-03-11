package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	dbConfig
	apiConfig
	jwtConfig
}

type dbConfig struct {
	Driver   string `mapstructure:"DB_DRIVER"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

type apiConfig struct {
	APIPort string `mapstructure:"API_PORT"`
}

type jwtConfig struct {
	Secret    string `mapstructure:"JWT_SECRET"`
	ExpiresIn int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth *jwtauth.JWTAuth
}

func LoadConfig(path string) *conf {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("H256", []byte(cfg.Secret), nil)
	return cfg
}
