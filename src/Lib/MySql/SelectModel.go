package MySqlLib

import (
	DataSource2 "github.com/lazypandatg/framework-lib/src/Lib/DataSource"
	"strings"
)

type SearchModel struct {
	FieldList []DataSource2.FieldModel
	ValueList []interface{}
	FieldBaseModel
}

func (_this *SearchModel) Sql() (string, []interface{}) {

	field, _, value := _this.GetField(_this.FieldList)
	_this.ValueList = append(_this.ValueList, value...)

	table := _this.GetTable(_this.FieldList)

	where, value := _this.GetWhere(_this.FieldList)

	group := _this.Group(_this.FieldList)

	order := _this.Order(_this.FieldList)

	page := _this.Page(_this.FieldList)
	_this.ValueList = append(_this.ValueList, value...)
	var sql string

	if len(field) == 0 {
		sql = " select * from " + table + where + group + order + page
	} else {
		sql = " select " + strings.Join(field, ",") + " form "+ table + where + group + order + page
	}

	return sql, _this.ValueList
}
