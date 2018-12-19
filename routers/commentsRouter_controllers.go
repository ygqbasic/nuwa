package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/ygqbasic/nuwa/controllers:ClusterController"] = append(beego.GlobalControllerRouter["github.com/ygqbasic/nuwa/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "RetrieveClusters",
            Router: `/retrieveclusters`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ygqbasic/nuwa/controllers:ClusterController"] = append(beego.GlobalControllerRouter["github.com/ygqbasic/nuwa/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "Save",
            Router: `/save`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
