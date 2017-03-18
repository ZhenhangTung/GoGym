package GoGym

import (
	"fmt"
	"net/http"
	"reflect"
)

// Gym is a service container
type Gym struct {
	Router   *Router
	Request  *Request
	Response *Response
	service  map[string]GymService
}

// RegisterService registers service into service container
// Maybe there would be some other logic in the future.
// It would be better to call RegisterService instead of calling bindService,
// that's why it is a private method
func (g *Gym) RegisterService(name string, service GymService) {
	g.bindService(name, service)
}

func (g *Gym) RegisterServices(services map[string]GymService) {
	for name, service := range services {
		g.bindService(name, service)
	}
}

func (g *Gym) bindService(name string, service GymService) {
	g.service[name] = service
}

func (g *Gym) GetService(name string) GymService {
	return g.service[name]
}

func (r *Request) CallServiceMethod(service GymService, method string, param []interface{}) {
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
	reflect.ValueOf(service).MethodByName(method).Call(in)
}

func (g *Gym) Prepare() *Gym {
	g.service = make(map[string]GymService)
	g.Router = new(Router)
	g.Router.Prepare(g)
	g.Request = new(Request)
	g.Request.Prepare(g)
	g.Response = new(Response)
	g.Response.Prepare(g)
	return g
}

// Serve is a function
func (g *Gym) Open(port int) {
	g.Router.RegisterHandleFunc()
	fullPort := fmt.Sprintf(":%d", port)
	http.ListenAndServe(fullPort, nil)
}