![](http://tongzhenhang.me/wp-content/uploads/2017/03/dumbbell-1.png)
<div style="font-size: 10px;">Icons made by <a href="http://www.flaticon.com/authors/vectors-market" title="Vectors Market">Vectors Market</a> from <a href="http://www.flaticon.com" title="Flaticon">www.flaticon.com</a> is licensed by <a href="http://creativecommons.org/licenses/by/3.0/" title="Creative Commons BY 3.0" target="_blank">CC 3.0 BY</a></div>


# GoGym  



[![Build Status](https://travis-ci.org/ZhenhangTung/GoGym.svg?branch=master)](https://travis-ci.org/ZhenhangTung/GoGym)


```GoGym ``` is a micro-framework for building RESTful APIs, which is written in ```Golang```. It is inspired by an artisan framework [Laravel](https://laravel.com/).

## Import Package
* Install the package from the command line: 

	```bash
	$ go get github.com/ZhenhangTung/GoGym
	```

## How to use ```GoGym```

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
* ```Options(path, controllerWithActionString string)``` : Options is a method handles Options requests
* ```Delete(path, controllerWithActionString string)``` : Delete is a method handles Delete requests
* ```RegisterController``` : RegisterControllers is a method registers a controller into controllerRegistry
* ```RegisterControllers``` : RegisterControllers is a method registers a struct of controllers into controllerRegistry

### ```Request``` 
* ```Request.Method```: It gets the method of the http request
* ```Request.Header```: It gets the header of the http request
* ```Request.Query``` : It parses query of the http request
* ```Request.Form``` : It parses request form of the http request

### ```Response```
* ```JsonResponse(resp interface{}, statusCode int, header http.Header)```: ```JsonResponse ``` accept response, status code and http header to generate http JSON response.


## Want to implement your own service?
1. Implement ```GymService``` interface
2. Pass your service into ```Gym``` using method ```RegisterService``` or ```RegisterServices```
3. Get your service using method ```GetService```
4. Call the service's method using ```CallServiceMethod```



## Code Example (How to get a Hello World)

```go
package main

import (
    "net/url"
    "net/http"
    "github.com/ZhenhangTung/GoGym"
)

type HelloController struct {
}

func (h *HelloController) SayHello(api *GoGym.Gym) {
	api.Response.JsonResponse(map[string]string{"hello": "world"}, 200, http.Header{})
}

func main() {
	var gym = new(GoGym.Gym)
	gym.Prepare()
	gym.Router.RegisterController(&HelloController{})
	gym.Router.Get("/", "HelloController@SayHello")
	gym.OpenAt(3000)
}

// Then open the http://localhost:3000 to see the result

```


## Notice
```GoGym``` now is still in development which means that it is still <b>unstable</b>, and it has changed a lot since the time when it was open source. I believe these changes are impressive and brings more flexibility to ```GoGym```.  ```v0.1``` is on schedule and would be comming soon.


## License

`GoGym` is released under the [MIT License](http://opensource.org/licenses/MIT).