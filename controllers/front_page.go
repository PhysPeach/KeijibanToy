package controllers

import (
	"github.com/astaxie/beego"
	"KeijibanToy/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Github"] = "github.com/PhysPeach/KeijibanToy"

	posts := models.GetAllPosts()
	c.Data["posts"] = posts

	c.TplName = "front_page.tpl"
}
