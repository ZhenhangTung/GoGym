package GoGym

import (
	// "fmt"
	// "encoding/json"
	"net/http"
	"net/url"
	// "reflect"
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

type Request struct {
	boss *Gym // Service Container

	Method string
	Header http.Header
	Query  url.Values
	Form   url.Values
}

func (r *Request) Prepare(g *Gym) {
	r.WhoIsYourBoss(g)
}

func (r *Request) WhoIsYourBoss(g *Gym) {
	r.boss = g
}

func (r *Request) CallYourBoss() *Gym {
	return r.boss
}

func (r *Request) accept(request *http.Request) {
	request.ParseForm()
	r.Method = request.Method
	r.Query = request.Form
	r.Form = request.PostForm
	r.Header = request.Header
}
