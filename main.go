package main

import (
	"fmt"

	_ "KeijibanToy/routers"
	"github.com/astaxie/beego"

	_ "KeijibanToy/models"
)

func main() {
	fmt.Println("Hello")
	beego.Run()
}

