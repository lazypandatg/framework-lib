package MySqlLib

import (
	"github.com/lazypandatg/framework-lib/src/Lib/DataSource"
	"github.com/lazypandatg/framework-lib/src/Lib/MySql/Config"
	"strings"
)

type FieldBaseModel struct {
}

func (_this *FieldBaseModel) GetField(FieldList []DataSourceLib.FieldModel) ([]string, []string, []interface{}) {
	var list []string
	var setList []string
	var valueList []interface{}
	for i := 0; i < len(FieldList); i++ {
		if FieldList[i].GetType() != MySqlConfigLib.Field {
			continue
		}
		if FieldList[i].GetValue() != nil {
			valueList = append(valueList, FieldList[i].GetValue())
		}
		list = append(list, FieldList[i].GetName())
		setList = append(setList, FieldList[i].GetExpression())
	}
	return list, setList, valueList
}
func (_this *FieldBaseModel) GetUpdateField(FieldList []DataSourceLib.FieldModel) ([]string, []interface{}) {
	var list []string
	var valueList []interface{}
	for i := 0; i < len(FieldList); i++ {
		if FieldList[i].GetType() != MySqlConfigLib.Field {
			continue
		}
		if FieldList[i].GetValue() != nil {
			valueList = append(valueList, FieldList[i].GetValue())
		}
		list = append(list, FieldList[i].GetName()+" = "+FieldList[i].GetExpression())
	}
	return list, valueList
}
func (_this *FieldBaseModel) GetTable(FieldList []DataSourceLib.FieldModel) string {
	var list []string

	for i := 0; i < len(FieldList); i++ {
		if FieldList[i].GetType() != MySqlConfigLib.Table {
			continue
		}
		list = append(list, FieldList[i].GetExpression())
	}

	if len(list) > 0 {
		return strings.Join(list, " left jon ")
	} else {
		return ""
	}
}

func (_this *FieldBaseModel) GetFieldTable(FieldList []DataSourceLib.FieldModel) string {
	var list []string

	for i := 0; i < len(FieldList); i++ {
		list = append(list, FieldList[i].GetExpression())
	}

	if len(list) > 0 {
		return " from " + strings.Join(list, " left jon ")
	} else {
		return ""
	}
}

func (_this *FieldBaseModel) GetWhere(FieldList []DataSourceLib.FieldModel) (string, []interface{}) {
	var list []string
	var valueList []interface{}
	for i := 0; i < len(FieldList); i++ {
		if FieldList[i].GetType() != MySqlConfigLib.Where {
			continue
		}
		if FieldList[i].GetValue() != nil {
			valueList = append(valueList, FieldList[i].GetValue())
		}
		list = append(list, FieldList[i].GetExpression())
	}
	if len(list) > 0 {
		return " where " + strings.Join(list, " and "), valueList
	} else {
		return "", []interface{}{}
	}
}
func (_this *FieldBaseModel) Group(FieldList []DataSourceLib.FieldModel) string {
	var list []string
	for i := 0; i < len(FieldList); i++ {
		if FieldList[i].GetType() != MySqlConfigLib.Group {
			continue
		}
		list = append(list, FieldList[i].GetExpression())
	}
	if len(list) == 0 {
		return ""
	}
	return " group by " + strings.Join(list, ",")
}
func (_this *FieldBaseModel) Order(FieldList []DataSourceLib.FieldModel) string {
	var list []string
	for i := 0; i < len(FieldList); i++ {
		if FieldList[i].GetType() != MySqlConfigLib.Order {
			continue
		}
		list = append(list, FieldList[i].GetExpression())
	}
	if len(list) == 0 {
		return ""
	}
	return " order by " + strings.Join(list, ",")
}

func (_this *FieldBaseModel) Page(FieldList []DataSourceLib.FieldModel) string {
	var list []string
	for i := 0; i < len(FieldList); i++ {
		if FieldList[i].GetType() != MySqlConfigLib.Page {
			continue
		}
		list = append(list, FieldList[i].GetExpression())
	}
	return " limit " + strings.Join(list, ",")
}
