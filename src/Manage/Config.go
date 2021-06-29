package Manage

import (
	"framework-lib/src/Lib/Message"
	"framework-lib/src/Lib/MySql"
)

type Config struct {
	CsdnPool     int               `yaml:"csdn_pool"`
	ProxyDynamic string            `yaml:"proxy_dynamic"`
	DataBase     MySqlLib.Config   `yaml:"data_base"`
	Service      MessageLib.Config `yaml:"service"`
	Source       MessageLib.Config `yaml:"source"`
	Gateway      MessageLib.Config `yaml:"gateway"`
}
