package DataCenterModel

type Config struct {
	Status    bool `json:"status"`
	Start     string `json:"start"`
	Min       int    `json:"Min"`
	TableName string `json:"TableName"`
	Title     string `json:"Title"`
	Content   string `json:"Content"`
	Tag       string `json:"Tag "`
	Confine   string `json:"Confine"`
	AllUrl    bool   `json:"AllUrl"`
}
