package MessageLib

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type Queue struct {
	Pre       string
	PushName  string
	PopName   string
	Count     int
	PushQueue *redis.Client
	PopQueue  *redis.Client
}

func NewQueue(config Config) *Queue {
	q := &Queue{}
	q.Init(config)
	return q
}

func (_this *Queue) Init(config Config) {
	_this.Pre = config.Pre
	_this.PushName = config.PushName
	_this.PopName = config.PopName
	_this.PopQueue = redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.DataBase,
	})
	_this.PushQueue = redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.DataBase,
	})
}

func (_this *Queue) Listener() {
	go func() {
		for true {
			data := QueueItem{}
			result, err := _this.PopQueue.BLPop(10*time.Second, _this.PopName).Result()
			if err != nil {
				continue
			}
			err = json.Unmarshal([]byte(result[1]), &data)
			if err != nil {
				log.Println("数据格式不正确：", result)
				log.Println(err)
				continue
			}
			go RunAction(data, func(callResult QueueItem) {})
		}
	}()
}

func (_this *Queue) Push(name string, val interface{}) {
	_this.PushService(_this.PushName, name, val)
}

func (_this *Queue) PushSync(name string, val interface{}) {
	_this.PushService(_this.PushName, name, val)
}

func (_this *Queue) PushService(service, name string, val interface{}) {
	marshal, err := json.Marshal(val)
	if err != nil {
		return
	}
	post, err := json.Marshal(QueueItem{Name: name, Data: string(marshal)})
	if err != nil {
		return
	}
	//log.Println(service)
	_this.PushQueue.LPush(service, post)
}
