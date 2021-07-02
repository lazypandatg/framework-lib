package main

import (
	_ "github.com/lazypandatg/framework-lib/src/Application/DataCenter/Control"
	"github.com/lazypandatg/framework-lib/src/Lib/HttpService"
	"github.com/lazypandatg/framework-lib/src/Manage"
	_ "github.com/lazypandatg/framework-lib/src/Manage"
)

func main() {
	//go func() {
	//	for true {
	//		if !DataCenterService.Csdn.Job().Status {
	//			time.Sleep(5 * time.Second)
	//		}
	//	}
	//}()

	//go DataCenterService.WangYi.Listen()
	//go DataCenterService.Jyb.Listen()
	//go DataCenterService.Ithome.Listen()

	go Manage.Service.Listener()

	go HttpService.Base()

	select {}
}
