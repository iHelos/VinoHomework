package main

import (
	. "github.com/iHelos/VinoHomework/controller"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

var control BusinessTransaction

func main() {
	control = BusinessTransaction{}
	control.Start()
	app := iris.New()
	//app.Use(recovery.New())
	//app.Config.IsDevelopment = false
	//app.Config.Gzip = false
	//app.Config.Charset = "UTF-8"
	//app.Config.Sessions.DisableSubdomainPersistence = false
	//app.StaticServe("./static/web_files", "/static")
	//
	//cors_obj := cors.New()
	//
	//app.Use(cors_obj)
	app.Adapt(
		// adapt a logger which prints all errors to the os.Stdout
		iris.DevLogger(),
		// adapt the adaptors/httprouter or adaptors/gorillamux
		httprouter.New(),
		// 5 template engines are supported out-of-the-box:
		//
		// - standard html/template
		// - amber
		// - django
		// - handlebars
		// - pug(jade)
		//
		// Use the html standard engine for all files inside "./views" folder with extension ".html"
		//view.HTML("./views", ".html"),
		// Cors wrapper to the entire application, allow all origins.
		cors.New(cors.Options{AllowedOrigins: []string{"*"}}))
	//CREATE
	app.Post("/dish/create", func(ctx *iris.Context) {
		control.CreateDish(ctx)
	})
	app.Post("/ingredient/create", func(ctx *iris.Context) {
		control.CreateIngredient(ctx)
	})
	app.Post("/kitchen/create", func(ctx *iris.Context) {
		control.CreateKitchen(ctx)
	})

	//DELETE

	app.Post("/dish/delete/:id", func(ctx *iris.Context) {
		control.RemoveDish(ctx)
	})
	app.Post("/ingredient/delete/:id", func(ctx *iris.Context) {
		control.RemoveIngredient(ctx)
	})
	app.Post("/kitchen/delete/:id", func(ctx *iris.Context) {
		control.RemoveKitchen(ctx)
	})

	//READ

	app.Get("/dish/read/:id", func(ctx *iris.Context) {
		control.ReadDish(ctx)
	})
	app.Get("/ingredient/read/:id", func(ctx *iris.Context) {
		control.ReadIngredient(ctx)
	})
	app.Get("/kitchen/read/:id", func(ctx *iris.Context) {
		control.ReadKitchen(ctx)
	})

	//UPDATE

	app.Post("/dish/update/:id", func(ctx *iris.Context) {
		control.UpdateDish(ctx)
	})
	app.Post("/ingredient/update/:id", func(ctx *iris.Context) {
		control.UpdateIngredient(ctx)
	})
	app.Post("/kitchen/update/:id", func(ctx *iris.Context) {
		control.UpdateKitchen(ctx)
	})

	app.Listen("127.0.0.1:8080")
}
