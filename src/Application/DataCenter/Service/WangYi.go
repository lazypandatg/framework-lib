package DataCenterService

import (
	"encoding/json"
	"github.com/lazypandatg/framework-lib/src/Application/DataCenter/Model"
	"io/ioutil"
)

func init() {
	file, err := ioutil.ReadFile("collect.json")
	if err != nil {
		return
	}
	var configList []DataCenterModel.Config
	err = json.Unmarshal(file, &configList)
	if err != nil {
		return
	}
	for _, v := range configList {
		if !v.Status {
			continue
		}

		item := DataCenterModel.Collect{
			Start:     v.Start,
			TableName: v.TableName,
			Confine:   v.Confine,
			Title:     v.Title,
			Content:   v.Content,
			Tag:       v.Tag,
		}
		go item.Listen()
	}
}
