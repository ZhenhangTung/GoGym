package GoGym

import (
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestRequest_Prepare(t *testing.T) {
	gym := new(Gym)
	request := Request{}
	request.Prepare(gym)
	if !reflect.DeepEqual(request.App, gym) {
		t.Error("Failed to prepare request")
	}
	if !reflect.DeepEqual(request.PathVar, make(map[string]string)) {
		t.Error("Failed to prepare request")
	}
}

func TestRequest_InjectServiceContainer(t *testing.T) {
	gym := new(Gym)
	request := Request{}
	request.InjectServiceContainer(gym)
	if !reflect.DeepEqual(request.App, gym) {
		t.Error("Failed to inject service container")
	}
}

func TestRequest_GetServiceContainer(t *testing.T) {
	gym := new(Gym)
	request := Request{}
	request.InjectServiceContainer(gym)
	g := request.GetServiceContainer()
	if !reflect.DeepEqual(gym, g) {
		t.Error("Failed to get service container")
	}
}

func TestRequest_BindPathVar(t *testing.T) {
	request := Request{}
	request.PathVar = make(map[string]string)
	var tokens []Token
	tk1 := Token{Name: "Foo", Value: "Foo", IsParam: false}
	tk2 := Token{Name: "Bar", Value: "yes", IsParam: true}
	tokens = append(tokens, tk1, tk2)
	request.BindPathVar(tokens)
	expected := map[string]string{"Bar": "yes"}
	if !reflect.DeepEqual(expected, request.PathVar) {
		t.Error("Something went wrong when binding path var")
	}
}

func TestRequest_Accept(t *testing.T) {
	gym := new(Gym)
	gym.Prepare()
	go gym.OpenAt(2010)
	httpRequest, _ := http.NewRequest("GET", "http://localhost:2010", strings.NewReader(""))
	httpRequest.Header.Add("Foo", "Bar")
	q := httpRequest.URL.Query()
	q.Add("api_key", "gogym")
	httpRequest.URL.RawQuery = q.Encode()
	_, err := myClient.Do(httpRequest)
	if err != nil {
		t.Error(err)
	}
	query := url.Values{}
	query["api_key"] = []string{"gogym"}
	if !reflect.DeepEqual(query, gym.Request.Query) {
		t.Error("Error when accepting query form")
	}
	if gym.Request.Method != "GET" {
		t.Error("Error when accepting request method")
	}
	if gym.Request.Header.Get("Foo") != "Bar" {
		t.Error("Error when accepting request header")
	}
}

func TestRequest_Accept_PostForm(t *testing.T) {
	gym := new(Gym)
	gym.Prepare()
	go gym.OpenAt(2011)
	requestForm := url.Values{"foo": {"bar", "baz"}}
	myClient.PostForm("http://localhost:2011/requests/form-method/form", requestForm)
	if !reflect.DeepEqual(requestForm, gym.Request.Form) {
		t.Error("Error when accepting request form")
	}
}

