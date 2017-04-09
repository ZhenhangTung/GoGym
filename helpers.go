package GoGym

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"time"
)

// GetType is a function gets the type of value
func GetType(value interface{}) string {
	if t := reflect.TypeOf(value); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

// CallServiceMethodWithReflect is a method calls user's own service method
func CallServiceMethodWithReflect(g GymService, method string, param []interface{}) []reflect.Value {
	length := len(param)
	var in []reflect.Value
	if length > 0 {
		in = make([]reflect.Value, length)
		for k, v := range param {
			in[k] = reflect.ValueOf(v)
		}
	} else {
		in = []reflect.Value{}
	}
	results := reflect.ValueOf(g).MethodByName(method).Call(in)
	return results
}

func GetJson(method, url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	request, _ := http.NewRequest(method, url, strings.NewReader(""))
	response, err := myClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(target)
}
