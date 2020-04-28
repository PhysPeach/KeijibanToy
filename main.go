package main

import (
	"fmt"

	_ "KeijibanToy/routers"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("Hello")
	beego.Run()
}

