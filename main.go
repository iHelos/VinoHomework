package main

import (
	"github.com/kataras/iris"
	. "github.com/iHelos/VinoHomework/controller"
	"github.com/iris-contrib/middleware/recovery"
	"github.com/iris-contrib/middleware/cors"
)

var control BusinessTransaction

func init(){
	control = BusinessTransaction{}
	control.Start()

	iris.Use(recovery.New())
	iris.Config.IsDevelopment = false
	iris.Config.Gzip = false
	iris.Config.Charset = "UTF-8"
	iris.Config.Sessions.DisableSubdomainPersistence = false
	iris.StaticServe("./static/web_files", "/static")

	cors_config := cors.Options{
		AllowedOrigins:[]string{"*"},
		AllowedMethods:[]string{"GET", "POST", "OPTIONS", ""},
		AllowCredentials:true,
		MaxAge:5,
		Debug:false,
	}

	cors_obj := cors.New(cors_config)

	iris.Use(cors_obj)
}

func main() {
	//CREATE

	iris.Post("/dish/create", func(ctx *iris.Context) {
		control.CreateDish(ctx)
	})
	iris.Post("/ingredient/create", func(ctx *iris.Context) {
		control.CreateIngredient(ctx)
	})
	iris.Post("/kitchen/create", func(ctx *iris.Context) {
		control.CreateKitchen(ctx)
	})

	//DELETE

	iris.Post("/dish/delete/:id", func(ctx *iris.Context) {
		control.CreateDish(ctx)
	})
	iris.Post("/ingredient/delete/:id", func(ctx *iris.Context) {
		control.CreateIngredient(ctx)
	})
	iris.Post("/kitchen/delete/:id", func(ctx *iris.Context) {
		control.CreateKitchen(ctx)
	})

	//READ

	iris.Get("/dish/read/:id", func(ctx *iris.Context) {
		control.ReadDish(ctx)
	})
	iris.Get("/ingredient/read/:id", func(ctx *iris.Context) {
		control.ReadIngredient(ctx)
	})
	iris.Get("/kitchen/read/:id", func(ctx *iris.Context) {
		control.ReadKitchen(ctx)
	})

	//UPDATE

	iris.Get("/dish/update/:id", func(ctx *iris.Context) {
		control.CreateDish(ctx)
	})
	iris.Get("/ingredient/update/:id", func(ctx *iris.Context) {
		control.CreateIngredient(ctx)
	})
	iris.Get("/kitchen/update/:id", func(ctx *iris.Context) {
		control.CreateKitchen(ctx)
	})

	iris.Listen("127.0.0.1:8080")
}
