package GoGym

import (
	"reflect"
)

// getType is a function gets the type of value
func getType(value interface{}) string {
	if t := reflect.TypeOf(value); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
