package Manage

import (
	"fmt"
	"framework-lib/src/Lib/Message"
	"framework-lib/src/Lib/MySql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var DataBase = &MySqlLib.DataSource{}
var Service = &MessageLib.Queue{}
var Source = &MessageLib.Queue{}
var Gateway = &MessageLib.Queue{}

func init() {
	myConfig := load()
	DataBase.Init(myConfig.DataBase)
	Source.Init(myConfig.Source)
	Service.Init(myConfig.Service)
	Gateway.Init(myConfig.Gateway)
}
func load() Config {
	yamlFile, err := ioutil.ReadFile("develop.yaml")
	if err != nil {
		fmt.Printf("failed to read yaml file : %v\n", err)
		return Config{}
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("failed to unmarshal : %v\n", err)
		return Config{}
	}
	return config
}
