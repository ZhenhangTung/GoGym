# 响应

## 返回JSON响应
 ```JsonResponse(resp interface{}, statusCode int, header http.Header)```: ```JsonResponse ```接收返回值，http状态码，还有http header来生成JSON响应。

```
type HelloController struct {
}

func (h *HelloController) SayHello(api *GoGym.Gym) {
	api.Response.JsonResponse(map[string]string{"msg": "Hello World!"}, 200, http.Header{})
}
```
