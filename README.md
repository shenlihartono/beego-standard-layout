## Beego Standard Package Layout

This project is created using [Beego](https://github.com/beego/bee) API framework.

Beego is good for bootstrapping a standard Golang API Server, but in this project we're going to re-structure the project according the [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1) introduced by Ben Johnson.

### Instruction

How to run locally:
```
1. go build ./...
2. bee generate docs
3. go run cmd/beego/main.go
``` 

There's another style of Router/Controller setup using Beego, you can find it [here](https://beego.me/docs/mvc/controller/router.md)

### Friendly Warning
There may be warning such as `0001 Cannot find the object: groot.StructRequest` when generate docs. 

This is because we don't use package `models` in this project. There's nothing to worry, just ignore it.