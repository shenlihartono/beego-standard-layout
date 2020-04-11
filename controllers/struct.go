// Package controllers is for mapping controller functions to endpoints in routers.
package controllers

import (
	groot "beego-standard-layout"
	"beego-standard-layout/inmemory"
	"beego-standard-layout/json"
	"beego-standard-layout/postgres"
	"beego-standard-layout/service"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var repo groot.StructRepository

func InitPostgresRepo(orm orm.Ormer) {
	repo = postgres.NewStructRepository(orm)
}

func InitInmemoryRepo() {
	repo = inmemory.NewStructRepository()
}

// Operations about struct
type StructController struct {
	beego.Controller
}

// @Description Create Struct
// @Param	body		body 	groot.StructRequest		true	"The struct content, e.q {"value": 1}"
// @Success 200 {result} the Struct
// @Failure 400 bad request
// @Failure 500 internal server error
// @router / [post]
func (s *StructController) Create() {
	var req groot.StructRequest
	err := json.ConvertRequest(s.Ctx.Input.RequestBody, &req)
	if err != nil {
		s.CustomAbort(400, "bad request")
	}

	svc := service.NewStructService(repo)
	resp := svc.CreateStruct(req)

	if resp.Status != 200 {
		body := resp.Result.(string)
		s.CustomAbort(resp.Status, body)
	}

	s.Data["json"] = resp.Result
	s.ServeJSON()
}

// @Description Find one struct by id
// @Param	structId	path 	string		true	"the struct id you want to find"
// @Success 200 {result} the Struct
// @Failure 404 struct not found
// @Failure 500 internal server error
// @router /:structId [get]
func (s *StructController) Struct() {
	structId := s.Ctx.Input.Param(":structId")

	svc := service.NewStructService(repo)
	resp := svc.Struct(structId)

	if resp.Status != 200 {
		body := resp.Result.(string)
		s.CustomAbort(resp.Status, body)
	}

	s.Data["json"] = resp.Result
	s.ServeJSON()
}

// @Description Find all structs
// @Success 200 {result} the Struct
// @Failure 404 struct not found
// @Failure 500 internal server error
// @router / [get]
func (s *StructController) Structs() {
	svc := service.NewStructService(repo)
	resp := svc.Structs()

	if resp.Status != 200 {
		body := resp.Result.(string)
		s.CustomAbort(resp.Status, body)
	}

	s.Data["json"] = resp.Result
	s.ServeJSON()
}

// @Description Update a Struct
// @Param	structId	path 	string					true	"the struct id you want to update"
// @Param	body		body 	groot.StructRequest		true	"The struct content, e.q {"value": 1}"
// @Success 200 {result} the Struct
// @Failure 400 bad request
// @Failure 404 struct not found
// @Failure 500 internal server error
// @router /:structId [put]
func (s *StructController) Update() {
	structId := s.Ctx.Input.Param(":structId")
	if structId == "" {
		s.CustomAbort(400, "bad request")
	}

	var req groot.StructRequest
	err := json.ConvertRequest(s.Ctx.Input.RequestBody, &req)
	if err != nil {
		s.CustomAbort(400, "bad request")
	}

	svc := service.NewStructService(repo)
	resp := svc.UpdateStruct(structId, req)

	if resp.Status != 200 {
		body := resp.Result.(string)
		s.CustomAbort(resp.Status, body)
	}

	s.Data["json"] = resp.Result
	s.ServeJSON()
}

// @Description Delete one struct by StructID
// @Param	structId	path 	string		true	"the struct id you want to delete"
// @Success 200 {string} success delete struct
// @Failure 404 struct not found
// @Failure 500 internal server error
// @router /:structId [delete]
func (s *StructController) DeleteStruct() {
	structId := s.Ctx.Input.Param(":structId")

	svc := service.NewStructService(repo)
	resp := svc.Delete(structId)

	if resp.Status != 200 {
		body := resp.Result.(string)
		s.CustomAbort(resp.Status, body)
	}

	s.Data["json"] = resp.Result
	s.ServeJSON()
}
