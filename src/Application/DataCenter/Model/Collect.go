package DataCenterModel

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/lazypandatg/framework-lib/src/CoreManage"
	"github.com/lazypandatg/framework-lib/src/Lib/DataSource"
	"github.com/lazypandatg/framework-lib/src/Lib/Message"
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
	list = append(list, CoreManage.DataBase.Table(_this.TableName, ""))
	id, err := CoreManage.DataBase.Insert(list)
	if err != nil {
		log.Println(err)
		return -1
	}
	return id
}
func (_this *collect) JobItem(v List) {
	myUrl, err := url.Parse(strings.Trim(v.Url, " "))
	//log.Println(err,_this,myUrl )
	if err != nil {
		return
	}
	file := "d://data_center/" + _this.TableName + "/" + myUrl.Host + "/"
	for i := 0; i+4 < len(myUrl.Path); i = i + 4 {
		file = file + myUrl.Path[i:i+4]
	}
	file = file + ".json"
	_, err = os.Stat(file)

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
		myMd5 := fmt.Sprintf("%x", md5.Sum([]byte(itemUrl)))
		md5url := "d:/data_center_md5/"
		for i := 0; i+4 < len(myMd5); i = i + 4 {
			md5url = md5url + "/" + myMd5[i:i+4]
		}
		_, err := os.Stat(md5url)
		if err == nil {
			return
		}
		err = os.MkdirAll(md5url[:strings.LastIndex(md5url, "/")], os.ModePerm)
		if err != nil {
			log.Println(err)
			return
		}

		err = ioutil.WriteFile(md5url, []byte("true"), 777)
		if err != nil {
			log.Println(err)
			return
		}
		CoreManage.Service.Push("add", []DataSourceLib.FieldModel{
			CoreManage.DataBase.Table(_this.TableName, ""),
			CoreManage.DataBase.Set("url", itemUrl),
			CoreManage.DataBase.Set("mark", fmt.Sprintf("%x", md5.Sum([]byte(itemUrl)))),
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
	err := CoreManage.DataBase.Base.GetOneColumn("show tables like '"+_this.TableName+"'", []interface{}{}, &data)
	log.Println(data)
	if err != nil {
		return
	}
	if data == "" {
		_, err := CoreManage.DataBase.Base.Connection.Exec("CREATE TABLE `" + _this.TableName + "`  (\n  `id` int(11) NOT NULL AUTO_INCREMENT,\n  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,\n  `mark` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,\n  PRIMARY KEY (`id`) USING BTREE,\n  UNIQUE INDEX `mark`(`mark`) USING BTREE\n) CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;")
		if err != nil {
			return
		}
	}
	for true {
		if len(_this.Job().Data) == 0 {
			time.Sleep(10 * time.Second)
		}
	}
}
func (_this *collect) Job() MessageLib.QueueItem {
	var contentList []*List
	err := CoreManage.DataBase.Select([]DataSourceLib.FieldModel{
		CoreManage.DataBase.Table(_this.TableName, ""),
		CoreManage.DataBase.Greater("Id", _this.Min),
		CoreManage.DataBase.Page(1, 100),
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
