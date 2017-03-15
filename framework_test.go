package GoGym

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

func init() {
	fmt.Println("Tests started")
}

var formTest url.Values

type HelloFormat struct {
	Hello string
}

var helloResponse = map[string]string{"Hello": "World"}

type IndexController struct {
}

func (IndexController *IndexController) Index(request map[string]url.Values, headers http.Header) (statusCode int, response interface{}) {
	return 200, helloResponse
}

func (IndexController *IndexController) Post(request map[string]url.Values, headers http.Header) (statusCode int, response interface{}) {
	return 200, helloResponse
}

func (IndexController *IndexController) Put(request map[string]url.Values, headers http.Header) (statusCode int, response interface{}) {
	return 200, helloResponse
}

func (IndexController *IndexController) Delete(request map[string]url.Values, headers http.Header) (statusCode int, response interface{}) {
	return 200, helloResponse
}

func (IndexController *IndexController) Options(request map[string]url.Values, headers http.Header) (statusCode int, response interface{}) {
	return 200, helloResponse
}

func (IndexController *IndexController) Patch(request map[string]url.Values, headers http.Header) (statusCode int, response interface{}) {
	return 200, helloResponse
}

func (IndexController *IndexController) QueryForm(request map[string]url.Values, headers http.Header) (statusCode int, response interface{}) {
	formTest = request["query"]
	return 200, helloResponse
}

func (IndexController *IndexController) PostForm(request map[string]url.Values, headers http.Header) (statusCode int, response interface{}) {
	formTest = request["form"]
	return 200, helloResponse
}

func (IndexController *IndexController) SetHeaders(request map[string]url.Values, headers http.Header) (statusCode int, response interface{}, responseHeader http.Header) {

	return 200, helloResponse, http.Header{"Foo": {"Bar", "Baz"}, "Gogym": {"Yeah"}}
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

func TestRequestWithQuery(t *testing.T) {
	var apiService = Prepare()
	apiService.Get("/requests/form-method/query", "IndexController@QueryForm")
	apiService.RegisterController(&IndexController{})
	go apiService.Serve(3000)
	var requestQuery = url.Values{"api_key": {"gogym"}, "foo": {"bar&baz"}}
	request, _ := http.NewRequest("GET", "http://localhost:3000/requests/form-method/query", nil)
	q := request.URL.Query()
	q.Add("api_key", "gogym")
	q.Add("foo", "bar&baz")
	request.URL.RawQuery = q.Encode()
	response, err := myClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	if !reflect.DeepEqual(formTest, requestQuery) {
		t.Error("received query is not same as requested query")
	}
}

func TestRequestWithForm(t *testing.T) {
	var apiService = Prepare()
	apiService.Post("/requests/form-method/form", "IndexController@PostForm")
	apiService.RegisterController(&IndexController{})
	go apiService.Serve(3000)
	requestForm := url.Values{"foo": {"bar", "baz"}}
	myClient.PostForm("http://localhost:3000/requests/form-method/form", requestForm)
	if !reflect.DeepEqual(formTest, requestForm) {
		t.Error("received form is not same as requested form")
	}
}

func TestHeader(t *testing.T) {
	var apiService = Prepare()
	apiService.Get("/requests/headers", "IndexController@SetHeaders")
	apiService.RegisterController(&IndexController{})
	go apiService.Serve(3000)
	r, err := myClient.Get("http://localhost:3000/requests/headers")
	if err != nil {
		t.Error(err)
	}
	responseHeaders := r.Header
	expectedHeaders := http.Header{"Foo": {"Bar", "Baz"}, "Gogym": {"Yeah"}, "Content-Type": {"application/json"}}

	// Check if all expected headers exist in response
	for k, v := range expectedHeaders {
		header, isset := responseHeaders[k]
		if !isset {
			t.Error("response headers didn't match as expected")
		}
		if !reflect.DeepEqual(v, header) {
			t.Error("content of the response header didn't match as expected")
		}
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
