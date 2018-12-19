package main

import (
	_ "github.com/ygqbasic/nuwa/routers"
	_ "github.com/ygqbasic/nuwa/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//beego.SetStaticPath("/swagger", "swagger")
	beego.Run()
}
