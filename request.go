package GoGym

import (
	"net/http"
	"net/url"
	"reflect"
)

const (
	GETMethod     = "GET"
	POSTMethod    = "POST"
	PUTMethod     = "PUT"
	PATCHMethod   = "PATCH"
	DELETEMethod  = "DELETE"
	OPTIONSMethod = "OPTIONS"
)

const (
	ServiceRequest = "Request"
)

// Request service
type Request struct {
	App *Gym // Service Container

	Method string
	Header http.Header
	Query  url.Values
	Form   url.Values
}

// Prepare is a method prepares the Request service
func (r *Request) Prepare(g *Gym) {
	r.InjectServiceContainer(g)
}

// WhoIsYourBoss is a method sets the service container into the Request
func (r *Request) InjectServiceContainer(g *Gym) {
	r.App = g
}

// CallYourBoss is a method gets the service container
func (r *Request) GetServiceContainer() *Gym {
	return r.App
}

func (r *Request) CallService(method string, param []interface{}) []reflect.Value {
	return []reflect.Value{}
}

// accept is a method gets the http request and parse it
func (r *Request) accept(request *http.Request) {
	request.ParseForm()
	r.Method = request.Method
	r.Query = request.Form
	r.Form = request.PostForm
	r.Header = request.Header
}
