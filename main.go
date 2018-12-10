package main

import (
	_ "github.com/ygqbasic/nuwa/routers"
	_ "github.com/ygqbasic/nuwa/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
