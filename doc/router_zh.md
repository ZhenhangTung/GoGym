# 路由

# 基本路由

### Get

```Get(path, action string)``` : 负责处理Get请求。

```
Gym.Router.Get("/users", "UsersController@Get")
```

### Post

```Post(path, action string)``` : 负责处理Post请求。

```
Gym.Router.Post("/users", "UsersController@Post")
```

### Put

```Put(path, action string)``` : 负责处理Put请求。

```
Gym.Router.Post("/users", "UsersController@Post")
```

### Patch

```Patch(path, action string)``` : 负责处理Patch请求。

```
Gym.Router.Patch("/users", "UsersController@Patch")
```

### Options

```Options(path, action string)``` : 负责处理Options请求。

```
Gym.Router.Options("/users", "UsersController@Options")
```

### Delete

```Delete(path, action string)``` : 负责处理Delete请求。

```
Gym.Router.Delete("/users", "UsersController@Delete")
```

## 路由变量
```
var gym = new(GoGym.Gym)
gym.Prepare()
gym.Router.Get("/users/{id}", "UsersController@Get")
```


## 注册控制器

### RegisterController

```RegisterController(controller interface{})``` : 注册单个控制器。



```
type HelloController struct {}
var gym = new(GoGym.Gym)
gym.Prepare()
gym.Router.RegisterController(&HelloController{})

```

### RegisterControllers

 ```RegisterControllers(controllers []interface{})``` : 注册多个控制器。
 
```
var gym = new(GoGym.Gym)
gym.Prepare()
controllers := []interface{}{}
controllers = append(controllers, &HelloController{})
gym.Router.RegisterControllers(controllers)
```
 