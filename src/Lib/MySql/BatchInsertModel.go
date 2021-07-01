package MySqlLib

import "strings"

type BatchInsertModel struct {
	InsertList []InsertModel
	ValueList  []interface{}
	FieldBaseModel
}

type BatchInsertItemModel struct {
	TableName string
	sql       []string
	Field     string
	Param     []interface{}
}

func (_this *BatchInsertItemModel) SetTableName(string2 string) {
	_this.TableName = string2
}
func (_this *BatchInsertItemModel) AddSql(string2 string) {
	_this.sql = append(_this.sql, string2)
}

func (_this *BatchInsertItemModel) AddParam(paramList []interface{}) {
	_this.Param = append(_this.Param, paramList...)
}

func (_this *BatchInsertItemModel) SetField(string2 string) {
	_this.Field = string2
}
func (_this *BatchInsertItemModel) Sql() (string, []interface{}) {
	return "insert into " + _this.TableName + " (" + _this.Field + ") values " + strings.Join(_this.sql, ","), _this.Param
}

func (_this *BatchInsertModel) Add(item InsertModel) *BatchInsertModel {
	_this.InsertList = append(_this.InsertList, item)
	return _this
}

func (_this *BatchInsertModel) Sql() map[string]*BatchInsertItemModel {
	var itemMap = map[string]*BatchInsertItemModel{}
	for i := 0; i < len(_this.InsertList); i++ {
		table, field, sql, param := _this.InsertList[i].GetBatchSql()
		item, ok := itemMap[table]
		if !ok {
			item = &BatchInsertItemModel{}
			item.SetField(field)
			item.TableName = table
			itemMap[table] = item
		}
		item.AddSql(sql)
		item.AddParam(param)
	}
	return itemMap
}
