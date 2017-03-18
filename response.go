package GoGym

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	HTTPMethodNotAllowed = 405
	HTTPOk               = 200
)

const (
	ServiceResponse = "Response"
)

type Response struct {
	boss *Gym // Service Container

	rw         http.ResponseWriter
	statusCode int
	respone    interface{}
	header     http.Header
}

func (r *Response) Prepare(g *Gym) {
	r.WhoIsYourBoss(g)
}

func (r *Response) WhoIsYourBoss(g *Gym) {
	r.boss = g
}

func (r *Response) CallYourBoss() *Gym {
	return r.boss
}

func (r *Response) JsonResponse(resp interface{}, statusCode int, header http.Header) {
	r.respone = resp
	r.statusCode = statusCode
	var respHeader http.Header
	if header != nil {
		respHeader = header
		respHeader.Add("Content-Type", "application/json")
	} else {
		respHeader.Add("Content-Type", "application/json")
	}

	r.header = respHeader
}

func (r *Response) wait(rw http.ResponseWriter) {
	r.rw = rw
	r.statusCode = HTTPOk
}

func (r *Response) send() {
	for k, v := range r.header {
		for _, h := range v {
			r.rw.Header().Add(k, h)
		}
	}
	r.rw.WriteHeader(r.statusCode)
	rsp, err := json.Marshal(r.respone)
	if err != nil {
		// TODO: logging error
		fmt.Println("JSON err:", err)
	}
	r.rw.Write(rsp)
}
