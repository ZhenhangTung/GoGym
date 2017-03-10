# GoGym

```GoGym ``` is a micro-framework for building RESTful APIs, which is written in ```Golang```.

## Import Package
* Install the package from the command line: 

```bash
$ go install github.com/ZhenhangTung/GoGym
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
```
$ curl localhost:3000/index
{"hello":"world"}
```

## License

`GoGym` is released under the [MIT License](http://opensource.org/licenses/MIT).