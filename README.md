<img src="data:image/svg+xml;utf8;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iaXNvLTg4NTktMSI/Pgo8IS0tIEdlbmVyYXRvcjogQWRvYmUgSWxsdXN0cmF0b3IgMTkuMC4wLCBTVkcgRXhwb3J0IFBsdWctSW4gLiBTVkcgVmVyc2lvbjogNi4wMCBCdWlsZCAwKSAgLS0+CjxzdmcgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgdmVyc2lvbj0iMS4xIiBpZD0iTGF5ZXJfMSIgeD0iMHB4IiB5PSIwcHgiIHZpZXdCb3g9IjAgMCA1MDQgNTA0IiBzdHlsZT0iZW5hYmxlLWJhY2tncm91bmQ6bmV3IDAgMCA1MDQgNTA0OyIgeG1sOnNwYWNlPSJwcmVzZXJ2ZSIgd2lkdGg9IjUxMnB4IiBoZWlnaHQ9IjUxMnB4Ij4KPGNpcmNsZSBzdHlsZT0iZmlsbDojODREQkZGOyIgY3g9IjI1MiIgY3k9IjI1MiIgcj0iMjUyIi8+CjxwYXRoIHN0eWxlPSJmaWxsOiMzMjRBNUU7IiBkPSJNNDM3LjksMjY0SDY2LjJjLTYuNiwwLTEyLTUuNC0xMi0xMmwwLDBjMC02LjYsNS40LTEyLDEyLTEyaDM3MS43YzYuNiwwLDEyLDUuNCwxMiwxMmwwLDAgIEM0NDkuOSwyNTguNiw0NDQuNSwyNjQsNDM3LjksMjY0eiIvPgo8cGF0aCBzdHlsZT0iZmlsbDojRkZGRkZGOyIgZD0iTTM1NywzMzkuOUwzNTcsMzM5LjljLTE0LjUsMC0yNi4zLTExLjgtMjYuMy0yNi4zVjE5MC40YzAtMTQuNSwxMS44LTI2LjMsMjYuMy0yNi4zbDAsMCAgYzE0LjUsMCwyNi4zLDExLjgsMjYuMywyNi4zdjEyMy4yQzM4My4zLDMyOC4xLDM3MS41LDMzOS45LDM1NywzMzkuOXoiLz4KPHBhdGggc3R5bGU9ImZpbGw6I0U2RTlFRTsiIGQ9Ik00MDkuNiwzMTUuNUw0MDkuNiwzMTUuNWMtMTQuNSwwLTI2LjMtMTEuOC0yNi4zLTI2LjN2LTc0LjRjMC0xNC41LDExLjgtMjYuMywyNi4zLTI2LjNsMCwwICBjMTQuNSwwLDI2LjMsMTEuOCwyNi4zLDI2LjN2NzQuNEM0MzUuOSwzMDMuNyw0MjQuMSwzMTUuNSw0MDkuNiwzMTUuNXoiLz4KPHBhdGggc3R5bGU9ImZpbGw6I0ZGRkZGRjsiIGQ9Ik0xNDcsMTY0LjFMMTQ3LDE2NC4xYzE0LjUsMCwyNi4zLDExLjgsMjYuMywyNi4zdjEyMy4yYzAsMTQuNS0xMS44LDI2LjMtMjYuMywyNi4zbDAsMCAgYy0xNC41LDAtMjYuMy0xMS44LTI2LjMtMjYuM1YxOTAuNEMxMjAuNywxNzUuOSwxMzIuNSwxNjQuMSwxNDcsMTY0LjF6Ii8+CjxwYXRoIHN0eWxlPSJmaWxsOiNFNkU5RUU7IiBkPSJNOTQuNCwxODguNUw5NC40LDE4OC41YzE0LjUsMCwyNi4zLDExLjgsMjYuMywyNi4zdjc0LjRjMCwxNC41LTExLjgsMjYuMy0yNi4zLDI2LjNsMCwwICBjLTE0LjUsMC0yNi4zLTExLjgtMjYuMy0yNi4zdi03NC40QzY4LjEsMjAwLjMsNzkuOSwxODguNSw5NC40LDE4OC41eiIvPgo8Zz4KPC9nPgo8Zz4KPC9nPgo8Zz4KPC9nPgo8Zz4KPC9nPgo8Zz4KPC9nPgo8Zz4KPC9nPgo8Zz4KPC9nPgo8Zz4KPC9nPgo8Zz4KPC9nPgo8Zz4KPC9nPgo8Zz4KPC9nPgo8Zz4KPC9nPgo8Zz4KPC9nPgo8Zz4KPC9nPgo8Zz4KPC9nPgo8L3N2Zz4K" style="width:200px; height: 200px;"/>
<div style="font-size: 10px;">Icons made by <a href="http://www.flaticon.com/authors/vectors-market" title="Vectors Market">Vectors Market</a> from <a href="http://www.flaticon.com" title="Flaticon">www.flaticon.com</a> is licensed by <a href="http://creativecommons.org/licenses/by/3.0/" title="Creative Commons BY 3.0" target="_blank">CC 3.0 BY</a></div>
# GoGym

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

func (IndexController *IndexController) Index(values url.Values, headers http.Header) (statusCode int, response interface{}) {
	return 200, map[string]string{"hello": "world"}
}

type BarController struct {
}

func (*BarController) Bar(values url.Values, headers http.Header) (statusCode int, response interface{}) {
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
```GoGym``` now is still in development, it needs:

1. Unit tests
2. Some optimization for data structure
3. Error Handling with detail information
4. User could set his own headers

## License

`GoGym` is released under the [MIT License](http://opensource.org/licenses/MIT).