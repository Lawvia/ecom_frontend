package modules

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/viper"
)

//Configuration is a class for get all configuration from config folder
type Configuration struct{}

var config *viper.Viper

//Init for initialize all configuration
func (c Configuration) Init() {
	var err error

	files, err := ioutil.ReadDir("./app/config/")
	if err != nil {
		log.Fatal(err)
	}

	var names []string

	for _, f := range files {
		fmt.Println(f.Name())
		names = append(names, strings.Split(f.Name(), ".")[0])
	}

	config = viper.New()

	config.SetConfigType("json")
	config.AddConfigPath("./app/config/")
	for i, n := range names {
		config.SetConfigName(n)
		if i == 0 {
			err = config.ReadInConfig()
		} else {
			err = config.MergeInConfig()
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	config.AutomaticEnv()
}

//GetConfig is used to get viper config and get all configuration
func (c Configuration) GetConfig() *viper.Viper {
	return config
}
