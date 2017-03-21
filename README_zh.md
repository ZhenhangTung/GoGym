# GoGym  


[![Build Status](https://travis-ci.org/ZhenhangTung/GoGym.svg?branch=master)](https://travis-ci.org/ZhenhangTung/GoGym)

```GoGym```是一个用Golang构建的RESTful APIs的框架。它深受[Laravel](https://laravel.com/)启发，希望能帮助用户优雅快速地构建API服务。

![](http://tongzhenhang.me/wp-content/uploads/2017/03/GoGym_Logo_256.png)   
项目Icon由 @Beth Wardolf 设计制作

## 安装项目
* 通过命令行输入
	```bash
	$ go get github.com/ZhenhangTung/GoGym
	```

## 如何使用 ```GoGym```

### 1. 定义控制器和方法

```
type HelloController struct {
}

func (h *HelloController) SayHello(api *GoGym.Gym) {
    // Your logic goes there
}
```

### 2. 为开张你自己的Gym做准备
```
var gym = new(GoGym.Gym)
gym.Prepare()
```

### 3. 注册你定义的控制器，让它成为Gym家族的一员
```
gym.Router.RegisterController(&HelloController{})
```

### 4. 将控制器和对应的方法与请求的方法和路径绑定
```
gym.Router.Get("/", "HelloController@SayHello")
```

### 5. 祝贺！你的Gym已经可以开张了！
```
gym.OpenAt(3000)
```

## 核心服务

### ```Gym```
```Gym``` 是一个服务容器:   

* ```RegisterService(name string, service GymService)``` : 将用户自定义的服务注册到容器中
* ```RegisterServices(services map[string]GymService)``` : 将用户自定义的多个服务注册到容器中
* ```GetService(name string) GymService``` : 从```Gym```容器中获得服务
* ```CallServiceMethod(service GymService, method string, param []interface{}) []reflect.Value``` : 用来调用服务的方法
* ```Prepare() *Gym``` : 用来准备```Gym```
* ```OpenAt(port int)``` : 在用户定义的端口启动服务

### ```Router```：路由
* ```Get(path, controllerWithActionString string)``` : 将控制器和其方法注册为负责接收路由为```path ```的GET请求
* ```Post(path, controllerWithActionString string)``` : 将控制器和其方法注册为负责接收路由为```path ```的POST请求
* ```Put(path, controllerWithActionString string)``` : 将控制器和其方法注册为负责接收路由为```path ```的PUT请求
* ```Patch(path, controllerWithActionString string)``` : 将控制器和其方法注册为负责接收路由为```path ```的PATCH请求
* ```Options(path, controllerWithActionString string)``` : 将控制器和其方法注册为负责接收路由为```path ```的OPTIONS请求
* ```Delete(path, controllerWithActionString string)``` : 将控制器和其方法注册为负责接收路由为```path ```的DELETE请求
* ```RegisterController``` : 负责注册控制器
* ```RegisterControllers``` : 负责注册多个控制器

### ```Request``` ： 请求
* ```Request.Method```: 返回HTTP请求的方法
* ```Request.Header```: 返回HTTP请求的header
* ```Request.Query``` : 返回HTTP请求时候所带的的Query
* ```Request.Form``` : 返回HTTP请求时候所带的表单

### ```Response```：响应
* ```JsonResponse(resp interface{}, statusCode int, header http.Header)```: 接受结果，HTTP状态码和HTTP header，生成JSON响应

### ```Helpers```：辅助方法
* ```GetType(value interface{}) string``` : 获取value的类型
* ```CallServiceMethodWithReflect(g GymService, method string, param []interface{}) []reflect.Value``` : 这个方法是封装了调用用户定义service中的method，用户也可以选择不使用这个方法，自己实现

## 如何注册自己定义的服务到容器中?
1. 接入```GymService```这个接口，并且实现```GymService```下面的方法
2. 通过使用```RegisterService```或者```RegisterServices```，将你的服务注册到```Gym```中
3. 通过```GetService```获取你的服务
4. 通过```CallMethod ```来调用服务内部的方法
5. 你可以自己实现```CallMethod ```或者在```CallMethod ```中使用```CallServiceMethodWithReflect()```这个预先封装好的方法



## 示例代码


```go
package main

import (
	"fmt"
	"github.com/ZhenhangTung/GoGym"
	"net/http"
	"reflect"
)

type HelloController struct {
}

func (h *HelloController) SayHello(api *GoGym.Gym) {
	method := api.Request.Method
	fmt.Println(method)
	api.Response.JsonResponse(map[string]string{"hello": "world"}, 200, http.Header{})
}

type FooService struct {
	boss *GoGym.Gym
}

func (f *FooService) Prepare(g *GoGym.Gym) {
	f.WhoIsYourBoss(g)
}

func (f *FooService) WhoIsYourBoss(g *GoGym.Gym) {
	f.boss = g
}

func (f *FooService) CallYourBoss() *GoGym.Gym {
	return f.boss
}

func (f *FooService) Test() {
	fmt.Println("oh yes")
}

func (f *FooService) CallMethod(method string, param []interface{}) []reflect.Value {
	r := GoGym.CallServiceMethodWithReflect(f, method, param)
	return r
}

func main() {
	var gym = new(GoGym.Gym)
	gym.Prepare()
	gym.Router.RegisterController(&HelloController{})
	gym.Router.Get("/", "HelloController@SayHello")
	gym.RegisterService("Foo", new(FooService))
	gym.GetService("Foo").CallMethod("Test", nil)
	gym.OpenAt(3000)
}


// Then open the http://localhost:3000 to see the result

```



## 想成为```GoGym```的一员

十分欢迎参与```GoGym```的开发
你可以通过如下的方式参与:

* [发布issue或者feedback](https://github.com/ZhenhangTung/GoGym/issues)
* 通过Pull Request来提交bug fix或者new feature
* 写文档或者修饰文档


## 项目贡献成员
[https://github.com/ZhenhangTung/GoGym/graphs/contributors](https://github.com/ZhenhangTung/GoGym/graphs/contributors)



## License

```GoGym```遵守[MIT License](http://opensource.org/licenses/MIT).

