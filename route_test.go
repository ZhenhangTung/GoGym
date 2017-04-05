package GoGym

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRoute_ExtractTokens_WithParams(t *testing.T) {
	var route Route
	route.Uri = "/foo/{bar}/baz"
	route.ExtractTokens(route.Uri)
	expected := []Token{}
	tk1 := Token{Var: "foo", Name: "foo", Value: "foo", IsParam: false}
	tk2 := Token{Var: "{bar}", Name: "bar", IsParam: true}
	tk3 := Token{Var: "baz", Name: "baz", Value: "baz", IsParam: false}
	expected = append(expected, tk1, tk2, tk3)
	if !reflect.DeepEqual(route.Compiled.Tokens, expected) {
		t.Error("tokens are not same as expected")
	}
}

func TestRoute_ExtractTokens_NoParam(t *testing.T) {
	var route Route
	route.Uri = "/foo/bar"
	route.ExtractTokens(route.Uri)
	expected := []Token{}
	tk1 := Token{Var: "foo", Name: "foo", Value: "foo", IsParam: false}
	tk2 := Token{Var: "bar", Name: "bar", Value: "bar", IsParam: false}
	expected = append(expected, tk1, tk2)
	if !reflect.DeepEqual(route.Compiled.Tokens, expected) {
		t.Error("tokens are not same as expected")
	}
}

func TestRoute_ExtractTokens_RouteIsConsistedOfParams(t *testing.T) {
	var route Route
	route.Uri = "/{foo}/{bar}"
	route.ExtractTokens(route.Uri)
	expected := []Token{}
	tk1 := Token{Var: "{foo}", Name: "foo", IsParam: true}
	tk2 := Token{Var: "{bar}", Name: "bar", IsParam: true}
	expected = append(expected, tk1, tk2)
	if !reflect.DeepEqual(route.Compiled.Tokens, expected) {
		t.Error("tokens are not same as expected")
	}
}

func TestRoute_Compile_WithParams(t *testing.T) {
	var route Route
	route.Uri = "/foo/{bar}/baz"
	route.ExtractTokens(route.Uri)
	route.Compile(route.Uri)
	expectedRegExp := "\\/foo\\/(\\w+)\\/baz$"
	if expectedRegExp != route.Compiled.RegExp {
		t.Error("regexp is not same as expected")
	}
}

func TestRoute_Compile_NoParam(t *testing.T) {
	var route Route
	route.Uri = "/foo/bar"
	route.ExtractTokens(route.Uri)
	route.Compile(route.Uri)
	expectedRegExp := "\\/foo\\/bar$"
	if expectedRegExp != route.Compiled.RegExp {
		t.Error("regexp is not same as expected")
	}
}

func TestRoute_Compile_RouteIsConsistedOfParams(t *testing.T) {
	var route Route
	route.Uri = "/{foo}/{bar}/{baz}"
	route.ExtractTokens(route.Uri)
	route.Compile(route.Uri)
	expectedRegExp := "\\/(\\w+)\\/(\\w+)\\/(\\w+)$"
	if expectedRegExp != route.Compiled.RegExp {
		t.Error("regexp is not same as expected")
	}
}

func TestRoute_Match_NoParam(t *testing.T) {
	var route Route
	route.Uri = "/foo/bar/baz"
	route.ExtractTokens(route.Uri)
	route.Compile(route.Uri)
	var request string
	var match bool
	request = "/foo/bar/baz"
	match = route.Match(request)
	if match != true {
		t.Error("match result is not same as expected")
	}
	request = "/foo/bar"
	match = route.Match(request)
	if match == true {
		t.Error("match result is not same as expected")
	}
	request = "/foo/bar/baz/hey"
	match = route.Match(request)
	if match == true {
		t.Error("match result is not same as expected")
	}
}

func TestRoute_Match_WithParams(t *testing.T) {
	var route Route
	route.Uri = "/{foo}/bar/baz"
	route.ExtractTokens(route.Uri)
	route.Compile(route.Uri)
	var request string
	var match bool
	request = "/test/bar/baz"
	match = route.Match(request)
	if match != true {
		t.Error("match result is not same as expected")
	}
	request = "/test/bar"
	match = route.Match(request)
	if match == true {
		t.Error("match result is not same as expected")
	}
	request = "/test/bar/baz/hey"
	match = route.Match(request)
	if match == true {
		t.Error("match result is not same as expected")
	}
}

func TestRoute_Match_RouteIsConsistedOfParams(t *testing.T) {
	var route Route
	route.Uri = "/{foo}/{bar}/{baz}"
	route.ExtractTokens(route.Uri)
	route.Compile(route.Uri)
	var request string
	var match bool
	request = "/music/jazz/yeah"
	match = route.Match(request)
	if match != true {
		t.Error("match result is not same as expected")
	}
	request = "/test/bar"
	match = route.Match(request)
	if match == true {
		t.Error("match result is not same as expected")
	}
	request = "/test/bar/baz/hey"
	match = route.Match(request)
	if match == true {
		fmt.Println("re", route.Compiled.RegExp)
		t.Error("match result is not same as expected")
	}
}

func TestRoute_AssignValuesToTokens(t *testing.T) {
	var route Route
	route.Uri = "/{foo}/bar/{test1}/baz/{test}"
	route.ExtractTokens(route.Uri)
	route.Compile(route.Uri)
	var request string
	request = "/test/bar/hello/baz/world"
	route.AssignValuesToTokens(request)
	expected := []Token{}
	tk1 := Token{Var: "{foo}", Name: "foo", Value: "test", IsParam: true}
	tk2 := Token{Var: "bar", Name: "bar", Value: "bar", IsParam: false}
	tk3 := Token{Var: "{test1}", Name: "test1", Value: "hello", IsParam: true}
	tk4 := Token{Var: "baz", Name: "baz", Value: "baz", IsParam: false}
	tk5 := Token{Var: "{test}", Name: "test", Value: "world", IsParam: true}
	expected = append(expected, tk1, tk2, tk3, tk4, tk5)
	if !reflect.DeepEqual(route.Compiled.Tokens, expected) {
		t.Error("tokens' value are not same as expected")
	}
}
