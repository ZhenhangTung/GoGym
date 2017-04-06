package GoGym

import (
	"reflect"
	"testing"
	//"fmt"
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
	rt := Route{Uri: uri, Methods: methods, Action: action, Compiled: comp}
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

func TestRouter_ServeHTTP(t *testing.T) {

}

func TestRouter_FindRoute(t *testing.T) {
	router := Router{}
	router.NewRoute("/foo/{bar}", []string{GETMethod}, "IndexController@Get")
	router.NewRoute("/foo/{bar}", []string{POSTMethod}, "IndexController@Post")
	router.NewRoute("/foo/bar/baz", []string{DELETEMethod}, "IndexController@Delete")
	var route Route
	route.Uri = "/foo/{bar}"
	route.Methods = []string{GETMethod}
	route.Action = "IndexController@Get"
	route.ExtractTokens(route.Uri)
	route.Compile(route.Uri)
	for k, tk := range route.Compiled.Tokens {
		if tk.Name == "bar" {
			route.Compiled.Tokens[k].Value = "yeah"
		}
	}
	var route1 Route
	route1.Uri = "/foo/{bar}"
	route1.Methods = []string{POSTMethod}
	route1.Action = "IndexController@Post"
	route1.ExtractTokens(route.Uri)
	route1.Compile(route.Uri)
	var expected []Route
	expected = append(expected, route, route1)
	for k, tk := range route1.Compiled.Tokens {
		if tk.Name == "bar" {
			route1.Compiled.Tokens[k].Value = "yeah"
		}
	}
	if !reflect.DeepEqual(expected, router.FindRoute("/foo/yeah")) {
		t.Error("failed to find routes")
	}
}

func TestRouter_Handle(t *testing.T) {

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
