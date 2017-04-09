# Gym

```Gym``` is a service container

## Prepare The Service
```Prepare() *Gym``` :  Prepare is a function prepares the service container.
```
var app = new(GoGym.Gym)
app.Prepare()
```

## Serve The App
```OpenAt(port int)``` : OpenAt is a function which is used to serve the app.

```
var app = new(GoGym.Gym)
app.OpenAt(3000)
```