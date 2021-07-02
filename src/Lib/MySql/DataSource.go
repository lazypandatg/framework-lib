//go:generate accessor -type=DataSource
package MySqlLib

import (
	"github.com/lazypandatg/framework-lib/src/Lib/DataSource"
	"github.com/lazypandatg/framework-lib/src/Lib/MySql/Config"
	"log"
	"strconv"
)

type DataSource struct {
	Base *BaseModel
}

func NewDataSource(config Config) *DataSource {
	d := &DataSource{}
	d.Init(config)
	return d
}

func (_this *DataSource) Init(config Config) {
	_this.Base = NewBase(config)
	err := _this.Base.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(_this.Base.Connection.Ping())
}

func (_this *DataSource) Select(paramList []DataSourceLib.FieldModel, Data interface{}, DataType interface{}) error {
	mySearch := SearchModel{FieldList: paramList}
	sql, Parameter := mySearch.Sql()
	return _this.Base.Select(sql, Parameter, Data, DataType)
}

func (_this *DataSource) Insert(paramList []DataSourceLib.FieldModel) (int64, error) {
	mySearch := InsertModel{FieldList: paramList}
	sql, Parameter := mySearch.Sql()
	return _this.Base.Insert(sql, Parameter)
}

func (_this *DataSource) Update(paramList []DataSourceLib.FieldModel) (int64, error) {
	mySearch := UpdateModel{FieldList: paramList}
	sql, Parameter := mySearch.Sql()
	return _this.Base.Update(sql, Parameter)
}

func (_this *DataSource) Connect() error {
	return _this.Base.Connect()
}

func (_this *DataSource) Table(table string, join string) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: table, Type: MySqlConfigLib.Table, Value: "", Expression: table + join}
}

func (_this *DataSource) TableExpression(table string, join string) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Type: MySqlConfigLib.Table, Value: "", Expression: table + join}
}

func (_this *DataSource) Set(name string, value interface{}) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: name, Type: MySqlConfigLib.Field, Value: value, Expression: " ? "}
}

func (_this *DataSource) Show(name string) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: name, Type: MySqlConfigLib.Field, Value: "", Expression: name}
}

func (_this *DataSource) ShowExpression(Expression string) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Type: MySqlConfigLib.Field, Value: "", Expression: Expression}
}

func (_this *DataSource) Like(name string, value string) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: name, Type: MySqlConfigLib.Where, Value: value, Expression: name + " like CONCAT('&',?,'%') "}
}

func (_this *DataSource) Equal(name string, value string) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: name, Type: MySqlConfigLib.Where, Value: value, Expression: name + " = ? "}
}

func (_this *DataSource) EqualGreater(name string, value int) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: name, Type: MySqlConfigLib.Where, Value: value, Expression: name + " >= ?"}
}

func (_this *DataSource) EqualLess(name string, value int) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: name, Type: MySqlConfigLib.Where, Value: value, Expression: name + " <= ? "}
}

func (_this *DataSource) Expression(name string, value int) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: name, Type: MySqlConfigLib.Where, Value: value, Expression: name + " < ?"}
}

func (_this *DataSource) Greater(name string, value int) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: name, Type: MySqlConfigLib.Where, Value: value, Expression: name + " > ? "}
}

func (_this *DataSource) Less(name string, value int) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: name, Type: MySqlConfigLib.Where, Value: value, Expression: name + " < ? "}
}

func (_this *DataSource) Group(name string) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: name, Type: MySqlConfigLib.Group, Value: "", Expression: name}
}

func (_this *DataSource) OrderAsc(name string) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: name, Type: MySqlConfigLib.Order, Value: "", Expression: name + " asc "}
}

func (_this *DataSource) OrderDesc(name string) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Name: name, Type: MySqlConfigLib.Order, Value: "", Expression: name + " desc "}
}

func (_this *DataSource) Page(page int, count int) DataSourceLib.FieldModel {
	return DataSourceLib.FieldModel{Type: MySqlConfigLib.Page, Value: "", Expression: strconv.Itoa((page-1)*count) + "," + strconv.Itoa(count)}
}
