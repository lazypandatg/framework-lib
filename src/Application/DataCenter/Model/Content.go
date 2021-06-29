package DataCenterModel

type Content struct {
	Id         int    `db:"Id"`
	Title      string `db:"Name"`
	Content    string `db:"Content"`
	Tag        string `db:"Content"`
	AddTime    int    `db:"AddTime"`
	UpdateTime int    `db:"UpdateTime"`
}
