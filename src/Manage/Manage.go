package Manage

import (
	"fmt"
	"framework-lib/src/Lib/Message"
	"framework-lib/src/Lib/MySql"
	"framework-lib/src/Lib/Queue"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var DataBase = &MySqlLib.DataSource{}
var Service = &MessageLib.Queue{}
var Source = &MessageLib.Queue{}
var Gateway = &MessageLib.Queue{}
var BatchInsert = &Queue.BatchInsertQueue{TableList: map[string][]MySqlLib.InsertModel{}}

func init() {
	myConfig := load()
	DataBase.Init(myConfig.DataBase)
	Source.Init(myConfig.Source)
	Service.Init(myConfig.Service)
	Gateway.Init(myConfig.Gateway)
	BatchInsert.DataBase = DataBase
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
