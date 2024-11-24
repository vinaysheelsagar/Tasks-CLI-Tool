package configuration

import (
	"path"
	"strings"

	"github.com/spf13/viper"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

type Config struct {
	DbPath string `mapstructure:"db_path"`
}

func GetConfig() Config {
	config := Config{}

	dbName := "tasks.db"
	configFile := "config.yaml"

	taclDir := utilities.GetTaclDir()

	viper.AddConfigPath(taclDir)

	configFileDetails := strings.Split(configFile, ".")

	viper.SetConfigName(configFileDetails[0])
	viper.SetConfigType(configFileDetails[1])

	dbPath := path.Join(taclDir, dbName)

	viper.SetDefault("db_path", dbPath)

	err := viper.ReadInConfig()
	if err != nil {
		// fmt.Print(err)
		_, configFileNotFound := err.(viper.ConfigFileNotFoundError)
		if configFileNotFound {
			err = viper.WriteConfigAs(path.Join(taclDir, configFile))
			utilities.CheckNil(
				err,
				"Could not create config file",
				"Created config file.",
			)
		} else {
			panic("Some error occured while reading config.")
		}
	}

	err = viper.Unmarshal(&config)
	utilities.CheckNil(err, "", "")

	return config
}

func UpdateConfigFile() {

}
