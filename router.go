package GoGym

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"reflect"
	"strings"
)

const (
	ServiceRouter = "Router"
)

// Router service
type Router struct {
	// App is the Service Container
	App *Gym
	// ControllerRegistry is where all registered controllers exist
	ControllerRegistry map[string]interface{}
	MethodVerbs        []string
	RouteCollection    []Route
}

// Prepare is a method prepares the router service
func (r *Router) Prepare(g *Gym) {
	r.InjectServiceContainer(g)
	r.ControllerRegistry = make(map[string]interface{})
	r.MethodVerbs = []string{GETMethod, POSTMethod, PUTMethod, PATCHMethod, DELETEMethod, OPTIONSMethod}
}

// InjectServiceContainer is a method sets the service container into the Router
func (r *Router) InjectServiceContainer(g *Gym) {
	r.App = g
}

// GetServiceContainer is a method gets the service container
func (r *Router) GetServiceContainer() *Gym {
	return r.App
}

func (r *Router) CallMethod(method string, param []interface{}) []reflect.Value {
	return []reflect.Value{}
}

func (r *Router) NewRoute(uri string, methods []string, action string) {
	if !r.IsActionLegal(action) {
		log.Fatalf("Action %s is illegal", action)
		return
	}
	var route Route
	route.Uri = uri
	route.Methods = methods
	route.Action = action
	route.ExtractTokens(route.Uri)
	route.Compile(route.Uri)
	r.RouteCollection = append(r.RouteCollection, route)
}

func (r *Router) IsActionLegal(action string) bool {
	result := false
	if strings.Contains(action, "@") {
		result = true
	}
	return result
}

// Get is a method handles GET requests
func (r *Router) Get(path, action string) {
	methods := []string{GETMethod}
	r.NewRoute(path, methods, action)
}

// Post is a method handles POST requests
func (r *Router) Post(path, action string) {
	methods := []string{POSTMethod}
	r.NewRoute(path, methods, action)
}

// Put is a method handles PUT requests
func (r *Router) Put(path, action string) {
	methods := []string{PUTMethod}
	r.NewRoute(path, methods, action)
}

// Patch is a method handles PATCH requests
func (r *Router) Patch(path, action string) {
	methods := []string{PATCHMethod}
	r.NewRoute(path, methods, action)
}

// Options is a method handles Options requests
func (r *Router) Options(path, action string) {
	methods := []string{OPTIONSMethod}
	r.NewRoute(path, methods, action)
}

// Delete is a method handles Delete requests
func (r *Router) Delete(path, action string) {
	methods := []string{DELETEMethod}
	r.NewRoute(path, methods, action)
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	r.GetServiceContainer().Request.accept(request)
	r.GetServiceContainer().Response.wait(rw)
	//fmt.Println("req", request.URL.Path)
	routes := r.FindRoute(request.URL.Path)
	if routes == nil {
		rsp := make(map[string]interface{})
		rsp["err"] = "Not found"
		r.GetServiceContainer().Response.JsonResponse(rsp, HTTPStatusNotFound, http.Header{})

	} else {
		methodMatch := false
		var handlingRoute Route
		for k, route := range routes {
			for _, mth := range route.Methods {
				if mth == request.Method {
					methodMatch = true
					handlingRoute = routes[k]
				}
			}
		}

		if !methodMatch {
			rsp := make(map[string]interface{})
			rsp["err"] = "Method not allowed"
			r.GetServiceContainer().Response.JsonResponse(rsp, HTTPStatusNotFound, http.Header{})
		} else {
			r.Handle(handlingRoute, rw, request)
		}
	}
	r.GetServiceContainer().Response.send()

}

func (r *Router) FindRoute(uri string) []Route {
	var matchedCollection []Route
	matched := false
	for _, route := range r.RouteCollection {
		if route.Match(uri) {
			matchedCollection = append(matchedCollection, route)
			matched = true
		}
	}
	if matched {
		return matchedCollection
	}
	return nil
}

func (r *Router) Handle(route Route, rw http.ResponseWriter, request *http.Request) {
	actionSlice := strings.Split(route.Action, "@")
	method := actionSlice[1]
	controllerKey := "*" + actionSlice[0]
	controller := r.ControllerRegistry[controllerKey]
	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(r.GetServiceContainer())
	reflect.ValueOf(controller).MethodByName(method).Call(in)
}

func (r *Router) RegisterController(controller interface{}) {
	controllerType := GetType(controller)
	r.ControllerRegistry[controllerType] = controller
}
