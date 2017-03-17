package GoGym

import (
	"fmt"
	"net/http"
)

type Response struct {
	statusCode int
	respone    interface{}
	header     http.Header
}

func (r *Response) JsonResponse(statusCode int, resp interface{}, header http.Header) {

}
