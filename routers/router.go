// @APIVersion 1.0.0
// @Title kubernetes management api
// @Description manage api docs.
// @Contact ygqbasic@gmail.com
package routers

import (
	"github.com/ygqbasic/nuwa/controllers"

	"github.com/astaxie/beego"
)

func init() {

	//TotalCollectRate
	beego.Router("/totalcollectrate/index", &controllers.TotalCollectRateController{}, "*:Index")
	beego.Router("/totalcollectrate/datagrid", &controllers.TotalCollectRateController{}, "Get,Post:DataGrid")

	//TerminalTrace
	beego.Router("/terminaltrace/index", &controllers.TerminalTraceController{}, "*:Index")
	beego.Router("/terminaltrace/datagrid", &controllers.TerminalTraceController{}, "Get,Post:DataGrid")

	//TotalDtuRows
	beego.Router("/totaldturows/index", &controllers.TotalDtuRowsController{}, "*:Index")
	beego.Router("/totaldturows/datagrid", &controllers.TotalDtuRowsController{}, "Get,Post:DataGrid")
	beego.Router("/totaldturows/datalist", &controllers.TotalDtuRowsController{}, "Post:DataList")

	//SystemVal
	beego.Router("/systemval/index", &controllers.SystemValController{}, "*:Index")
	beego.Router("/systemval/datagrid", &controllers.SystemValController{}, "Get,Post:DataGrid")
	beego.Router("/systemval/datalist", &controllers.SystemValController{}, "Post:DataList")
	beego.Router("/systemval/edit/?:id", &controllers.SystemValController{}, "Get,Post:Edit")
	beego.Router("/systemval/delete", &controllers.SystemValController{}, "Post:Delete")

	//websocket
	beego.Router("/websocketwidget/index", &controllers.WebsocketWidgetController{}, "*:Index")
	beego.Router("/websocketwidget/ws", &controllers.WebsocketWidgetController{}, "Get:Get")

	//logintrace
	beego.Router("/logintrace/index", &controllers.LoginTraceController{}, "*:Index")
	beego.Router("/logintrace/datagrid", &controllers.LoginTraceController{}, "Get,Post:DataGrid")
	beego.Router("/logintrace/datalist", &controllers.LoginTraceController{}, "Post:DataList")

	//icons
	beego.Router("/icons/index", &controllers.IconsController{}, "*:Index")

	//totalactivepowerminute
	beego.Router("/totalactivepowerminute/index", &controllers.TotalActivePowerMinuteController{}, "*:Index")
	beego.Router("/totalactivepowerminute/datagrid", &controllers.TotalActivePowerMinuteController{}, "Get,Post:DataGrid")
	beego.Router("/totalactivepowerminute/datalist", &controllers.TotalActivePowerMinuteController{}, "Post:DataList")

	//collectbaseinfo
	beego.Router("/collectbaseinfo/index", &controllers.CollectBaseInfoController{}, "*:Index")
	beego.Router("/collectbaseinfo/datagrid", &controllers.CollectBaseInfoController{}, "Get,Post:DataGrid")
	beego.Router("/collectbaseinfo/datalist", &controllers.CollectBaseInfoController{}, "Post:DataList")

	//Cluster
	beego.Router("/cluster/index", &controllers.ClusterController{}, "*:Index")
	beego.Router("/cluster/datagrid", &controllers.ClusterController{}, "Get,Post:DataGrid")
	beego.Router("/cluster/retrieveclusters", &controllers.ClusterController{}, "Get,Post:RetrieveClusters")
	beego.Router("/cluster/datalist", &controllers.ClusterController{}, "Post:DataList")
	beego.Router("/cluster/?:clusterid/detail", &controllers.ClusterController{}, "*:Detail")
	beego.Router("/cluster/edit/?:id", &controllers.ClusterController{}, "Get,Post:Edit")
	beego.Router("/cluster/delete", &controllers.ClusterController{}, "Post:Delete")

	// beego.Router("/cluster/?:clusterid/hosts", &controllers.HostController{}, "Get,Post:RetrieveHosts")
	// beego.Router("/cluster/?:clusterid/hosts/?:host_id", &controllers.HostController{}, "Get,Post:retrieveHosts")
	// beego.Router("/cluster/?:clusterid/components", &controllers.ClusterController{}, "Get,Post:Component")
	// beego.Router("/cluster/?:clusterid/components/?:component_id", &controllers.ClusterController{}, "Get,Post:Component")

	//ClusterHost
	beego.Router("/clusterhost/index", &controllers.ClusterHostController{}, "*:Index")
	beego.Router("/clusterhost/retrieveHosts/?:cluster_id", &controllers.ClusterHostController{}, "Get,Post:RetrieveHosts")
	beego.Router("/clusterhost/datagrid", &controllers.ClusterHostController{}, "Get,Post:DataGrid")
	beego.Router("/clusterhost/datalist", &controllers.ClusterHostController{}, "Post:DataList")
	beego.Router("/clusterhost/edit/?:id", &controllers.ClusterHostController{}, "Get,Post:Edit")
	beego.Router("/clusterhost/delete", &controllers.ClusterHostController{}, "Post:Delete")

	//ClusterComponent
	beego.Router("/clustercomponent/index", &controllers.ClusterComponentController{}, "*:Index")
	beego.Router("/clustercomponent/retrievecomponents/?:cluster_id", &controllers.ClusterComponentController{}, "Get,Post:RetrieveComponents")
	beego.Router("/clustercomponent/datagrid", &controllers.ClusterComponentController{}, "Get,Post:DataGrid")
	beego.Router("/clustercomponent/datalist", &controllers.ClusterComponentController{}, "Post:DataList")
	beego.Router("/clustercomponent/edit/?:id", &controllers.ClusterComponentController{}, "Get,Post:Edit")
	beego.Router("/clustercomponent/delete", &controllers.ClusterComponentController{}, "Post:Delete")

	//Host
	beego.Router("/host/index", &controllers.HostController{}, "*:Index")
	beego.Router("/host/retrieveHosts/?:cluster_id", &controllers.HostController{}, "Get,Post:RetrieveHosts")
	beego.Router("/host/datagrid", &controllers.HostController{}, "Get,Post:DataGrid")
	beego.Router("/host/datalist", &controllers.HostController{}, "Post:DataList")
	beego.Router("/host/edit/?:id", &controllers.HostController{}, "Get,Post:Edit")
	beego.Router("/host/delete", &controllers.HostController{}, "Post:Delete")

	//BackendConf
	beego.Router("/backendconf/index", &controllers.BackendConfController{}, "*:Index")
	beego.Router("/backendconf/datagrid", &controllers.BackendConfController{}, "Get,Post:DataGrid")
	beego.Router("/backendconf/datalist", &controllers.BackendConfController{}, "Post:DataList")
	beego.Router("/backendconf/edit/?:id", &controllers.BackendConfController{}, "Get,Post:Edit")
	beego.Router("/backendconf/delete", &controllers.BackendConfController{}, "Post:Delete")

	//用户角色路由
	beego.Router("/role/index", &controllers.RoleController{}, "*:Index")
	beego.Router("/role/datagrid", &controllers.RoleController{}, "Get,Post:DataGrid")
	beego.Router("/role/edit/?:id", &controllers.RoleController{}, "Get,Post:Edit")
	beego.Router("/role/delete", &controllers.RoleController{}, "Post:Delete")
	beego.Router("/role/datalist", &controllers.RoleController{}, "Post:DataList")
	beego.Router("/role/allocate", &controllers.RoleController{}, "Post:Allocate")
	beego.Router("/role/updateseq", &controllers.RoleController{}, "Post:UpdateSeq")

	//资源路由
	beego.Router("/resource/index", &controllers.ResourceController{}, "*:Index")
	beego.Router("/resource/treegrid", &controllers.ResourceController{}, "POST:TreeGrid")
	beego.Router("/resource/edit/?:id", &controllers.ResourceController{}, "Get,Post:Edit")
	beego.Router("/resource/parent", &controllers.ResourceController{}, "Post:ParentTreeGrid")
	beego.Router("/resource/delete", &controllers.ResourceController{}, "Post:Delete")

	//快速修改顺序
	beego.Router("/resource/updateseq", &controllers.ResourceController{}, "Post:UpdateSeq")

	//通用选择面板
	beego.Router("/resource/select", &controllers.ResourceController{}, "Get:Select")
	beego.Router("/resource/chooseIcon", &controllers.ResourceController{}, "Get:ChooseIcon")

	//用户有权管理的菜单列表（包括区域）
	beego.Router("/resource/usermenutree", &controllers.ResourceController{}, "POST:UserMenuTree")
	beego.Router("/resource/checkurlfor", &controllers.ResourceController{}, "POST:CheckUrlFor")

	//后台用户路由
	beego.Router("/backenduser/index", &controllers.BackendUserController{}, "*:Index")
	beego.Router("/backenduser/datagrid", &controllers.BackendUserController{}, "POST:DataGrid")
	beego.Router("/backenduser/edit/?:id", &controllers.BackendUserController{}, "Get,Post:Edit")
	beego.Router("/backenduser/delete", &controllers.BackendUserController{}, "Post:Delete")

	//后台用户中心
	beego.Router("/usercenter/profile", &controllers.UserCenterController{}, "Get:Profile")
	beego.Router("/usercenter/basicinfosave", &controllers.UserCenterController{}, "Post:BasicInfoSave")
	beego.Router("/usercenter/uploadimage", &controllers.UserCenterController{}, "Post:UploadImage")
	beego.Router("/usercenter/passwordsave", &controllers.UserCenterController{}, "Post:PasswordSave")

	beego.Router("/home/index", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/index2", &controllers.HomeController{}, "*:Index2")
	beego.Router("/home/login", &controllers.HomeController{}, "*:Login")
	beego.Router("/home/dologin", &controllers.HomeController{}, "Post:DoLogin")
	beego.Router("/home/logout", &controllers.HomeController{}, "*:Logout")
	beego.Router("/home/404", &controllers.HomeController{}, "*:Page404")
	beego.Router("/home/error/?:error", &controllers.HomeController{}, "*:Error")

	//beego.Router("/home/metercount", &controllers.HomeController{}, "*:GetMeterCount")
	//beego.Router("/home/collectrowstoday", &controllers.HomeController{}, "*:GetCollectRowsToday")

	beego.Router("/home/configvalue", &controllers.HomeController{}, "*:GetConfigValue")
	beego.Router("/home/weather", &controllers.HomeController{}, "*:GetWeather")

	beego.Router("/", &controllers.HomeController{}, "*:Index")

	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/cluster",
				beego.NSInclude(
					&controllers.ClusterController{},
				),
			),
			beego.NSNamespace("/host",
				beego.NSInclude(
					&controllers.HostController{},
				),
			),
		)
	beego.AddNamespace(ns)

}
