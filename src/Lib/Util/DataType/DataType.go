package DataTypeUtil

import (
	"bytes"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func JsonFileToStruct(filePath string, data interface{}) error {
	u, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	e := json.Unmarshal(u, data)
	if e != nil {
		log.Println(e)
		return e
	}
	return nil
}
func YamlFileToStruct(filePath string, data interface{}) error {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, data)
	if err != nil {
		return err
	}
	return nil
}
func ObjectsToJson(data interface{}) string {
	byte, e := json.Marshal(data)
	if e != nil {
		log.Println(e)
	}
	return string(byte)
}
func JsonToObjects(u []byte, data interface{}) interface{} {
	e := json.Unmarshal(u, &data)
	if e != nil {
		log.Println(e)
	}
	return data
}
func ToString(value interface{}) string {
	switch value.(type) {
	case string:
		return value.(string)
	case int, int8, int64, int16, int32:
		return strconv.FormatInt(value.(int64), 10)
	case float32, float64:
		return strconv.FormatFloat(value.(float64), 'E', -1, 64)
	}
	o, _ := json.Marshal(value)
	return string(o)
}
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func Format(data interface{}, FormatString string) string {
	t := template.New("")
	t1, _ := t.Parse(FormatString)
	b := bytes.NewBufferString("")
	_ = t1.Execute(b, data)
	return b.String()
}
func JoinFormat(data []interface{}, FormatString []string, sep string, RemoveDuplicates bool) string {
	var strList []string
	for _, v := range FormatString {
		ss := strings.Trim(Format(data[0], v), " ")
		if ss != "" {
			strList = append(strList, ss)
		}
	}
	if RemoveDuplicates {
		return strings.Join(RemoveRepeatedElement(strList), sep)
	} else {

		return strings.Join(strList, sep)
	}
}
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func indexOf(a []string, e string) int {
	n := len(a)
	var i int = 0
	for ; i < n; i++ {
		if a[i] == e {
			return i
		}
	}
	return -1
}
func GetField(Columns []string, val reflect.Value, values []interface{}, pre string) {
	//time.NewTimer()
	t := GetTypeElem(val.Type())
	//log.Println(val.String())
	for i := 0; i < t.NumField(); i++ {
		f := GetTypeElem(t.Field(i).Type)
		vf := GetValueElem(GetValueElem(val).Field(i))
		url := t.Field(i).Tag.Get("db")
		if pre != "" {
			url = pre + t.Field(i).Tag.Get("db")
		}
		if t.Field(i).Tag.Get("db") == "" {
			continue
		}
		//log.Println(t.Field(i).Name, f.Name() == "Time")
		if f.Kind().String() == "struct" && f.Name() != "Time" {
			GetField(Columns, GetValueElem(val).Field(i), values, url+".")
		} else {
			//log.Println(url)
			if indexOf(Columns, url) > -1 {
				//if f.Name() == "Time" {
				//	aa := reflect.ValueOf("")
				//	values[indexOf(Columns, url)] = aa.Interface()
				//} else {
				set := reflect.New(vf.Type())
				vf.Set(reflect.ValueOf(set.Elem().Interface()))
				values[indexOf(Columns, url)] = vf.Addr().Interface()
				//}
				//log.Println(reflect.Indirect(vf).Type())
			}
		}
	}
}
func SetField(data map[string]interface{}, val reflect.Value, pre string) {
	t := GetTypeElem(val.Type())
	for i := 0; i < t.NumField(); i++ {
		f := GetTypeElem(t.Field(i).Type)
		vf := GetValueElem(val.Field(i))
		if t.Field(i).Tag.Get("db") == "" {
			continue
		}
		if f.Kind().String() == "struct" {
			SetField(data, val.Field(i), pre+pre+t.Field(i).Tag.Get("db")+".")
		} else {
			if !vf.CanSet() {
				vf.Set(reflect.New(t.Field(i).Type))
			}
			if v, ok := data[pre+t.Field(i).Tag.Get("db")]; ok {
				vf.Set(reflect.ValueOf(v))
			}
		}
	}
}
func GetTypeElem(v reflect.Type) reflect.Type {
	for {
		if v.Kind().String() == "ptr" {
			v = v.Elem()
		} else {
			return v
		}
	}
}
func GetValueElem(v reflect.Value) reflect.Value {
	for {
		//fmt.Println(v.Kind(), v.Type().String())
		if v.Kind().String() == "ptr" {
			v = v.Elem()
		} else {
			return v
		}
	}
}
