package HttpServiceLib

import (
	"flag"
	"github.com/lazypandatg/framework-lib/src/Lib/Message"
	"log"
	"net/http"
)

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
type HttpService struct {
	Config Config
}

func (_this *HttpService) Init(config Config) {
	_this.Config = config
	log.Println(_this.Config)
	go _this.Listen()
}

func (_this *HttpService) Listen() {
	host := flag.String("host", _this.Config.Host, "listen host")
	port := flag.String("port", _this.Config.Port, "listen port")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.RequestURI)

		err := request.ParseForm()
		if err != nil {
			return
		}

		MessageLib.RunFromAction(request.URL.Path, request.Form, func(callResult MessageLib.QueueItem) {
			_, err := writer.Write([]byte(callResult.Data))
			if err != nil {
				log.Println(err)
				return
			}
		})
	})
	log.Println(*host+":"+*port)
	err := http.ListenAndServe(*host+":"+*port, nil)

	if err != nil {
		panic(err)
	}
}
