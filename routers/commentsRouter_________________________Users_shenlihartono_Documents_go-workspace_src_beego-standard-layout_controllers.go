package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego-standard-layout/controllers:HealthController"] = append(beego.GlobalControllerRouter["beego-standard-layout/controllers:HealthController"],
        beego.ControllerComments{
            Method: "Health",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-standard-layout/controllers:StructController"] = append(beego.GlobalControllerRouter["beego-standard-layout/controllers:StructController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
