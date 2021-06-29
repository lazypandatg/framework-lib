package DataCenterModel

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	DataSourceLib "framework-lib/src/Lib/DataSource"
	MessageLib "framework-lib/src/Lib/Message"
	"framework-lib/src/Manage"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

type Collect = collect
type collect struct {
	Start     string
	Min       int
	QueuePath string
	TableName string
	Title     string
	Content   string
	Tag       string
	Confine   string
	AllUrl    bool
}

func (_this *collect) Add(list []DataSourceLib.FieldModel) int64 {
	list = append(list, Manage.DataBase.Table(_this.TableName, ""))
	id, err := Manage.DataBase.Insert(list)
	if err != nil {
		log.Println(err)
		return -1
	}
	return id
}
func (_this *collect) JobItem(v List) {
	myUrl, _ := url.Parse(v.Url)
	file := "d://data_center/" + _this.TableName + "/" + myUrl.Host + "/"
	for i := 0; i+4 < len(myUrl.Path); i = i + 4 {
		file = file + myUrl.Path[i:i+4]
	}
	file = file + ".json"
	_, err := os.Stat(file)

	if err == nil {
		log.Println("已经存在", v)
		return
	}

	status := true
	content := Content{}
	c := colly.NewCollector()
	c.OnHTML("a", func(element *colly.HTMLElement) {
		if strings.Index(element.Attr("href"), _this.Confine) == -1 {
			return
		}

		itemUrl := element.Attr("href")
		if strings.Index(itemUrl, "?") != -1 && !_this.AllUrl {
			itemUrl = itemUrl[:strings.Index(itemUrl, "?")]
		}

		Manage.Service.Push("add", []DataSourceLib.FieldModel{
			Manage.DataBase.Table(_this.TableName, ""),
			Manage.DataBase.Set("url", itemUrl),
			Manage.DataBase.Set("mark", fmt.Sprintf("%x", md5.Sum([]byte(itemUrl)))),
		})
	})
	c.OnHTML(_this.Title, func(element *colly.HTMLElement) {
		html, err := element.DOM.Html()
		if err != nil {
			status = false
			return
		}
		content.Title = html
	})

	c.OnHTML(_this.Content, func(element *colly.HTMLElement) {
		html, err := element.DOM.Html()
		if err != nil {
			status = false
			return
		}
		content.Content = html
	})

	c.OnHTML(_this.Tag, func(element *colly.HTMLElement) {
		html, err := element.DOM.Html()
		if err != nil {
			status = false
			return
		}
		content.Tag = html
	})
	err = c.Visit(v.Url)
	if err != nil {
		log.Println(err)
	}
	str, err := json.Marshal(content)
	if err != nil {
		log.Println(err)
		return
	}
	err = os.MkdirAll(file[:strings.LastIndex(file, "/")], os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}
	err = ioutil.WriteFile(file, str, 777)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("成功：", v)
}
func (_this *collect) Listen() {
	var data string
	err := Manage.DataBase.Base.GetOneColumn("show tables like '"+_this.TableName+"'", []interface{}{}, &data)
	log.Println(data)
	if err != nil {
		return
	}
	if data == "" {
		_, err := Manage.DataBase.Base.Connection.Exec("CREATE TABLE `" + _this.TableName + "`  (\n  `id` int(11) NOT NULL AUTO_INCREMENT,\n  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,\n  `mark` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,\n  PRIMARY KEY (`id`) USING BTREE,\n  UNIQUE INDEX `mark`(`mark`) USING BTREE\n) CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;")
		if err != nil {
			return
		}
	}
	for true {
		if len(_this.Job().Data) == 0 {
			time.Sleep(5 * time.Second)
		}
	}
}
func (_this *collect) Job() MessageLib.QueueItem {
	var contentList []*List
	err := Manage.DataBase.Select([]DataSourceLib.FieldModel{
		Manage.DataBase.Table(_this.TableName, ""),
		Manage.DataBase.Greater("Id", _this.Min),
		Manage.DataBase.Page(1, 100),
	}, &contentList, List{})
	if len(contentList) == 0 && _this.Min == 0 && _this.Start != "" {
		_this.JobItem(List{Url: _this.Start})
		return MessageLib.NewQueueItem(len(contentList) > 0, contentList)
	}
	if err != nil {
		return MessageLib.NewQueueItem(false, contentList)
	}

	for _, v := range contentList {
		if v.Id > _this.Min {
			_this.Min = v.Id
		}
		_this.JobItem(List{Id: v.Id, Name: v.Name, Url: v.Url, AddTime: v.AddTime, UpdateTime: v.UpdateTime})
	}
	return MessageLib.NewQueueItem(len(contentList) > 0, contentList)
}
