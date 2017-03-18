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

// Router service
type Router struct {
	boss *Gym // Service Container

	// controllerRegistry is where all registered controllers exist
	controllerRegistry map[string]interface{}
	//registeredPathAndController is a mapping of paths and controllers
	registeredPathAndController map[string]map[string]map[string]string
}

// Prepare is a method prepares the router service
func (r *Router) Prepare(g *Gym) {
	r.WhoIsYourBoss(g)
	r.controllerRegistry = make(map[string]interface{})
	r.registeredPathAndController = make(map[string]map[string]map[string]string)
}

// WhoIsYourBoss is a method sets the service container into the Router
func (r *Router) WhoIsYourBoss(g *Gym) {
	r.boss = g
}

// CallYourBoss is a method gets the service container
func (r *Router) CallYourBoss() *Gym {
	return r.boss
}

func (r *Router) CallMethod(method string, param []interface{}) []reflect.Value {
	return []reflect.Value{}
}

// Get is a fucntion handles GET requests
func (r *Router) Get(path, controllerWithActionString string) {
	mapping := r.mappingRequestMethodWithControllerAndActions(GETMethod, path, controllerWithActionString)
	r.registeredPathAndController[path] = mapping
}

// Post is a fucntion handles POST requests
func (r *Router) Post(path, controllerWithActionString string) {
	mapping := r.mappingRequestMethodWithControllerAndActions(POSTMethod, path, controllerWithActionString)
	r.registeredPathAndController[path] = mapping
}

// Put is a method handles PUT requests
func (r *Router) Put(path, controllerWithActionString string) {
	mapping := r.mappingRequestMethodWithControllerAndActions(PUTMethod, path, controllerWithActionString)
	r.registeredPathAndController[path] = mapping
}

// Patch is a method handles PATCH requests
func (r *Router) Patch(path, controllerWithActionString string) {
	mapping := r.mappingRequestMethodWithControllerAndActions(PATCHMethod, path, controllerWithActionString)
	r.registeredPathAndController[path] = mapping
}

// Options is a method handles Options requests
func (r *Router) Options(path, controllerWithActionString string) {
	mapping := r.mappingRequestMethodWithControllerAndActions(OPTIONSMethod, path, controllerWithActionString)
	r.registeredPathAndController[path] = mapping
}

// Delete is a method handles Delete requests
func (r *Router) Delete(path, controllerWithActionString string) {
	mapping := r.mappingRequestMethodWithControllerAndActions(DELETEMethod, path, controllerWithActionString)
	r.registeredPathAndController[path] = mapping
}

// mappingRequestMethodWithControllerAndActions is a method for mapping request method with controllers
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

// HandleRequest is a method to handle http request
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

// RegisterHandleFunc is a method registers a handle function to handle request from path
func (r *Router) RegisterHandleFunc() {
	for k, v := range r.registeredPathAndController {
		path := k
		if !strings.HasPrefix(k, "/") {
			path = fmt.Sprintf("/%v", k)
		}
		http.HandleFunc(path, r.HandleRequest(v))
	}
}

// RegisterControllers is a method registers a struct of controllers into controllerRegistry
func (r *Router) RegisterControllers(controllers []interface{}) {
	for _, v := range controllers {
		r.RegisterController(v)
	}
}

// RegisterControllers is a method registers a controller into controllerRegistry
func (r *Router) RegisterController(controller interface{}) {
	controllerType := GetType(controller)
	r.controllerRegistry[controllerType] = controller
}
