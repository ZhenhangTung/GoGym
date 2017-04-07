package GoGym

import (
	"reflect"
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

}
