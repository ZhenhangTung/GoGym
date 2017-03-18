package GoGym

import (
	"net/http"
	"net/url"
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
	boss *Gym // Service Container

	Method string
	Header http.Header
	Query  url.Values
	Form   url.Values
}

// Prepare is a function prepares the Request service
func (r *Request) Prepare(g *Gym) {
	r.WhoIsYourBoss(g)
}

// WhoIsYourBoss is a function sets the service container into the Request
func (r *Request) WhoIsYourBoss(g *Gym) {
	r.boss = g
}

// CallYourBoss is a function gets the service container
func (r *Request) CallYourBoss() *Gym {
	return r.boss
}

// accept is a function gets the http request and parse it
func (r *Request) accept(request *http.Request) {
	request.ParseForm()
	r.Method = request.Method
	r.Query = request.Form
	r.Form = request.PostForm
	r.Header = request.Header
}
