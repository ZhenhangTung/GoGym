# Request

## 获取请求的方法
```Request.Method```

```
func (u *UsersController) Get(Gym *GoGym.Gym) {
	method := Gym.Request.Method
}
```
## 获取请求的header
```Request.Header```

```
func (u *UsersController) Get(Gym *GoGym.Gym) {
	fooHeader := Gym.Request.Header.Get("Foo")
}
```

## 获取请求时候的query参数
```Request.Query```

```
func (u *UsersController) Get(Gym *GoGym.Gym) {
	fooValue := Gym.Request.Query.Get("Foo")
}
```

## 获取表单

```Request.Form```

```
func (u *UsersController) Get(Gym *GoGym.Gym) {
	fooValue := Gym.Request.Form.Get("foo")
}
```
