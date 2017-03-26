


# GoGym


[![Build Status](https://travis-ci.org/ZhenhangTung/GoGym.svg?branch=master)](https://travis-ci.org/ZhenhangTung/GoGym)

[中文文档](./README_zh.md)


```GoGym ``` is a framework for building RESTful APIs, which is written in ```Golang```. It is inspired by an artisan framework [Laravel](https://laravel.com/).

![](http://tongzhenhang.me/wp-content/uploads/2017/03/GoGym_Logo_256.png)
Icon made by @Beth Wardolf

## Import Package
* Install the package from the command line:

	```bash
	$ go get github.com/ZhenhangTung/GoGym
	```

## How to use ```GoGym```

### 1. Define your Own Controllers And Actions

```
type HelloController struct {
}

func (h *HelloController) SayHello(api *GoGym.Gym) {
    // Your logic goes there
}
```

### 2. Do preparations for your Gym's opening
```
var gym = new(GoGym.Gym)
gym.Prepare()
```
### 3. Register your controller, and make it as a part of your Gym family
```
gym.Router.RegisterController(&HelloController{})
```

### 4. Be ready for getting requests. Let's smile and say hello. :)
```
gym.Router.Get("/", "HelloController@SayHello")
```

### 5. Now your Gym are opening. Let's cheer!
```
gym.OpenAt(3000)
```

## Core Services

### ```Gym```
```Gym``` is a service container :

* ```RegisterService(name string, service GymService)``` : RegisterService registers user's own service into service container
* ```RegisterServices(services map[string]GymService)``` : RegisterServices is a function registers mutiple services
* ```GetService(name string) GymService``` : GetService is a fucntion gets a service from ```Gym```
* ```CallServiceMethod(service GymService, method string, param []interface{}) []reflect.Value``` : CallServiceMethod is a function call a method of service
* ```Prepare() *Gym``` :  Prepare is a function prepares the service container
* ```OpenAt(port int)``` : OpenAt is a function which is used to serve the service

### ```Router```
* ```Get(path, controllerWithActionString string)``` : Get is a fucntion handles GET requests
* ```Post(path, controllerWithActionString string)``` : Post is a fucntion handles POST requests
* ```Put(path, controllerWithActionString string)``` : Put is a method handles PUT requests
* ```Patch(path, controllerWithActionString string)``` : Patch is a method handles PATCH requests
* ```Options(path, controllerWithActionString string)``` : Options is a method handles OPTIONS requests
* ```Delete(path, controllerWithActionString string)``` : Delete is a method handles DELETE requests
* ```RegisterController``` : RegisterControllers is a method registers a controller into controllerRegistry
* ```RegisterControllers``` : RegisterControllers is a method registers a struct of controllers into controllerRegistry

### ```Request```
* ```Request.Method```: It gets the method of the http request
* ```Request.Header```: It gets the header of the http request
* ```Request.Query``` : It parses query of the http request
* ```Request.Form``` : It parses request form of the http request

### ```Response```
* ```JsonResponse(resp interface{}, statusCode int, header http.Header)```: ```JsonResponse ``` accept response, status code and http header to generate http JSON response.

### ```Helpers```
* ```GetType(value interface{}) string``` : GetType is a function gets the type of value
* ```CallServiceMethodWithReflect(g GymService, method string, param []interface{}) []reflect.Value``` : CallServiceMethodWithReflect is a functon calls user's own service method

## Want to implement your own service?
1. Implement ```GymService``` interface
2. Pass your service into ```Gym``` using method ```RegisterService``` or ```RegisterServices```
3. Get your service using method ```GetService```
4. Call the service's method using ```CallMethod ```
5. You could write your own ```CallMethod ``` or use the helper function ```CallServiceMethodWithReflect()```



## Code Example


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


## Contribution

Your contribution to ```GoGym``` development is very welcomed!
You may contribute in the following ways:

* [Post issues and feedbacks](https://github.com/ZhenhangTung/GoGym/issues)
* Submit fixes, features via Pull Request
* Write/polish documentation


## Contributors
[https://github.com/ZhenhangTung/GoGym/graphs/contributors](https://github.com/ZhenhangTung/GoGym/graphs/contributors)



## License

`GoGym` is released under the [MIT License](http://opensource.org/licenses/MIT).
