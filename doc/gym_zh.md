# Gym

```Gym``` 是服务容器，也是框架的核心，负责启动项目。

## Prepare The Service
```Prepare() *Gym``` :  Prepare为Gym准备必要的服务。

```
var app = new(GoGym.Gym)
app.Prepare()
```

## Serve The App
```OpenAt(port int)``` : OpenAt会将服务在传入的端口启动。


```
var app = new(GoGym.Gym)
app.OpenAt(3000)
```