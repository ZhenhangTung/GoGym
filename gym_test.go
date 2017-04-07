package GoGym

import (
	"reflect"
	"testing"
	//"fmt"
)

func TestGym_Prepare(t *testing.T) {
	gym := new(Gym)
	gym.Prepare()
	g := new(Gym)
	g.services = make(map[string]GymService)
	g.Router = new(Router)
	g.Router.Prepare(g)
	g.Request = new(Request)
	g.Request.Prepare(g)
	g.Response = new(Response)
	g.Response.Prepare(g)
	if !reflect.DeepEqual(gym, g) {
		t.Error("Failed to prepare gym")
	}
}

func TestGym_BindService(t *testing.T) {
	gym := Gym{}
	req := Request{}
	gym.Prepare()
	gym.bindService("Request", &req)
	if !reflect.DeepEqual(&req, gym.GetService("Request")) {
		t.Error("Error when binding service")
	}
}

func TestGym_GetService(t *testing.T) {
	gym := Gym{}
	req := Request{}
	gym.Prepare()
	gym.bindService("Request", &req)
	if !reflect.DeepEqual(&req, gym.GetService("Request")) {
		t.Error("Error when binding service")
	}
}

func TestGym_RegisterService(t *testing.T) {
	gym := Gym{}
	req := Request{}
	gym.Prepare()
	gym.RegisterService("Request", &req)
	if !reflect.DeepEqual(&req, gym.GetService("Request")) {
		t.Error("Error when binding service")
	}
}

func TestGym_RegisterServices(t *testing.T) {
	gym := Gym{}
	req := Request{}
	resp := Response{}
	gym.Prepare()
	svcs := map[string]GymService{"Request": &req, "Response": &resp}
	gym.RegisterServices(svcs)
	if !reflect.DeepEqual(&req, gym.GetService("Request")) {
		t.Error("Error when binding services")
	}
	if !reflect.DeepEqual(&resp, gym.GetService("Response")) {
		t.Error("Error when binding services")
	}
}

func TestGym_OpenAt(t *testing.T) {
	gym := Gym{}
	gym.Prepare()
	go gym.OpenAt(2048)
	err := GetJson("GET", "http://localhost:2048", &map[string]string{})
	if err != nil {
		t.Error("Error when opening service on specific port")
	}
	e := GetJson("GET", "http://localhost:2049", &map[string]string{})
	if e == nil {
		t.Error("Error when opening service on specific port")
	}
}
