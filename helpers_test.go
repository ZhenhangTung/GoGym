package GoGym

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHelpers_GetType(t *testing.T) {
	tp := Type{}
	type1 := GetType(tp)
	if type1 != "Type" {
		t.Error("Error when getting type")
	}
	type2 := GetType(&tp)
	if type2 != "*Type" {
		t.Error("Error when getting type")
	}
}

func TestHelpers_CallServiceMethodWithReflect(t *testing.T) {

}

func TestHelpers_GetJson(t *testing.T) {
	var gym = new(Gym)
	gym.Prepare()
	gym.Router.Get("/", "JsonController@GetJson")
	gym.Router.RegisterController(&JsonController{})
	go gym.OpenAt(2018)
	jsonRsp := &Json{}
	baseUri = fmt.Sprintf("http://localhost:%v", 2018)
	err := getJson(GETMethod, baseUri, jsonRsp)
	if err != nil {
		t.Error(err)
	}
	if jsonRsp.Json != "Yes" {
		t.Error("resp is not equal as expected")
	}
}

type Type struct {
}

type JsonController struct {
}

func (j *JsonController) GetJson(g *Gym) {
	rsp := map[string]string{"Json": "Yes"}
	g.Response.JsonResponse(rsp, 200, http.Header{})
}

type Json struct {
	Json string
}
