package MySqlLib

import (
	DataSource2 "github.com/lazypandatg/framework-lib/src/Lib/DataSource"
	"strings"
)

type InsertModel struct {
	FieldList []DataSource2.FieldModel
	ValueList []interface{}
	FieldBaseModel
}

func (_this *InsertModel) GetBatchSql() (string, string, string, []interface{}) {
	field, setList, value := _this.GetField(_this.FieldList)
	ValueList := value
	table := _this.GetTable(_this.FieldList)
	sql := "(" + strings.Join(setList, ",") + ")"
	return table, strings.Join(field, ","), sql, ValueList
}

func (_this *InsertModel) Sql() (string, []interface{}) {

	field, setList, value := _this.GetField(_this.FieldList)

	ValueList := value

	table := _this.GetTable(_this.FieldList)

	where, whereValue := _this.GetWhere(_this.FieldList)

	ValueList = append(ValueList, whereValue...)

	if where == "" {
		sql := " insert into " + table + " (" + strings.Join(field, ",") + ") values (" + strings.Join(setList, ",") + ")"
		return sql, ValueList
	} else {
		sql := " insert into " + table + " (" + strings.Join(field, ",") + ") select " + strings.Join(setList, ",") + where
		return sql, ValueList
	}
}
