#Router

# Basic Routing

### Get

```Get(path, action string)``` : Get is a fucntion handles GET requests

```
Gym.Router.Get("/users", "UsersController@Get")
```

### Post

```Post(path, action string)``` : Post is a fucntion handles POST requests

```
Gym.Router.Post("/users", "UsersController@Post")
```

### Put

```Put(path, action string)``` : Put is a method handles PUT requests

```
Gym.Router.Post("/users", "UsersController@Post")
```

### Patch

```Patch(path, action string)``` : Patch is a method handles PATCH requests

```
Gym.Router.Patch("/users", "UsersController@Patch")
```

### Options

```Options(path, action string)``` : Options is a method handles OPTIONS requests

```
Gym.Router.Options("/users", "UsersController@Options")
```

### Delete

```Delete(path, action string)``` : Delete is a method handles DELETE requests

```
Gym.Router.Delete("/users", "UsersController@Delete")
```

## Route Parameters
### Required Parameters
```
var gym = new(GoGym.Gym)
gym.Prepare()
gym.Router.Get("/users/{id}", "UsersController@Get")
```


## Regitser Controller

### RegisterController

```RegisterController(controller interface{})``` : RegisterControllers is a method registers a controller.

Usage Example:

```
type HelloController struct {}
var gym = new(GoGym.Gym)
gym.Prepare()
gym.Router.RegisterController(&HelloController{})

```

### RegisterControllers

 ```RegisterControllers(controllers []interface{})``` : RegisterControllers is a method registers a bunch of controllers.
 
```
var gym = new(GoGym.Gym)
gym.Prepare()
controllers := []interface{}{}
controllers = append(controllers, &HelloController{})
gym.Router.RegisterControllers(controllers)
```
 