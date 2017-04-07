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

var port int
var baseUri string

func init() {
	fmt.Println("Tests started")
	port = 3000
}

var formTest url.Values

type HelloFormat struct {
	Hello string
}

var helloResponse = map[string]string{"Hello": "World"}

type IndexController struct {
}

func (IndexController *IndexController) Index(g *Gym) {
	g.Response.JsonResponse(helloResponse, 200, http.Header{})
}

func (IndexController *IndexController) Post(g *Gym) {
	g.Response.JsonResponse(helloResponse, 200, http.Header{})
}

func (IndexController *IndexController) Put(g *Gym) {
	g.Response.JsonResponse(helloResponse, 200, http.Header{})
}

func (IndexController *IndexController) Delete(g *Gym) {
	g.Response.JsonResponse(helloResponse, 200, http.Header{})
}

func (IndexController *IndexController) Options(g *Gym) {
	g.Response.JsonResponse(helloResponse, 200, http.Header{})
}

func (IndexController *IndexController) Patch(g *Gym) {
	g.Response.JsonResponse(helloResponse, 200, http.Header{})
}

func (IndexController *IndexController) QueryForm(g *Gym) {
	formTest = g.Request.Query
	g.Response.JsonResponse(helloResponse, 200, http.Header{})
}

func (IndexController *IndexController) PostForm(g *Gym) {
	formTest = g.Request.Form
	g.Response.JsonResponse(helloResponse, 200, http.Header{})
}

func (IndexController *IndexController) SetHeaders(g *Gym) {
	g.Response.JsonResponse(helloResponse, 200, http.Header{"Foo": {"Bar", "Baz"}, "Gogym": {"Yeah"}})
}

func TestGet(t *testing.T) {
	var gym = new(Gym)
	gym.Prepare()
	gym.Router.Get("/", "IndexController@Index")
	gym.Router.RegisterController(&IndexController{})
	go gym.OpenAt(port)
	hello := &HelloFormat{}
	baseUri = fmt.Sprintf("http://localhost:%v", port)
	port++
	err := getJsonWithGetMethod(baseUri, hello)
	if err != nil {
		t.Error("resp error")
	}
	if hello.Hello != "World" {
		t.Error("resp is not equal as expected")
	}
}

func TestPost(t *testing.T) {
	var gym = new(Gym)
	gym.Prepare()
	gym.Router.Post("/requests/post", "IndexController@Post")
	gym.Router.RegisterController(&IndexController{})
	go gym.OpenAt(port)
	hello := &HelloFormat{}
	baseUri = fmt.Sprintf("http://localhost:%v", port)
	port++
	err := getJsonWithPostMethod(baseUri+"/requests/post", hello)
	if err != nil {
		t.Error("resp error")
	}
	if hello.Hello != "World" {
		t.Error("resp is not equal as expected")
	}
}

func TestPut(t *testing.T) {
	var gym = new(Gym)
	gym.Prepare()
	gym.Router.Put("/requests/put", "IndexController@Put")
	gym.Router.RegisterController(&IndexController{})
	go gym.OpenAt(port)
	hello := &HelloFormat{}
	baseUri = fmt.Sprintf("http://localhost:%v", port)
	port++
	err := getJson(PUTMethod, baseUri+"/requests/put", hello)
	if err != nil {
		t.Error("resp error")
	}
	if hello.Hello != "World" {
		t.Error("resp is not equal as expected")
	}
}

func TestDelete(t *testing.T) {
	var gym = new(Gym)
	gym.Prepare()
	gym.Router.Delete("/requests/delete", "IndexController@Delete")
	gym.Router.RegisterController(&IndexController{})
	go gym.OpenAt(port)
	hello := &HelloFormat{}
	baseUri = fmt.Sprintf("http://localhost:%v", port)
	port++
	err := getJson(DELETEMethod, baseUri+"/requests/delete", hello)
	if err != nil {
		t.Error("resp error")
	}
	if hello.Hello != "World" {
		t.Error("resp is not equal as expected")
	}
}

func TestOptions(t *testing.T) {
	var gym = new(Gym)
	gym.Prepare()
	gym.Router.Options("/requests/options", "IndexController@Options")
	gym.Router.RegisterController(&IndexController{})
	go gym.OpenAt(port)
	hello := &HelloFormat{}
	baseUri = fmt.Sprintf("http://localhost:%v", port)
	port++
	err := getJson(OPTIONSMethod, baseUri+"/requests/options", hello)
	if err != nil {
		t.Error("resp error")
	}
	if hello.Hello != "World" {
		t.Error("resp is not equal as expected")
	}
}

func TestPatch(t *testing.T) {
	var gym = new(Gym)
	gym.Prepare()
	gym.Router.Patch("/requests/patch", "IndexController@Patch")
	gym.Router.RegisterController(&IndexController{})
	go gym.OpenAt(port)
	hello := &HelloFormat{}
	baseUri = fmt.Sprintf("http://localhost:%v", port)
	port++
	err := getJson(PATCHMethod, baseUri+"/requests/patch", hello)
	if err != nil {
		t.Error("resp error")
	}
	if hello.Hello != "World" {
		t.Error("resp is not equal as expected")
	}
}

func TestRequestWithQuery(t *testing.T) {
	var gym = new(Gym)
	gym.Prepare()
	gym.Router.Get("/requests/form-method/query", "IndexController@QueryForm")
	gym.Router.RegisterController(&IndexController{})
	go gym.OpenAt(port)
	baseUri = fmt.Sprintf("http://localhost:%v", port)
	port++
	var requestQuery = url.Values{"api_key": {"gogym"}, "foo": {"bar&baz"}}
	request, _ := http.NewRequest("GET", baseUri+"/requests/form-method/query", nil)
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
	var gym = new(Gym)
	gym.Prepare()
	gym.Router.Post("/requests/form-method/form", "IndexController@PostForm")
	gym.Router.RegisterController(&IndexController{})
	go gym.OpenAt(port)
	baseUri = fmt.Sprintf("http://localhost:%v", port)
	port++
	requestForm := url.Values{"foo": {"bar", "baz"}}
	myClient.PostForm(baseUri+"/requests/form-method/form", requestForm)
	if !reflect.DeepEqual(formTest, requestForm) {
		t.Error("received form is not same as requested form")
	}
}

func TestHeader(t *testing.T) {
	var gym = new(Gym)
	gym.Prepare()
	gym.Router.Get("/requests/headers", "IndexController@SetHeaders")
	gym.Router.RegisterController(&IndexController{})
	go gym.OpenAt(port)
	baseUri = fmt.Sprintf("http://localhost:%v", port)
	port++
	r, err := myClient.Get(baseUri + "/requests/headers")
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
