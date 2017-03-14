package GoGym

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
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

func (IndexController *IndexController) Post(values url.Values, headers http.Header) (statusCode int, response interface{}) {
	return 200, helloResponse
}

func (IndexController *IndexController) Put(values url.Values, headers http.Header) (statusCode int, response interface{}) {
	return 200, helloResponse
}

func (IndexController *IndexController) Delete(values url.Values, headers http.Header) (statusCode int, response interface{}) {
	return 200, helloResponse
}

func (IndexController *IndexController) Options(values url.Values, headers http.Header) (statusCode int, response interface{}) {
	return 200, helloResponse
}

func (IndexController *IndexController) Patch(values url.Values, headers http.Header) (statusCode int, response interface{}) {
	return 200, helloResponse
}

func TestGet(t *testing.T) {
	var apiService = Prepare()
	apiService.Get("/", "IndexController@Index")
	apiService.RegisterController(&IndexController{})
	go apiService.Serve(3000)
	hello := &HelloFormat{}
	err := getJsonWithGetMethod("http://localhost:3000", hello)
	if err != nil {
		t.Error("resp error")
	}
	if hello.Hello != "World" {
		t.Error("resp is not equal as expected")
	}
}

func TestPost(t *testing.T) {
	var apiService = Prepare()
	apiService.Post("/requests/post", "IndexController@Post")
	apiService.RegisterController(&IndexController{})
	go apiService.Serve(3000)
	hello := &HelloFormat{}
	err := getJsonWithPostMethod("http://localhost:3000/requests/post", hello)
	if err != nil {
		t.Error("resp error")
	}
	if hello.Hello != "World" {
		t.Error("resp is not equal as expected")
	}
}

func TestPut(t *testing.T) {
	var apiService = Prepare()
	apiService.Put("/requests/put", "IndexController@Put")
	apiService.RegisterController(&IndexController{})
	go apiService.Serve(3000)
	hello := &HelloFormat{}
	err := getJson(PUTMethod, "http://localhost:3000/requests/put", hello)
	if err != nil {
		t.Error("resp error")
	}
	if hello.Hello != "World" {
		t.Error("resp is not equal as expected")
	}
}

func TestDelete(t *testing.T) {
	var apiService = Prepare()
	apiService.Delete("/requests/delete", "IndexController@Delete")
	apiService.RegisterController(&IndexController{})
	go apiService.Serve(3000)
	hello := &HelloFormat{}
	err := getJson(DELETEMethod, "http://localhost:3000/requests/delete", hello)
	if err != nil {
		t.Error("resp error")
	}
	if hello.Hello != "World" {
		t.Error("resp is not equal as expected")
	}
}

func TestOptions(t *testing.T) {
	var apiService = Prepare()
	apiService.Options("/requests/options", "IndexController@Options")
	apiService.RegisterController(&IndexController{})
	go apiService.Serve(3000)
	hello := &HelloFormat{}
	err := getJson(OPTIONSMethod, "http://localhost:3000/requests/options", hello)
	if err != nil {
		t.Error("resp error")
	}
	if hello.Hello != "World" {
		t.Error("resp is not equal as expected")
	}
}

func TestPatch(t *testing.T) {
	var apiService = Prepare()
	apiService.Patch("/requests/patch", "IndexController@Patch")
	apiService.RegisterController(&IndexController{})
	go apiService.Serve(3000)
	hello := &HelloFormat{}
	err := getJson(PATCHMethod, "http://localhost:3000/requests/patch", hello)
	if err != nil {
		t.Error("resp error")
	}
	if hello.Hello != "World" {
		t.Error("resp is not equal as expected")
	}
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJsonWithGetMethod(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func getJsonWithPostMethod(url string, target interface{}) error {
	r, err := myClient.Post(url, "application/json", strings.NewReader(""))
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func getJson(method, url string, target interface{}) error {
	request, _ := http.NewRequest(method, url, strings.NewReader(""))
	response, err := myClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(target)
}
