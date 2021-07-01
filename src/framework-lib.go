package main

import (
	_ "framework-lib/src/Application/DataCenter/Control"
	"framework-lib/src/Lib/HttpService"
	"framework-lib/src/Manage"
	_ "framework-lib/src/Manage"
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
