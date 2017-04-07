package GoGym

import (
	"reflect"
	"testing"
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"io"
)

func TestRouter_Prepare(t *testing.T) {
	gym := Gym{}
	router := Router{}
	router.Prepare(&gym)
	expectedRouter := Router{App: &gym, ControllerRegistry: map[string]interface{}{}, MethodVerbs: []string{GETMethod, POSTMethod, PUTMethod, PATCHMethod, DELETEMethod, OPTIONSMethod}}
	if !reflect.DeepEqual(expectedRouter, router) {
		t.Error("Failed to prepare router")
	}
}

func TestRouter_InjectServiceContainer(t *testing.T) {
	gym := Gym{}
	router := Router{}
	router.InjectServiceContainer(&gym)
	expectedRouter := Router{App: &gym}
	if !reflect.DeepEqual(expectedRouter, router) {
		t.Error("Failed to inject service container")
	}
}

func TestRouter_GetServiceContainer(t *testing.T) {
	gym := Gym{}
	router := Router{}
	router.InjectServiceContainer(&gym)
	container := router.GetServiceContainer()
	if !reflect.DeepEqual(&gym, container) {
		t.Error("Failed to get service container")
	}
}

func TestRouter_IsActionLegal(t *testing.T) {
	router := Router{}
	action1 := "IndexController#Foo"
	result1 := router.IsActionLegal(action1)
	if result1 {
		t.Error("failed to check leagal action")
	}
	action2 := "IndexController@Foo"
	result2 := router.IsActionLegal(action2)
	if !result2 {
		t.Error("failed to check leagal action")
	}
}

func TestRouter_NewRoute(t *testing.T) {
	router := Router{}
	uri := "/foo/{bar}"
	methods := []string{GETMethod, POSTMethod}
	action := "IndexController@Index"
	router.NewRoute(uri, methods, action)
	var comp Compiled
	expTk := []Token{}
	tk1 := Token{Var: "foo", Name: "foo", Value: "foo", IsParam: false}
	tk2 := Token{Var: "{bar}", Name: "bar", IsParam: true}
	expTk = append(expTk, tk1, tk2)
	regexp := "\\/foo\\/(\\w+)$"
	comp.Tokens = expTk
	comp.RegExp = regexp
	rt := Route{uri: uri, methods: methods, action: action, compiled: comp}
	var expected []Route
	expected = append(expected, rt)
	if !reflect.DeepEqual(expected, router.RouteCollection) {
		t.Error("routes are not same as expected")
	}
}

func TestRouter_Get(t *testing.T) {
	router := Router{}
	router.Get("/foo", "IndexController@Get")
	expected := Router{}
	expected.NewRoute("/foo", []string{GETMethod}, "IndexController@Get")
	if !reflect.DeepEqual(router.RouteCollection, expected.RouteCollection) {
		t.Error("Get method is not working as expected")
	}
}

func TestRouter_Post(t *testing.T) {
	router := Router{}
	router.Post("/foo", "IndexController@Get")
	expected := Router{}
	expected.NewRoute("/foo", []string{POSTMethod}, "IndexController@Get")
	if !reflect.DeepEqual(router.RouteCollection, expected.RouteCollection) {
		t.Error("Post method is not working as expected")
	}
}

func TestRouter_Put(t *testing.T) {
	router := Router{}
	router.Put("/foo", "IndexController@Get")
	expected := Router{}
	expected.NewRoute("/foo", []string{PUTMethod}, "IndexController@Get")
	if !reflect.DeepEqual(router.RouteCollection, expected.RouteCollection) {
		t.Error("Put method is not working as expected")
	}
}

func TestRouter_Patch(t *testing.T) {
	router := Router{}
	router.Patch("/foo", "IndexController@Get")
	expected := Router{}
	expected.NewRoute("/foo", []string{PATCHMethod}, "IndexController@Get")
	if !reflect.DeepEqual(router.RouteCollection, expected.RouteCollection) {
		t.Error("Patch method is not working as expected")
	}
}

func TestRouter_Options(t *testing.T) {
	router := Router{}
	router.Options("/foo", "IndexController@Get")
	expected := Router{}
	expected.NewRoute("/foo", []string{OPTIONSMethod}, "IndexController@Get")
	if !reflect.DeepEqual(router.RouteCollection, expected.RouteCollection) {
		t.Error("Options method is not working as expected")
	}
}

