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
	RouteCollection    map[int][]Route
}

// Prepare is a method prepares the router service
func (r *Router) Prepare(g *Gym) {
	r.InjectServiceContainer(g)
	r.ControllerRegistry = make(map[string]interface{})
	r.MethodVerbs = []string{GETMethod, POSTMethod, PUTMethod, PATCHMethod, DELETEMethod, OPTIONSMethod}
	r.RouteCollection = make(map[int][]Route)
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
	return nil
}

// NewRoute is a method that creates a new route to RouteCollection
func (r *Router) NewRoute(uri string, methods []string, action string) {
	if !r.IsActionLegal(action) {
		log.Fatalf("Action %s is illegal", action)
		return
	}
	var route Route
	route.uri = uri
	route.methods = methods
	route.action = action
	route.extractTokens(route.uri)
	route.compile(route.uri)
	route.node = r.calculateNode(route.uri)
	var routes []Route
	routes = r.RouteCollection[route.node]
	routes = append(routes, route)
	r.RouteCollection[route.node] = routes
}

// IsActionLegal checks if an action is legal
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

// ServeHTTP is a method serve http service
func (r *Router) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	requestService := r.GetServiceContainer().Request
	responseService := r.GetServiceContainer().Response
	requestService.accept(request)
	responseService.wait()
	routes := r.FindRoute(request.URL.Path)
	if routes == nil {
		rsp := make(map[string]interface{})
		rsp["err"] = "Not found"
		responseService.JsonResponse(rsp, HTTPStatusNotFound, http.Header{})

	} else {
		methodMatch := false
		var handlingRoute Route
		for k, route := range routes {
			for _, mth := range route.methods {
				if mth == request.Method {
					methodMatch = true
					handlingRoute = routes[k]
				}
			}
		}
		if !methodMatch {
			rsp := make(map[string]interface{})
			rsp["err"] = "Method not allowed"
			responseService.JsonResponse(rsp, HTTPStatusMethodNotAllowed, http.Header{})
		} else {
			// Binding path variables
			requestService.bindPathVar(handlingRoute.compiled.Tokens)
			// Handling request
			r.Handle(handlingRoute, rw, request)
		}
	}
	for k, v := range responseService.Header {
		for _, h := range v {
			rw.Header().Add(k, h)
		}
	}
	rw.WriteHeader(responseService.StatusCode)
	rw.Write(responseService.Response)
}

// FindRoute is a method finding a group of Route whose Uri is matched with request Uri
func (r *Router) FindRoute(uri string) []Route {
	var matchedCollection []Route
	matched := false
	node := r.calculateNode(uri)
	for i := node - 1; i >= 0; i-- {
		for _, route := range r.RouteCollection[node] {
			if route.match(uri) {
				route.assignValuesToTokens(uri)
				matchedCollection = append(matchedCollection, route)
				matched = true
			}
		}
		node--
	}
	if matched {
		return matchedCollection
	}
	return nil
}

// Handle is a method for using route passed in to handle the request
func (r *Router) Handle(route Route, rw http.ResponseWriter, request *http.Request) {
	actionSlice := strings.Split(route.action, "@")
	method := actionSlice[1]
	controllerKey := "*" + actionSlice[0]
	controller := r.ControllerRegistry[controllerKey]
	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(r.GetServiceContainer())
	reflect.ValueOf(controller).MethodByName(method).Call(in)
}

// RegisterController is a method registers controller
func (r *Router) RegisterController(controller interface{}) {
	controllerType := GetType(controller)
	r.ControllerRegistry[controllerType] = controller
}

// RegisterControllers is a method registers a bunch of controllers into controllerRegistry
func (r *Router) RegisterControllers(controllers []interface{}) {
	for _, v := range controllers {
		r.RegisterController(v)
	}
}

// calculateNode is a method calculates the node value of a uri
func (r *Router) calculateNode(uri string) int {
	var splitStr []string
	if uri[0:1] == Delimiter {
		splitStr = strings.Split(uri[1:], Delimiter)
	} else {
		splitStr = strings.Split(uri, Delimiter)
	}
	return len(splitStr)
}
