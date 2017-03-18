package GoGym

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

const (
	ServiceRouter = "Router"
)

type Router struct {
	boss *Gym // Service Container

	// controllerRegistry is where all registered controllers exist
	controllerRegistry map[string]interface{}
	//registeredPathAndController is a mapping of paths and controllers
	registeredPathAndController map[string]map[string]map[string]string
}

func (r *Router) Prepare(g *Gym) {
	r.WhoIsYourBoss(g)
	r.controllerRegistry = make(map[string]interface{})
	r.registeredPathAndController = make(map[string]map[string]map[string]string)
}

func (r *Router) WhoIsYourBoss(g *Gym) {
	r.boss = g
}

func (r *Router) CallYourBoss() *Gym {
	return r.boss
}

func (r *Router) Get(path, controllerWithActionString string) {
	mapping := r.mappingRequestMethodWithControllerAndActions(GETMethod, path, controllerWithActionString)
	r.registeredPathAndController[path] = mapping
}

func (r *Router) Post(path, controllerWithActionString string) {
	mapping := r.mappingRequestMethodWithControllerAndActions(POSTMethod, path, controllerWithActionString)
	r.registeredPathAndController[path] = mapping
}

func (r *Router) Put(path, controllerWithActionString string) {
	mapping := r.mappingRequestMethodWithControllerAndActions(PUTMethod, path, controllerWithActionString)
	r.registeredPathAndController[path] = mapping
}

func (r *Router) Patch(path, controllerWithActionString string) {
	mapping := r.mappingRequestMethodWithControllerAndActions(PATCHMethod, path, controllerWithActionString)
	r.registeredPathAndController[path] = mapping
}

func (r *Router) Options(path, controllerWithActionString string) {
	mapping := r.mappingRequestMethodWithControllerAndActions(OPTIONSMethod, path, controllerWithActionString)
	r.registeredPathAndController[path] = mapping
}

func (r *Router) Delete(path, controllerWithActionString string) {
	mapping := r.mappingRequestMethodWithControllerAndActions(DELETEMethod, path, controllerWithActionString)
	r.registeredPathAndController[path] = mapping
}

// mappingRequestMethodWithControllerAndActions is a function for mapping request method with controllers
// which containing actions
func (r *Router) mappingRequestMethodWithControllerAndActions(requestMethod, path, controllerWithActionString string) map[string]map[string]string {
	mappingResult := make(map[string]map[string]string)
	if length := len(r.registeredPathAndController[path]); length > 0 {
		mappingResult = r.registeredPathAndController[path]
	}
	controllerAndActionSlice := strings.Split(controllerWithActionString, "@")
	controller := controllerAndActionSlice[0]
	action := controllerAndActionSlice[1]
	controllerAndActionMap := map[string]string{controller: action}
	mappingResult[requestMethod] = controllerAndActionMap
	return mappingResult
}

// HandleRequest is a function to handle http request
func (r *Router) HandleRequest(controllers map[string]map[string]string) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		r.CallYourBoss().Request.accept(request)
		r.CallYourBoss().Response.wait(rw)
		macthedControllers, ok := controllers[r.CallYourBoss().Request.Method]
		if !ok {
			rw.WriteHeader(HTTPMethodNotAllowed)
		}
		for k, v := range macthedControllers {
			controllerKey := "*" + k
			controller := r.controllerRegistry[controllerKey]
			in := make([]reflect.Value, 1)
			in[0] = reflect.ValueOf(r.CallYourBoss())
			reflect.ValueOf(controller).MethodByName(v).Call(in)
			r.CallYourBoss().Response.send()
		}
	}
}

// RegisterHandleFunc is a function registers a handle function to handle request from path
func (r *Router) RegisterHandleFunc() {
	for k, v := range r.registeredPathAndController {
		path := k
		if !strings.HasPrefix(k, "/") {
			path = fmt.Sprintf("/%v", k)
		}
		http.HandleFunc(path, r.HandleRequest(v))
	}
}

// RegisterControllers is a function registers a struct of controllers into controllerRegistry
func (r *Router) RegisterControllers(controllers []interface{}) {
	for _, v := range controllers {
		r.RegisterController(v)
	}
}

// RegisterControllers is a function registers a controller into controllerRegistry
func (r *Router) RegisterController(controller interface{}) {
	controllerType := getType(controller)
	r.controllerRegistry[controllerType] = controller
}
