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

}
