# GoGym  


[![Build Status](https://travis-ci.org/ZhenhangTung/GoGym.svg?branch=master)](https://travis-ci.org/ZhenhangTung/GoGym)

```GoGym```是一个用Golang构建的RESTful APIs的框架。它深受[Laravel](https://laravel.com/)启发，希望能帮助用户优雅快速地构建API服务。

![](http://tongzhenhang.me/wp-content/uploads/2017/03/GoGym_Logo_256.png)   
项目Icon由 @Beth Wardolf 设计制作

## 安装项目
* 项目依赖:
	* [logrus](https://github.com/sirupsen/logrus)
	```
	$ go get github.com/sirupsen/logrus
	```

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


## 示例代码


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

## API 文档
* [Router](http://zhenhangtung.github.io/GoGym/doc/router_zh.html)
* [Request](http://zhenhangtung.github.io/GoGym/doc/request_zh.html)
* [Response](http://zhenhangtung.github.io/GoGym/doc/response_zh.html)
* [Service Container](http://zhenhangtung.github.io/GoGym/doc/gym_zh.html)


## 项目规划图
* v0.1: 接收请求，返回JSON响应 [Finished]
* v0.2: 支持路由变量 **[In development]**
* v0.3: 支持定义路由时候直接传入function
* v0.4: 支持中间件


## 想成为```GoGym```的一员

十分欢迎参与```GoGym```的开发
你可以通过如下的方式参与:

* [发布issue或者feedback](https://github.com/ZhenhangTung/GoGym/issues)
* 通过Pull Request来提交bug fix或者new feature
* 写文档或者修饰文档


## 项目贡献成员
感谢所有的[贡献成员](https://github.com/ZhenhangTung/GoGym/graphs/contributors)



## License

```GoGym```遵守[MIT License](http://opensource.org/licenses/MIT).

