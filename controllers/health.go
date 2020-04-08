// Package controllers for checking this service's Health.
package controllers

import (
	"github.com/astaxie/beego"
)

// HealthController is controller for Health operation.
type HealthController struct {
	beego.Controller
}

// @Description Health Controller
// @Success 200 {string} service up and running!
// @router / [get]
func (h *HealthController) Health() {
	h.Data["json"] = "service up and running!"
	h.ServeJSON()
}
