package shared

import (
	"bytes"
	_ "embed"
	"github.com/spf13/viper"
)

var KeyLocation = viper.GetString("KEY_LOCATION")

func init() {
	viper.SetConfigType("env")
	_ = viper.ReadConfig(bytes.NewBuffer(configFileContent))

}

//go:embed .env
var configFileContent []byte
