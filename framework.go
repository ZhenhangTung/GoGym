package goGym

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

const (
	GetMethod    = "GET"
	PostMethod   = "POST"
	PutMethod    = "PUT"
	PatchMethod  = "PATCH"
	DeleteMethod = "DELETE"
)

type APIService struct {
	controllerRegistry map[string]interface{}
	// TODO：拆分registeredPathAndController
	registeredPathAndController map[string]map[string]map[string]string
}

func (api *APIService) Get(path, controllerWithActionString string) {
	mapping := api.mappingRequestMethodWithControllerAndActions(GetMethod, controllerWithActionString)
	api.registeredPathAndController[path] = mapping
	fmt.Println(api.registeredPathAndController)

}

func (*APIService) mappingRequestMethodWithControllerAndActions(requestMethod, controllerWithActionString string) map[string]map[string]string {
	controllerAndActionSlice := strings.Split(controllerWithActionString, "@")
	controller := controllerAndActionSlice[0]
	action := controllerAndActionSlice[1]
	controllerAndActionMap := map[string]string{controller: action}
	mappingResult := map[string]map[string]string{requestMethod: controllerAndActionMap}
	return mappingResult
}

func (api *APIService) HandleRequest(controllers map[string]map[string]string) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		method := request.Method
		macthedControllers := controllers[method]

		for k, v := range macthedControllers {
			controllerKey := "*" + k
			controller := api.controllerRegistry[controllerKey]
			in := make([]reflect.Value, 2)
			in[0] = reflect.ValueOf(request.Form)
			in[1] = reflect.ValueOf(request.Header)
			reflect.ValueOf(controller).MethodByName(v).Call(in)

		}
	}
}

func (api *APIService) RegisterHandleFunc() {
	for k, v := range api.registeredPathAndController {
		path := fmt.Sprintf("/%v", k)
		http.HandleFunc(path, api.HandleRequest(v))
	}
}

func (api *APIService) RegisterControllers() {

}

func (api *APIService) RegisterController(controller interface{}) {
	controllerType := getType(controller)
	fmt.Println(controllerType)
	api.controllerRegistry[controllerType] = controller
	fmt.Println(api.controllerRegistry)
}

func getType(value interface{}) string {
	if t := reflect.TypeOf(value); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func (api *APIService) Serve(port int) {
	api.RegisterHandleFunc()
	fullPort := fmt.Sprintf(":%d", port)
	http.ListenAndServe(fullPort, nil)
}

func NewAPIService() *APIService {
	var apiService = new(APIService)
	apiService.controllerRegistry = make(map[string]interface{})
	apiService.registeredPathAndController = make(map[string]map[string]map[string]string)
	return apiService
}
