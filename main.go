package main

import (
	_ "github.com/ygqbasic/poseidon/routers"
	_ "github.com/ygqbasic/poseidon/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
