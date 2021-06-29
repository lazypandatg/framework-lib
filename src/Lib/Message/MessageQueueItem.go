package MessageLib

import "encoding/json"

type QueueItem struct {
	Name      string `json:"name"`
	Data      string `json:"data"`
	Status    bool   `json:"status"`
	Client    string `json:"client"`
	MessageId string `json:"messageId"`
}

func (_this *QueueItem) GetData(v interface{}) {
	err := json.Unmarshal([]byte(_this.Data), v)
	if err != nil {
		return
	}
}
func NewQueueItem(status bool, data interface{}) QueueItem {
	marshal, err := json.Marshal(data)
	if err != nil {
		return QueueItem{}
	}
	return QueueItem{Status: status, Data: string(marshal)}
}
