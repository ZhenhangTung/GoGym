![](http://tongzhenhang.me/wp-content/uploads/2017/03/dumbbell-1.png)
<div style="font-size: 10px;">Icons made by <a href="http://www.flaticon.com/authors/vectors-market" title="Vectors Market">Vectors Market</a> from <a href="http://www.flaticon.com" title="Flaticon">www.flaticon.com</a> is licensed by <a href="http://creativecommons.org/licenses/by/3.0/" title="Creative Commons BY 3.0" target="_blank">CC 3.0 BY</a></div>
# GoGym
---


[![Build Status](https://travis-ci.org/ZhenhangTung/GoGym.svg?branch=master)](https://travis-ci.org/ZhenhangTung/GoGym)


```GoGym ``` is a micro-framework for building RESTful APIs, which is written in ```Golang```. It is inspired by an artisan framework [Laravel](https://laravel.com/).

## Import Package
* Install the package from the command line: 

	```bash
	$ go get github.com/ZhenhangTung/GoGym
	```

## Steps of implementation
1. Define a controller

	```go
	type IndexController struct {
	}
	```
2. Define an action in the controller

	```go
	func (IndexController *IndexController) Index(values url.Values, headers http.Header) (statusCode int, response interface{}) {
		return 200, map[string]string{"hello": "world"}
	}
	```
3. Prepare the service
	
	```go
	var apiService = GoGym.Prepare()
	```

4. Set your path with Controller and Action


	```go
	apiService.Get("/", "IndexController@Index") // GET Method
	apiService.Get("users", "IndexController@GetUsers")
	apiService.Post("index", "IndexController@Post") // Post Method
	apiService.Put("index", "IndexController@Post") // Post Method
	apiService.Delete("index", "IndexController@Delete") // Delete Method
	apiService.Options("index", "IndexController@Options") // Options Method
	apiService.Patch("index", "IndexController@Patch") // Patch Method
	```
5. Register your controller
	* Register a single controller

	
	```go
	apiService.RegisterController(&FooController{})
	```
	* Register mutiple controllers

	
	```go
	controllers := []interface{}{&IndexController{}}
	apiService.RegisterControllers(controllers)
	```
6. Start to Serve


	``` go
	apiService.Serve(3000)
	```

## How to get request
In your controller's action, for example as below:

```go
func (IndexController *IndexController) QueryForm(request map[string]url.Values, headers http.Header) (statusCode int, response interface{}) {
	query := request["query"]
	form := request["form"]
}
```
* ```request["query"]``` is for getting query string for all requests
* ```request["form"]``` is for getting form when requests are ```POST```, ```PUT```, and ```PATCH``` requests



## Code Example

```go
package main

import (
    "net/url"
    "net/http"
    "github.com/ZhenhangTung/GoGym"
)

type IndexController struct {
}

func (IndexController *IndexController) Index(request map[string]url.Values, headers http.Header) (statusCode int, response interface{}) {
	return 200, map[string]string{"hello": "world"}
}

type BarController struct {
}

func (*BarController) Bar(request map[string]url.Values, headers http.Header) (statusCode int, response interface{}) {
	return 200, map[string]string{"GoTo": "Bar"}
}

func main() {
	var apiService = GoGym.Prepare()
	apiService.Get("index", "IndexController@Index")
	apiService.Post("bar", "BarController@Bar")
	controllers := []interface{}{&IndexController{}}
	apiService.RegisterControllers(controllers)
	apiService.RegisterController(&BarController{})
	apiService.Serve(3000)
}
```

## Running result
* Test GET Request

	```bash
	$ curl localhost:3000/index
	{"hello":"world"}
	```

* Test POST Request

	```bash
	$ curl -H "Content-Type: application/json" -d '{"hello":"world"}' http://localhost:3000/bar
	{"GoTo":"Bar"}
	```

## Notice
```GoGym``` now is still in development, features in development now:

1. <strike>Unit tests</strike>
2. Some optimization for data structure
3. Error Handling with detail information
4. User could set his own headers
5. <strike>Handle input form easily<strike>


## License

`GoGym` is released under the [MIT License](http://opensource.org/licenses/MIT).