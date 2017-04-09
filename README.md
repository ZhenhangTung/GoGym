


# GoGym


[![Build Status](https://travis-ci.org/ZhenhangTung/GoGym.svg?branch=master)](https://travis-ci.org/ZhenhangTung/GoGym)

[中文文档](./README_zh.md)


```GoGym ``` is a framework for building RESTful APIs, which is written in ```Golang```. It is inspired by an artisan framework [Laravel](https://laravel.com/).

![](http://tongzhenhang.me/wp-content/uploads/2017/03/GoGym_Logo_256.png)

Icon made by @Beth Wardolf


## Import Package
* Dependencies:
	* [logrus](https://github.com/sirupsen/logrus)
	```
	$ go get github.com/sirupsen/logrus
	```

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



## Code Example


```go
package main

import (
	"github.com/ZhenhangTung/GoGym"
	"net/http"
)

type HelloController struct {
}

func (h *HelloController) SayHello(api *GoGym.Gym) {
	api.Response.JsonResponse(map[string]string{"msg": "Hello World!"}, 200, http.Header{})
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

## API Documentation
* [Router](http://zhenhangtung.github.io/GoGym/doc/router.html)
* [Request](http://zhenhangtung.github.io/GoGym/doc/request.html)
* [Response](http://zhenhangtung.github.io/GoGym/doc/response.html)
* [Service Container](http://zhenhangtung.github.io/GoGym/doc/gym.html)


## Roadmap
* v0.1: Receive request and send JSON response. [Finished]
* v0.2: Support route parameters. **[In development]**
* v0.3: Support using functions directly when defining routes.
* v0.4: Support middleware.

## Contribution

Your contribution to ```GoGym``` development is very welcomed!
You may contribute in the following ways:

* [Post issues and feedbacks](https://github.com/ZhenhangTung/GoGym/issues).
* Submit fixes, features via the Pull Request.
* Write/polish the documentation. The documentation exists in folder ```doc```.


## Contributors
Thanks for all [contributors](https://github.com/ZhenhangTung/GoGym/graphs/contributors).



## License

`GoGym` is released under the [MIT License](http://opensource.org/licenses/MIT).
