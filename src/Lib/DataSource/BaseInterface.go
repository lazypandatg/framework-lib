package DataSourceLib

type BaseInterface interface {
	Select(sql string, Parameter []interface{}, Data interface{}, DataType interface{}) error
	Insert(sql string, Parameter []interface{}) (int64, error)
	Update(sql string, Parameter []interface{}) (int64, error)
	Connect() error
}