func TestRouter_Delete(t *testing.T) {
	router := Router{}
	router.Delete("/foo", "IndexController@Get")
	expected := Router{}
	expected.NewRoute("/foo", []string{DELETEMethod}, "IndexController@Get")
	if !reflect.DeepEqual(router.RouteCollection, expected.RouteCollection) {
		t.Error("Delete method is not working as expected")
	}
}

func TestRouter_ServeHTTP_HTTP200(t *testing.T) {
	gym := Gym{}
	gym.Prepare()
	gym.Router.RegisterController(&FooController{})
	gym.Router.Get("/foo", "FooController@Index")
	go gym.OpenAt(2000)
	uri := "http://localhost:2000/foo"
	response := FooResponse{}
	GetJson("GET", uri, &response)
	if response.Msg != "Hello World" {
		t.Error("The result is not same as expected")
	}

}

func TestRouter_ServeHTTP_HTTP405(t *testing.T) {
	gym := Gym{}
	gym.Prepare()
	gym.Router.RegisterController(&FooController{})
	gym.Router.Get("/foo", "FooController@Index")
	go gym.OpenAt(2001)
	errRsp := ErrResponse{}
	var client = &http.Client{Timeout: 10 * time.Second}
	request, _ := http.NewRequest("POST", "http://localhost:2001/foo", strings.NewReader(""))
	rsp, err := client.Do(request)
	if err != nil {
		t.Errorf("Response error: %s", err)
	}
	if rsp.StatusCode != 405 {
		t.Error("The status code is not same as expected")
	}
	json.NewDecoder(rsp.Body).Decode(&errRsp)
	defer rsp.Body.Close()
	if errRsp.Err != "Method not allowed" {
		t.Error("The error info is not same as expected")
	}
}

func TestRouter_FindRoute(t *testing.T) {
	router := Router{}
	router.NewRoute("/foo/{bar}", []string{GETMethod}, "IndexController@Get")
	router.NewRoute("/foo/{bar}", []string{POSTMethod}, "IndexController@Post")
	router.NewRoute("/foo/bar/baz", []string{DELETEMethod}, "IndexController@Delete")
	var route Route
	route.uri = "/foo/{bar}"
	route.methods = []string{GETMethod}
	route.action = "IndexController@Get"
	route.extractTokens(route.uri)
	route.compile(route.uri)
	for k, tk := range route.compiled.Tokens {
		if tk.Name == "bar" {
			route.compiled.Tokens[k].Value = "yeah"
		}
	}
	var route1 Route
	route1.uri = "/foo/{bar}"
	route1.methods = []string{POSTMethod}
	route1.action = "IndexController@Post"
	route1.extractTokens(route.uri)
	route1.compile(route.uri)
	var expected []Route
	expected = append(expected, route, route1)
	for k, tk := range route1.compiled.Tokens {
		if tk.Name == "bar" {
			route1.compiled.Tokens[k].Value = "yeah"
		}
	}
	if !reflect.DeepEqual(expected, router.FindRoute("/foo/yeah")) {
		t.Error("failed to find routes")
	}
}

func TestRouter_Handle(t *testing.T) {
	router := Router{}
	router.App = new(Gym)
	router.ControllerRegistry = make(map[string]interface{})
	router.RegisterController(new(FooController))
	reader := strings.NewReader("")
	request, err := http.NewRequest("GET", "/foo", io.LimitReader(reader, 90))
	if err != nil {
		t.Error("Error when creating new request.")
	}
	router.NewRoute("/foo", []string{GETMethod}, "FooController@Bar")
	route := router.RouteCollection[0]
	router.Handle(route, RW{}, request)
	if Bar.Msg != "Hello Bar" {
		t.Error("Something went wrong when handling the request")
	}
}

func TestRouter_RegisterController(t *testing.T) {
	type IndexController struct{}
	controller := &IndexController{}
	router := Router{ControllerRegistry: make(map[string]interface{})}
	router.RegisterController(controller)
	expected := map[string]interface{}{"*IndexController": controller}
	if !reflect.DeepEqual(expected, router.ControllerRegistry) {
		t.Error("failed to register controller")
	}
}

type FooController struct{}

type FooResponse struct {
	Msg string
}

func (f FooController) Index(g *Gym) {
	g.Response.JsonResponse(map[string]string{"msg": "Hello World"}, 200, http.Header{})
}

var Bar FooResponse

func (f FooController) Bar(g *Gym) {
	Bar.Msg = "Hello Bar"
}

type ErrResponse struct {
	Err string
}

type RW struct{}

func (rw RW) Header() http.Header {
	return http.Header{}
}

func (rw RW) Write([]byte) (int, error) {
	return 0, nil
}

func (rw RW) WriteHeader(int) {

}
