package Queue

import (
	"github.com/lazypandatg/framework-lib/src/Lib/DataSource"
	"github.com/lazypandatg/framework-lib/src/Lib/MySql"
	"log"
	"strconv"
	"sync"
)

type BatchInsertQueue struct {
	Batch     MySqlLib.BatchInsertModel
	TableList map[string][]MySqlLib.InsertModel
	Max       int
	a         sync.Mutex
	DataBase  *MySqlLib.DataSource
}

func (_this *BatchInsertQueue) Add(paramList []DataSourceLib.FieldModel) *BatchInsertQueue {
	_this.a.Lock()
	defer _this.a.Unlock()
	item := MySqlLib.InsertModel{FieldList: paramList}
	if _, ok := _this.TableList[item.GetTable(item.FieldList)]; !ok {
		_this.TableList[item.GetTable(item.FieldList)] = []MySqlLib.InsertModel{}
	} else {
		_this.TableList[item.GetTable(item.FieldList)] = append(_this.TableList[item.GetTable(item.FieldList)], item)
	}
	_this.Listen(item.GetTable(item.FieldList))
	return _this
}
func (_this *BatchInsertQueue) Listen(mark string) {
	//log.Println(_this.TableList[mark])
	if len(_this.TableList[mark]) > 20 {
		batchInsert := MySqlLib.BatchInsertModel{InsertList: _this.TableList[mark][0:20]}
		sqlInfoList := batchInsert.Sql()
		for _, v := range sqlInfoList {
			update, err := _this.DataBase.Base.Update(v.Sql())
			if err != nil {
				count := 0
				for _, item := range batchInsert.InsertList {
					_, err = _this.DataBase.Base.Insert(item.Sql())
					if err != nil {
						continue
					} else {
						count++
					}
				}
				log.Println("【去重入库】【" + mark + "】【" + strconv.Itoa(count) + "】")
				return
			} else {
				log.Println("【入库】【" + mark + "】【" + strconv.Itoa(int(update)) + "】")
			}
		}
		_this.TableList[mark] = _this.TableList[mark][20:]
	}
}

type BatchInsertItem struct {
}
