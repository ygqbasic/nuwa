package main

import (
	"flag"

	_ "github.com/ygqbasic/nuwa/routers"
	"github.com/ygqbasic/nuwa/runtime"
	_ "github.com/ygqbasic/nuwa/sysinit"

	"github.com/astaxie/beego"
)

func main() {

	workDir := flag.String("w", "./playbook", "ansible playbook should be placed in it")

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	flag.Parse()

	runtime.Run(*workDir)
	//beego.SetStaticPath("/swagger", "swagger")
	beego.Run()
}
