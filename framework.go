package goGym

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type RouteService struct {
}

type BaseController struct {
}

func JsonResponse() {

}

var controllerRegistry = make(map[string]interface{})

var Route map[string]map[string]map[string]string

func (rs *RouteService) Get(path, controllerWithActionString string) {
	controllerAndActionSlice := strings.Split(controllerWithActionString, "@")
	fmt.Println(controllerAndActionSlice)
	controllerAndActionMap := make(map[string]string)
	controllerAndActionMap[controllerAndActionSlice[0]] = controllerAndActionSlice[1]
	fmt.Println(controllerAndActionMap)
	routeAndControllerAction := map[string]map[string]string{"GET": controllerAndActionMap}
	Route = make(map[string]map[string]map[string]string)
	Route[path] = routeAndControllerAction
	fmt.Println(Route)

}

func HandleRequest(methodAndControllers map[string]map[string]string) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		// fmt.Println("ok")
		request.ParseForm()
		method := request.Method
		// values := request.Form
		controllers := methodAndControllers[method]
		fmt.Println(controllers)

		for k, v := range controllers {
			fmt.Println(k)
			fmt.Println(v)
			// controller := controllerRegistry[k]

		}
		// IndexController.index(values)
	}
}

func RegisterRouteService() {
	for k, v := range Route {
		path := fmt.Sprintf("/%v", k)
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println(path)
		http.HandleFunc(path, HandleRequest(v))
	}
}

func RegisterControllers() {

}

func RegisterController(controller interface{}, actionSlice []interface{}) {
	controllerType := getType(controller)
	fmt.Println(controllerType)
	controllerRegistry[controllerType] = controller
	fmt.Println(controllerRegistry)
	fmt.Println(actionSlice)
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
