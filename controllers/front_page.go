package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Github"] = "github.com/PhysPeach/KeijibanToy"
	c.TplName = "front_page.tpl"
}
