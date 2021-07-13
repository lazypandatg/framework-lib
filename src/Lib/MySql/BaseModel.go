//go:generate accessor -type=BaseModel,Config

package MySqlLib

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/lazypandatg/framework-lib/src/Lib/Util/DataType"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
	"reflect"
)

type BaseModel struct {
	Connection *sqlx.DB `access:"r,w"`
	Config     Config
}
type Config struct {
	Host         string `yaml:"host" access:"r,w"`
	Port         string `yaml:"port" access:"r,w"`
	DataBaseName string `yaml:"data_base_name" access:"r,w" `
	UserName     string `yaml:"user_name" access:"r,w" `
	UserPassword string `yaml:"user_password" access:"r,w"`
}

func NewBase(config Config) *BaseModel {
	base := &BaseModel{}
	base.Config = config
	err := base.Connect()
	if err != nil {
		return nil
	}
	return base
}

func (_this *BaseModel) Connect() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", _this.Config.UserName, _this.Config.UserPassword, _this.Config.Host, _this.Config.Port, _this.Config.DataBaseName, "utf8")
	logrus.Info(dsn)
	DB, err := sqlx.Open("mysql", dsn)
	if err != nil {
		logrus.Info("mysql connect failed, detail is [" + err.Error() + "]")
		return err
	} else {
		logrus.Info("mysql connect succeed")
		_this.Connection = DB
		return nil
	}
}

func (_this *BaseModel) Select(sql string, Parameter []interface{}, Data interface{}, DataType interface{}) error {
	rows, err := _this.Connection.Query(sql, Parameter...)
	defer rows.Close()
	if err != nil {
		logrus.Errorln(sql, Parameter)
		logrus.Errorln(err)
		return err
	}
	columns, _ := rows.Columns()
	cacheItemType := DataTypeUtil.GetTypeElem(reflect.TypeOf(DataType))
	setData := reflect.Indirect(reflect.ValueOf(Data))
	for rows != nil && rows.Next() {
		values := make([]interface{}, len(columns))
		cacheItem := reflect.New(cacheItemType)
		DataTypeUtil.GetField(columns, cacheItem, values, "")
		_ = rows.Scan(values...)
		setData.Set(reflect.Append(setData, cacheItem))
	}

	return nil
}

func (_this *BaseModel) Insert(sql string, Parameter []interface{}) (int64, error) {
	query, err := _this.Connection.Exec(sql, Parameter...)

	if err != nil {
		log.Println(err)
		return -1, err
	}
	id, err := query.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (_this *BaseModel) Update(sql string, Parameter []interface{}) (int64, error) {
	query, err := _this.Connection.Exec(sql, Parameter...)
	if err != nil {
		return -1, err
	}
	affected, err := query.RowsAffected()
	if err != nil {
		return -1, err
	}
	return affected, nil
}

func (_this *BaseModel) GetOne(sql string, Parameter []interface{}, Data interface{}, DataType interface{}) error {
	rows, err := _this.Connection.Query(sql, Parameter...)
	rows.Close()
	if err != nil {
		log.Println(err)
		return err
	}

	columns, _ := rows.Columns()
	cacheItemType := DataTypeUtil.GetTypeElem(reflect.TypeOf(DataType))
	setData := reflect.Indirect(reflect.ValueOf(Data))

	if rows != nil && rows.Next() {
		values := make([]interface{}, len(columns))
		cacheItem := reflect.New(cacheItemType)
		DataTypeUtil.GetField(columns, cacheItem, values, "")
		_ = rows.Scan(values...)
		setData.Set(reflect.Append(setData, cacheItem))
	}
	return nil
}
func (_this *BaseModel) GetOneColumn(sql string, Parameter []interface{}, Data interface{}) error {
	rows, err := _this.Connection.Query(sql, Parameter...)
	if err != nil {
		log.Println(err)
		return err
	}

	if rows != nil && rows.Next() {
		err := rows.Scan(Data)
		if err != nil {
			return err
		}
	}
	return nil
}
