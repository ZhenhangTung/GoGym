package GoGym

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
	"time"
)

type HelloFormat struct {
	Hello string
}

var helloResponse = map[string]string{"Hello": "World"}

type IndexController struct {
}

func (IndexController *IndexController) Index(values url.Values, headers http.Header) (statusCode int, response interface{}) {
	return 200, helloResponse
}

func TestGet(t *testing.T) {
	var apiService = Prepare()
	apiService.Get("/", "IndexController@Index")
	apiService.RegisterController(&IndexController{})
	go apiService.Serve(3000)
	hello := &HelloFormat{}
	err := getJson("http://localhost:3000", hello)
	if err != nil {
		t.Error("resp error")
	}
	if hello.Hello != "World" {
		t.Error("resp is not equal as expected")
	}
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
