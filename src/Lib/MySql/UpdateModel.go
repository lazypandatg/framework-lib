package MySqlLib

import (
	DataSource2 "github.com/lazypandatg/framework-lib/src/Lib/DataSource"
	"strings"
)

type UpdateModel struct {
	FieldList []DataSource2.FieldModel
	ValueList []interface{}
	FieldBaseModel
}

func (_this *UpdateModel) Sql() (string, []interface{}) {
	field, value := _this.GetUpdateField(_this.FieldList)
	_this.ValueList = append(_this.ValueList, value...)

	table := _this.GetTable(_this.FieldList)

	where, value := _this.GetWhere(_this.FieldList)
	_this.ValueList = append(_this.ValueList, value...)

	sql := " update " + table + " set " + strings.Join(field, ",") + where
	return sql, _this.ValueList
}
