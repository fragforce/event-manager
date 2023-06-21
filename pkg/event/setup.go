package event

import "github.com/spf13/viper"

var (
	users  viper.Viper
	teams  viper.Viper
	events viper.Viper
)

func LoadEvents(config viper.Viper) (err error) {
	teams.AddConfigPath(config.GetString("events.config.folder"))
	teams.SetConfigFile("teams.yml")

	users.AddConfigPath(config.GetString("events.config.folder"))
	users.SetConfigFile("users.yml")

	return
}
