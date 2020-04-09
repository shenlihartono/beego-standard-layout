## Introduction

This project is created using [Beego](https://beego.me/blog/beego_api) API framework.

Beego is good for bootstrapping a standard Golang API Server, but in this project we're going to re-structure the project according the [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1) introduced by Ben Johnson.

### Instruction

How to run locally:
```
1. go build ./...
2. bee generate docs
3. go run cmd/beego/main.go
``` 

There's another style of Router/Controller setup using Beego, you can find it [here](https://beego.me/docs/mvc/controller/router.md)
