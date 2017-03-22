package GoGym

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"net/http"
	"reflect"
)

const (
	HTTPStatusMethodNotAllowed    = 405
	HTTPStatusOK                  = 200
	HTTPStatusNotFound            = 404
	HTTPStatusInternalServerError = 500
)

const (
	ServiceResponse = "Response"
)

//MIME Types
const (
	MIME_APP_JSON             = "application/json"
	MIME_APP_JSON_CHARSETUTF8 = MIME_APP_JSON + ";" + "charset=UTF-8"
	MIME_APP_FORM             = "application/x-www-form-urlencoded"
	MIME_APP_PB               = "application/protobuf"
)

// Response service
type Response struct {
	boss *Gym // Service Container

	rw         http.ResponseWriter
	statusCode int
	respone    interface{}
	header     http.Header
}

// Prepare is a method prepares the Response service
func (r *Response) Prepare(g *Gym) {
	r.WhoIsYourBoss(g)
}

// WhoIsYourBoss is a method sets the service container into the Response
func (r *Response) WhoIsYourBoss(g *Gym) {
	r.boss = g
}

// CallYourBoss is a method gets the service container
func (r *Response) CallYourBoss() *Gym {
	return r.boss
}

func (r *Response) CallMethod(method string, param []interface{}) []reflect.Value {
	return []reflect.Value{}
}

// JsonResponse is a method prepares the JSON response
func (r *Response) JsonResponse(resp interface{}, statusCode int, header http.Header) {
	r.respone = resp
	r.statusCode = statusCode
	var respHeader http.Header
	if header != nil {
		respHeader = header
		respHeader.Add("Content-Type", MIME_APP_JSON)
	} else {
		respHeader.Add("Content-Type", MIME_APP_JSON)
	}

	r.header = respHeader
}

// wait is a method does preparation for sending response
func (r *Response) wait(rw http.ResponseWriter) {
	r.rw = rw
	r.statusCode = HTTPStatusOK
}

// send is a method sending the http response
func (r *Response) send() {
	for k, v := range r.header {
		for _, h := range v {
			r.rw.Header().Add(k, h)
		}
	}
	// r.rw.WriteHeader(r.statusCode)
	rsp, err := json.Marshal(r.respone)
	if err != nil {
		// TODO: logging error
		glog.Error(fmt.Sprintf("JSON err: %s", err))
		r.statusCode = HTTPStatusInternalServerError
		rsp, _ = json.Marshal(map[string]string{"error": "foo"})
	}
	r.rw.WriteHeader(r.statusCode)
	r.rw.Write(rsp)
}
