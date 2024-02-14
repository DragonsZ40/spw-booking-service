package utils

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

var ENV *Environment
var ENV_APP string

// Read init env
func Read(path string) (*Environment, error) {
	// Get environment parameter :: dev, uat, prd

	a := os.Args
	envInput := "local"
	if len(a) > 1 {
		envInput = os.Args[1]
	}
	ENV_APP = envInput
	viper.SetConfigName(envInput)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	err := viper.Unmarshal(&ENV)
	if err != nil {
		return nil, err
	}
	return ENV, nil
}

// Read config
func ReadConfig(path string) (resp SystemConfig, err error) {
	// Get environment parameter :: dev, uat, prd

	a := os.Args
	envInput := "local"
	if len(a) > 1 {
		envInput = os.Args[1]
	}
	ENV_APP = envInput
	viper.SetConfigName(envInput)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yml")
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(&resp)
	if err != nil {
		return
	}
	return
}

func ReadConfigDatabaseList(path string) (resp SystemConfigDatabaseList, err error) {
	// Get environment parameter :: dev, uat, prd

	a := os.Args
	envInput := "local"
	if len(a) > 1 {
		envInput = os.Args[1]
	}

	viper.SetConfigName(envInput)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yml")
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(&resp)
	if err != nil {
		return
	}
	return
}
