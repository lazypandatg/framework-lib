package DataCenterModel

type List struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Url        string `db:"url"`
	AddTime    int    `db:"addTime"`
	UpdateTime int    `db:"updateTime"`
}
