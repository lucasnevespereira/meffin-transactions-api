package configs

import "github.com/spf13/viper"

type Config struct {
	Env   string
	Port  int
	DbUrl string
}

func Load() Config {
	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	conf := Config{}
	conf.Env = viper.GetString("ENV")
	conf.Port = viper.GetInt("PORT")
	conf.DbUrl = viper.GetString("DATABASE_URL")

	return conf
}
