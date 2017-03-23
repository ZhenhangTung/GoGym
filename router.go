package GoGym

import (
	"fmt"
	"github.com/golang/glog"
	"net/http"
	// "net/url"
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

	// v0.2
	methodVerbs          []string
	routeCollection      []Route
	controllerCollection map[string]interface{} //Currently, controller struct is not needed
}

// Prepare is a method prepares the router service
func (r *Router) Prepare(g *Gym) {
	r.WhoIsYourBoss(g)
	r.controllerRegistry = make(map[string]interface{})
	r.registeredPathAndController = make(map[string]map[string]map[string]string)
	r.methodVerbs = []string{GETMethod, POSTMethod, PUTMethod, PATCHMethod, DELETEMethod, OPTIONSMethod}
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
func (r *Router) Get(path, action string) {
	mapping := r.mappingRequestMethodWithControllerAndActions(GETMethod, path, action)
	r.registeredPathAndController[path] = mapping

	// v0.2
	r.addRoute([]string{GETMethod}, path, action)
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
		// v0.2
		fmt.Println("v02")
		// fmt.Println("url", request.RequestURI)
		r.findRoute(request.RequestURI)
		fmt.Println("endv02")

		r.CallYourBoss().Request.accept(request)
		r.CallYourBoss().Response.wait(rw)
		macthedControllers, ok := controllers[r.CallYourBoss().Request.Method]
		if !ok {
			rw.WriteHeader(HTTPStatusMethodNotAllowed)
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
		// r.CallYourBoss().Mux.GetMux()
		r.CallYourBoss().Mux.HandleFunc(path, r.HandleRequest(v))
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

// add route to route collection
func (r *Router) addRoute(verbs []string, uri string, action interface{}) {
	// check if there is a "@" in action
	legal := r.isActionLegal(action)
	if !legal {
		glog.Error(fmt.Sprintf("Action %s is illegal", action))
		return
	}
	route := Route{}
	route.methods = verbs
	route.uri = uri
	route.action = action
	r.routeCollection = append(r.routeCollection, route)
}

// findRoute routes
func (r *Router) findRoute(uri string) {
	for _, v := range r.routeCollection {
		r.check(uri, v)
	}
}

// check if there is a matched route
func (r *Router) check(uri string, route Route) {
	// fmt.Println(route.uri)
	// fmt.Println()
	if uri == "/3" {
		fmt.Println("url match")

	}
}

func (r *Router) isActionLegal(action interface{}) bool {
	result := false
	if GetType(action) == "string" {
		actionString := action.(string)
		result = strings.Contains(actionString, "@")
	}
	return result
}
