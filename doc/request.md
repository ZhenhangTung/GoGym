# Request

## Get Request Method
```Request.Method```: It gets the method of the http request.

```
func (u *UsersController) Get(Gym *GoGym.Gym) {
	method := Gym.Request.Method
}
```
## Get Request Header
```Request.Header```: It gets the header of the http request.

```
func (u *UsersController) Get(Gym *GoGym.Gym) {
	fooHeader := Gym.Request.Header.Get("Foo")
}
```

## Get Request Query
```Request.Query``` : It parses query of the http request.

```
func (u *UsersController) Get(Gym *GoGym.Gym) {
	fooValue := Gym.Request.Query.Get("Foo")
}
```

## Get Request Form

```Request.Form``` : It parses request form of the http request.

```
func (u *UsersController) Get(Gym *GoGym.Gym) {
	fooValue := Gym.Request.Form.Get("foo")
}
```
