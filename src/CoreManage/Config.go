package CoreManage

import (
	HttpServiceLib "github.com/lazypandatg/framework-lib/src/Lib/HttpService"
	"github.com/lazypandatg/framework-lib/src/Lib/Message"
	"github.com/lazypandatg/framework-lib/src/Lib/MySql"
)

type Config struct {
	CsdnPool     int                   `yaml:"csdn_pool"`
	ProxyDynamic string                `yaml:"proxy_dynamic"`
	DataBase     MySqlLib.Config       `yaml:"data_base"`
	DataCenter   MySqlLib.Config       `yaml:"data_center"`
	HttpService  HttpServiceLib.Config `yaml:"http_service"`
	Service      MessageLib.Config     `yaml:"service"`
	Source       MessageLib.Config     `yaml:"source"`
	Gateway      MessageLib.Config     `yaml:"gateway"`

}
