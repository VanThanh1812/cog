package main

import (
	_ "cog/routers"

	"github.com/astaxie/beego"
	"cog/pay"
)

func main() {

	pay.InitializeContract()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
