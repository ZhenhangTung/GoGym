package GoGym

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
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
	App        *Gym // Service Container
	Rw         http.ResponseWriter
	StatusCode int
	Response   interface{}
	Header     http.Header
}

// Prepare is a method prepares the Response service
func (r *Response) Prepare(g *Gym) {
	r.InjectServiceContainer(g)
}

// WhoIsYourBoss is a method sets the service container into the Response
func (r *Response) InjectServiceContainer(g *Gym) {
	r.App = g
}

// CallYourBoss is a method gets the service container
func (r *Response) GetServiceContainer() *Gym {
	return r.App
}

func (r *Response) CallMethod(method string, param []interface{}) []reflect.Value {
	return []reflect.Value{}
}

// JsonResponse is a method prepares the JSON response
func (r *Response) JsonResponse(resp interface{}, statusCode int, header http.Header) {
	r.Response = resp
	r.StatusCode = statusCode
	var respHeader http.Header
	if header != nil {
		respHeader = header
		respHeader.Add("Content-Type", MIME_APP_JSON)
	} else {
		respHeader.Add("Content-Type", MIME_APP_JSON)
	}

	r.Header = respHeader
}

// wait is a method does preparation for sending response
func (r *Response) Wait(rw http.ResponseWriter) {
	r.Rw = rw
	r.StatusCode = HTTPStatusOK
}

// send is a method sending the http response
func (r *Response) Send() {
	for k, v := range r.Header {
		for _, h := range v {
			r.Rw.Header().Add(k, h)
		}
	}
	// r.rw.WriteHeader(r.statusCode)
	rsp, err := json.Marshal(r.Response)
	if err != nil {
		log.Error(fmt.Sprintf("JSON err: %s", err))
		r.StatusCode = HTTPStatusInternalServerError
		rsp, _ = json.Marshal(map[string]string{"error": "Error when parsing Json Response"})
	}
	r.Rw.WriteHeader(r.StatusCode)
	r.Rw.Write(rsp)
}
