package MessageLib

import (
	"encoding/json"
	"log"
	"net/url"
	"reflect"
)

var action = map[string]ActionModel{}

type ActionModel struct {
	Name          string
	ParameterType reflect.Type
	Action        reflect.Value
}

func AddAction(name string, myAction interface{}) ActionModel {
	action[name] = ActionModel{
		Name:          name,
		ParameterType: reflect.TypeOf(myAction).In(0),
		Action:        reflect.ValueOf(myAction),
	}
	return action[name]
}
func RunFromAction(name string, val url.Values, client func(callResult QueueItem)) {
	if _, ok := action[name]; ok {
		result := reflect.New(action[name].ParameterType)
		err := Unmarshal(val, result.Interface())
		if err != nil {
			return
		}
		callResult := action[name].Action.Call([]reflect.Value{result.Elem()})
		if len(callResult) > 0 {
			for i := 0; i < len(callResult); i++ {
				if re, ok := callResult[i].Interface().(QueueItem); ok {
					client(re)
				}
			}
		}
		return
	}
}
func RunAction(item QueueItem, client func(callResult QueueItem)) {
	if _, ok := action[item.Name]; !ok {
		//log.Println("action not exist：", item)
		return
	}
	result := reflect.New(action[item.Name].ParameterType)
	err := json.Unmarshal([]byte(item.Data), result.Interface())
	if err != nil {
		log.Println("data type error：", item)
	}
	callResult := action[item.Name].Action.Call([]reflect.Value{result.Elem()})
	if len(callResult) > 0 {
		for i := 0; i < len(callResult); i++ {
			if re, ok := callResult[i].Interface().(QueueItem); ok {
				client(re)
			}
		}
	}
}
