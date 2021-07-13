package CoreManage

import (
	"github.com/lazypandatg/framework-lib/src/Lib/HttpService"
	"github.com/lazypandatg/framework-lib/src/Lib/Message"
	"github.com/lazypandatg/framework-lib/src/Lib/MySql"
	"github.com/lazypandatg/framework-lib/src/Lib/Queue"
	"github.com/lazypandatg/framework-lib/src/Lib/Util/DataType"
	"log"
)

var DataBase = &MySqlLib.DataSource{}
var Service = &MessageLib.Queue{}
var Source = &MessageLib.Queue{}
var Gateway = &MessageLib.Queue{}
var BatchInsert = &Queue.BatchInsertQueue{TableList: map[string][]MySqlLib.InsertModel{}}
var DataCenter = &MySqlLib.DataSource{}
var HttpService = &HttpServiceLib.HttpService{}

func init() {
	myConfig := Config{}
	err := DataTypeUtil.YamlFileToStruct("develop.yaml", &myConfig)
	if err != nil {
		log.Println("develop.yaml配置文件读取失败：",err)
		return
	}
	DataBase.Init(myConfig.DataBase)
	DataCenter.Init(myConfig.DataCenter)
	Source.Init(myConfig.Source)
	Service.Init(myConfig.Service)
	Gateway.Init(myConfig.Gateway)
	HttpService.Init(myConfig.HttpService)
	BatchInsert.DataBase = DataBase
}