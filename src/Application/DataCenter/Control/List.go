package Control

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	DataCenterModel "github.com/lazypandatg/framework-lib/src/Application/DataCenter/Model"
	"github.com/lazypandatg/framework-lib/src/Application/DataCenter/Service"
	"github.com/lazypandatg/framework-lib/src/Lib/DataSource"
	"github.com/lazypandatg/framework-lib/src/Lib/Message"
	"github.com/lazypandatg/framework-lib/src/Manage"
)

var _ = MessageLib.AddAction("add", func(fieldList []DataSourceLib.FieldModel) MessageLib.QueueItem {
	//log.Println("list入库",fieldList)
	//go Manage.DataBase.Insert(fieldList)
	Manage.BatchInsert.Add(fieldList)
	return MessageLib.NewQueueItem(true, 0)
})
var _ = MessageLib.AddAction("content/add", func(list DataCenterModel.List) MessageLib.QueueItem {
	go DataCenterService.Csdn.JobItem(list)
	return MessageLib.NewQueueItem(true, 0)
})
var _ = MessageLib.AddAction("Content.UpdateJob", func(fieldList []DataSourceLib.FieldModel) MessageLib.QueueItem {
	return DataCenterService.Csdn.Job()
})

var _ = MessageLib.AddAction("/DataCenter/AddCsdnList", func(
	request struct {
	Data string `param:"data"`
},
) MessageLib.QueueItem {
	var list articlesList
	data := request.Data[len(`<html><head></head><body><pre style="word-wrap: break-word; white-space: pre-wrap;">`) : len(request.Data)-len(`</pre></body></html>`)]
	err := json.Unmarshal([]byte(data), &list)
	if err != nil {
		return MessageLib.QueueItem{}
	}

	for _, v := range list.Articles {
		DataCenterService.Csdn.Add([]DataSourceLib.FieldModel{
			Manage.DataBase.Set("url", v.Url),
			Manage.DataBase.Set("mark", fmt.Sprintf("%x", md5.Sum([]byte(v.Url)))),
		})
	}

	return MessageLib.NewQueueItem(true, nil)
})

type articlesList struct {
	Status       string      `json:"status"`
	LastViewTime interface{} `json:"last_view_time"`
	Message      string      `json:"message"`
	ShownOffset  int         `json:"shown_offset"`
	Articles     []struct {
		Comments      string        `json:"comments"`
		Avatarurl     string        `json:"avatarurl"`
		UserName      string        `json:"user_name"`
		RecommendType string        `json:"recommend_type"`
		CreatedAt     string        `json:"created_at"`
		Focus         bool          `json:"focus"`
		Digg          string        `json:"digg"`
		Recommend     string        `json:"recommend"`
		Title         string        `json:"title"`
		Url           string        `json:"url"`
		Tags          []interface{} `json:"tags"`
		ProductType   string        `json:"product_type"`
		ReportData    struct {
			EventClick bool `json:"eventClick"`
			Data       struct {
				Mod           string `json:"mod"`
				Extra         string `json:"extra"`
				DistRequestId string `json:"dist_request_id"`
				Index         string `json:"index"`
				Strategy      string `json:"strategy"`
			} `json:"data"`
			UrlParams struct {
				UtmMedium       string `json:"utm_medium"`
				Depth1UtmSource string `json:"depth_1-utm_source"`
			} `json:"urlParams"`
			EventView bool `json:"eventView"`
		} `json:"report_data"`
		ProductId   string `json:"product_id"`
		Nickname    string `json:"nickname"`
		Style       string `json:"style"`
		Views       string `json:"views"`
		Desc        string `json:"desc"`
		Id          string `json:"id"`
		Type        string `json:"type"`
		CategoryId  string `json:"category_id"`
		Category    string `json:"category"`
		IsPlan      bool   `json:"isPlan"`
		StrategyId  string `json:"strategy_id"`
		Strategy    string `json:"strategy"`
		TaceCode    string `json:"tace_code"`
		ShownOffset int    `json:"shown_offset"`
		ShownTime   string `json:"shown_time"`
		UserUrl     string `json:"user_url"`
		Avatar      string `json:"avatar"`
	} `json:"articles"`
	NoWartchers bool `json:"no_wartchers"`
}
