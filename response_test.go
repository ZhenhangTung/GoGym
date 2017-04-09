package GoGym

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestResponse_Prepare(t *testing.T) {
	resp := Response{}
	gym := new(Gym)
	resp.Prepare(gym)
	if !reflect.DeepEqual(gym, resp.GetServiceContainer()) {
		t.Error("Error when preparing response service")
	}
}

func TestResponse_InjectServiceContainer(t *testing.T) {
	resp := Response{}
	gym := new(Gym)
	resp.InjectServiceContainer(gym)
	if !reflect.DeepEqual(gym, resp.app) {
		t.Error("Error when preparing response service")
	}
}

func TestResponse_JsonResponse(t *testing.T) {
	resp := Response{}
	//resp.Header
	r := map[string]string{"hello": "world"}
	header := http.Header{}
	header.Add("Foo", "Bar")
	resp.JsonResponse(r, 200, header)
	if jsn, _ := json.Marshal(r); !reflect.DeepEqual(jsn, resp.Response) {
		t.Error("Error when marshaling json response")
	}
	if resp.StatusCode != 200 {
		t.Error("Error when setting http status code")
	}
	var respHeader http.Header
	respHeader = http.Header{}
	respHeader.Add("Content-Type", MIME_APP_JSON)
	respHeader.Add("Foo", "Bar")
	//fmt.Println(respHeader)
	if !reflect.DeepEqual(respHeader.Get("Foo"), resp.Header.Get("Foo")) {
		t.Error("Error when setting http header")
	}
	if !reflect.DeepEqual(respHeader.Get("Content-Type"), resp.Header.Get("Content-Type")) {
		t.Error("Error when setting http header")
	}
}

func TestResponse_GetServiceContainer(t *testing.T) {
	resp := Response{}
	gym := new(Gym)
	resp.InjectServiceContainer(gym)
	if !reflect.DeepEqual(gym, resp.GetServiceContainer()) {
		t.Error("Error when preparing response service")
	}
}
