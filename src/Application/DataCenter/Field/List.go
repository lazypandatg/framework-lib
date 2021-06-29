package DataCenterField

import MysqlLib "framework-lib/src/Lib/MySql"

var List = list{}

type list struct {
	Id         MysqlLib.FieldBaseModel
	Name       MysqlLib.FieldBaseModel
	Content    MysqlLib.FieldBaseModel
	AddTime    MysqlLib.FieldBaseModel
	UpdateTime MysqlLib.FieldBaseModel
}
