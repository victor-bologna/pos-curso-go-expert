package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *Conf

type Conf struct {
	DbConfig  `mapstructure:"DB"`
	ApiConfig `mapstructure:"API"`
	JwtConfig `mapstructure:"JWT"`
}

type DbConfig struct {
	Driver   string `mapstructure:"DRIVER"`
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Name     string `mapstructure:"NAME"`
}

type ApiConfig struct {
	APIPort string `mapstructure:"PORT"`
}

type JwtConfig struct {
	Secret    string `mapstructure:"SECRET"`
	ExpiresIn int    `mapstructure:"EXPIRESIN"`
	JWTAuth   *jwtauth.JWTAuth
}

func LoadConfig(path string) *Conf {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	cfg.JwtConfig.JWTAuth = jwtauth.New("HS256", []byte(cfg.JwtConfig.Secret), nil)
	return cfg
}
