package MessageLib

type Config struct {
	Pre      string
	Addr     string
	Port     string
	Password string
	DataBase int
	PushName string `yaml:"push_name"`
	PopName  string `yaml:"pop_name"`
}
