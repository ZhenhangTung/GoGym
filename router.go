package GoGym

import (
	"fmt"
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
	App *Gym // Service Container

	// controllerRegistry is where all registered controllers exist
	ControllerRegistry map[string]interface{}

	// v0.2
	MethodVerbs     []string
	RouteCollection []Route
	//ControllerCollection map[string]interface{}
}

// Prepare is a method prepares the router service
func (r *Router) Prepare(g *Gym) {
	r.InjectServiceContainer(g)
	r.ControllerRegistry = make(map[string]interface{})
	r.MethodVerbs = []string{GETMethod, POSTMethod, PUTMethod, PATCHMethod, DELETEMethod, OPTIONSMethod}
}

// WhoIsYourBoss is a method sets the service container into the Router
func (r *Router) InjectServiceContainer(g *Gym) {
	r.App = g
}

// CallYourBoss is a method gets the service container
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
	fmt.Print("hello world")
	route := r.FindRoute(request.RequestURI)
	if route.Action == "" {
		// 404
		fmt.Println("404")
		//rsp := map(string){"err": "not found"}
		rsp := make(map[string]interface{})
		r.GetServiceContainer().Response.JsonResponse(rsp, HTTPStatusNotFound, http.Header{})
		r.GetServiceContainer().Response.send()

	}
	methodMatch := false
	for _, m := range route.Methods {
		if m == request.Method {
			methodMatch = true
		}
	}
	if !methodMatch {
		// 405
		fmt.Println("405")
		return
	}
	r.Handle(route, rw, request)

}

func (r *Router) FindRoute(uri string) Route {
	for _, route := range r.RouteCollection {
		//r.check(uri, v)
		if route.Match(uri) {
			return route
		}
	}
	return Route{}
}

func (r *Router) Handle(route Route, rw http.ResponseWriter, request *http.Request) {
	actionSlice := strings.Split(route.Action, "@")
	method := actionSlice[1]
	controllerKey := "*" + actionSlice[0]
	controller := r.ControllerRegistry[controllerKey]
	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(r.GetServiceContainer())
	reflect.ValueOf(controller).MethodByName(method).Call(in)
	//r.GetServiceContainer().Response.send()
}

func (r *Router) RegisterController(controller interface{}) {
	controllerType := GetType(controller)
	r.ControllerRegistry[controllerType] = controller
}
