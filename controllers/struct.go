// Package controllers is for mapping controller functions to endpoints in routers.
package controllers

import (
	groot "beego-standard-layout"
	"beego-standard-layout/inmemory"
	"beego-standard-layout/json"
	"beego-standard-layout/response"
	"beego-standard-layout/service"
	"github.com/astaxie/beego"
)

// Operations about struct
type StructController struct {
	beego.Controller
}

// @Description Create Struct
// @Param	body		body 	groot.StructRequest		true	"The struct content"
// @Success 200 {object} groot.Struct
// @router / [post]
func (s *StructController) Create() {
	var req groot.StructRequest
	err := json.ConvertRequest(s.Ctx.Input.RequestBody, &req)
	if err != nil {
		s.Data["json"] = response.Error("invalid request")
		s.ServeJSON()
	}

	//TODO must be able to switch inmemory / postgres impl (possibly using conf)
	repo := inmemory.NewStructRepository()
	svc := service.NewStructService(repo)

	s.Data["json"] = svc.CreateStruct(req)
	s.ServeJSON()
}
