package main

import (
	"beeblog/controllers"
	"beeblog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true                      //打印所有orm信息，方便调试
	orm.RunSyncdb("default", false, true) //第一个是否每次需要重建表，，是否打印相关信息

	beego.Router("/", &controllers.MainController{})
	beego.Router("/category", &controllers.CategoryController{})
	//beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Run()
}
