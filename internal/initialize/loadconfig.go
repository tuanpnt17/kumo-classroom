package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tuanpnt17/kumo-classroom-be/global"
)

func LoadConfig() {
	// create viper instance
	myViper := viper.New()
	myViper.AddConfigPath("./config/")
	myViper.SetConfigName("local")
	myViper.SetConfigType("yaml")

	err := myViper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error when load config file: %s \n", err))
	}

	// read configuration and map to struct
	if err = myViper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("Fatal error when parse config file: %s \n", err))
	}
}
