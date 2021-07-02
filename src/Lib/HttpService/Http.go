package HttpService

import (
	"flag"
	"github.com/lazypandatg/framework-lib/src/Lib/Message"
	"log"
	"net/http"
)

func Base() {
	host := flag.String("host", "0.0.0.0", "listen host")
	port := flag.String("port", "9701", "listen port")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.RequestURI)
		err := request.ParseForm()
		if err != nil {
			return
		}
		MessageLib.RunFromAction(request.URL.Path, request.Form, func(callResult MessageLib.QueueItem) {
			log.Println()
		})
	})

	err := http.ListenAndServe(*host+":"+*port, nil)

	if err != nil {
		panic(err)
	}
}
