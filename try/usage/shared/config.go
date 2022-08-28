package shared

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("env")
	_ = viper.ReadConfig(bytes.NewBuffer(configFileContent))

	fmt.Println(viper.GetString("TEST"))
}

//go:embed .env
var configFileContent []byte
