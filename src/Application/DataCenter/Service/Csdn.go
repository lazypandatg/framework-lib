package DataCenterService

import (
	"github.com/lazypandatg/framework-lib/src/Application/DataCenter/Model"
)

var Csdn = DataCenterModel.Collect{QueuePath: "163/list/add", TableName: "a163_list"}


//var Csdn = csdn{}
//
//type csdn struct {
//	Min int
//}
//
//func (_this *csdn) Add(list []DataSourceLib.FieldModel) {
//	list = append(list, Manage.DataBase.Table("csdn_list", ""))
//	id, err := Manage.DataBase.Insert(list)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	log.Println("csdn_list增加" + strconv.FormatInt(id, 10))
//}
//
//func (_this *csdn) Update(list []DataSourceLib.FieldModel) {
//
//}
//
//func (_this *csdn) List(list []DataSourceLib.FieldModel) {
//
//}
//
//func (_this *csdn) JobItem(v DataCenterModel.List) {
//	myUrl, _ := url.Parse(v.Url)
//	file := "d://data_center/" + myUrl.Host + "/"
//	for i := 0; i+4 < len(myUrl.Path); i = i + 4 {
//		file = file + myUrl.Path[i:i+4]
//	}
//	file = file + ".json"
//	_, err := os.Stat(file)
//
//	if err == nil {
//		log.Println("已经存在", v)
//		return
//	}
//
//	status := true
//	content := DataCenterModel.Content{}
//	c := colly.NewCollector()
//	c.OnHTML("a", func(element *colly.HTMLElement) {
//		if strings.Index(element.Attr("href"), "blog.csdn.net") == -1 {
//			return
//		}
//		itemUrl := element.Attr("href")
//		if strings.Index(itemUrl, "?") != -1 {
//			itemUrl = itemUrl[:strings.Index(itemUrl, "?")]
//		}
//		Manage.Service.Push("list/add", []DataSourceLib.FieldModel{
//			Manage.DataBase.Set("url", itemUrl),
//			Manage.DataBase.Set("mark", fmt.Sprintf("%x", md5.Sum([]byte(itemUrl)))),
//		})
//	})
//	c.OnHTML(".title-article", func(element *colly.HTMLElement) {
//		html, err := element.DOM.Html()
//		if err != nil {
//			status = false
//			return
//		}
//		content.Title = html
//	})
//
//	c.OnHTML("#content_views", func(element *colly.HTMLElement) {
//		html, err := element.DOM.Html()
//		if err != nil {
//			status = false
//			return
//		}
//		content.Content = html
//	})
//
//	c.OnHTML("#artic-tag-box", func(element *colly.HTMLElement) {
//		html, err := element.DOM.Html()
//		if err != nil {
//			status = false
//			return
//		}
//		content.Tag = html
//	})
//	err = c.Visit(v.Url)
//	if err != nil {
//		//log.Println(err)
//
//		log.Println(err)
//	}
//	str, err := json.Marshal(content)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	err = os.MkdirAll(file[:strings.LastIndex(file, "/")], os.ModePerm)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	err = ioutil.WriteFile(file, str, 777)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	log.Println("成功：", v)
//}
//func (_this *csdn) Job() MessageLib.QueueItem {
//	var contentList []*DataCenterModel.List
//	err := Manage.DataBase.Select([]DataSourceLib.FieldModel{
//		Manage.DataBase.Table("csdn_list", ""),
//		Manage.DataBase.Greater("Id", _this.Min),
//		Manage.DataBase.Page(1, 100),
//	}, &contentList, DataCenterModel.List{})
//
//	if err != nil {
//		return MessageLib.NewQueueItem(false, contentList)
//	}
//
//	for _, v := range contentList {
//		if v.Id > _this.Min {
//			_this.Min = v.Id
//		}
//		_this.JobItem(DataCenterModel.List{Id: v.Id, Name: v.Name, Url: v.Url, AddTime: v.AddTime, UpdateTime: v.UpdateTime})
//	}
//	return MessageLib.NewQueueItem(len(contentList) > 0, contentList)
//}
