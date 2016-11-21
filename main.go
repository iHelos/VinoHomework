package main

import (
	"github.com/kataras/iris"
	. "github.com/iHelos/VinoHomework/controller"
)

var control BusinessTransaction

func init(){
	control = BusinessTransaction{}
	control.Start()
}

func main() {
	control.CreateDish()
	iris.Listen("127.0.0.1:8080")
}
