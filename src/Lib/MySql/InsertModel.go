package MySqlLib

import (
	DataSource2 "framework-lib/src/Lib/DataSource"
	"strings"
)

type InsertModel struct {
	FieldList []DataSource2.FieldModel
	ValueList []interface{}
	FieldBaseModel
}

func (_this *InsertModel) Sql() (string, []interface{}) {

	field, setList, value := _this.GetField(_this.FieldList)
	_this.ValueList = append(_this.ValueList, value...)

	table := _this.GetTable(_this.FieldList)

	where, value := _this.GetWhere(_this.FieldList)

	_this.ValueList = append(_this.ValueList, value...)
	//log.Println(where)
	if where == "" {
		sql := " insert into " + table + " (" + strings.Join(field, ",") + ") value (" + strings.Join(setList, ",") + ")"
		return sql, _this.ValueList
	} else {
		sql := " insert into " + table + " (" + strings.Join(field, ",") + ") select " + strings.Join(setList, ",") + where
		return sql, _this.ValueList
	}
}
