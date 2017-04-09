# Response

## Return JSON Response
 ```JsonResponse(resp interface{}, statusCode int, header http.Header)```: ```JsonResponse ``` accept response, status code and http header to generate http JSON response.

```
type HelloController struct {
}

func (h *HelloController) SayHello(api *GoGym.Gym) {
	api.Response.JsonResponse(map[string]string{"msg": "Hello World!"}, 200, http.Header{})
}
```
