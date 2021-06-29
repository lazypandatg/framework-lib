//go:generate accessor -type=FieldModel

package DataSourceLib

type FieldModel struct {
	Name       string
	Type       int
	Value      interface{}
	Expression string
	TableName  string
}
