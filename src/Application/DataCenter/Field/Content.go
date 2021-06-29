package DataCenterField

import MysqlLib "framework-lib/src/Lib/MySql"

var Content = content{}

type content struct {
	Id         MysqlLib.FieldBaseModel
	Name       MysqlLib.FieldBaseModel
	Title      MysqlLib.FieldBaseModel
	Content    MysqlLib.FieldBaseModel
	AddTime    MysqlLib.FieldBaseModel
	UpdateTime MysqlLib.FieldBaseModel
}
