package controllers

import (
	"time"
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

func (c *MainController) Post() {
	post := models.Post{
		Name: c.GetString("inputName"),
		Text: c.GetString("inputText"),
		Created: time.Now()}
	models.AddPost(&post)
	c.Ctx.Redirect(302, "/")
}