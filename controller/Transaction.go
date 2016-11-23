package controller

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/iHelos/VinoHomework/model/ingredient"
	"github.com/iHelos/VinoHomework/model/dish"
	"github.com/iHelos/VinoHomework/model/kitchen"
	"log"
	"github.com/iHelos/VinoHomework/model/link_DI"
	"github.com/iHelos/VinoHomework/model/link_DK"
	"github.com/kataras/iris"
	"encoding/json"
	"strconv"
)

type BusinessTransaction struct {
	connection_pool *sql.DB
}

func (*BusinessTransaction) Start() (bool){
	db, err := sql.Open("postgres", "postgresql://admin:keks@139.59.133.90:5432/KitchenDB")
	if err != nil {
		log.Fatal(err)
	}

	ingredient.Start(db)
	kitchen.Start(db)
	dish.Start(db)
	link_DI.Start(db)
	link_DK.Start(db)
	return true
}

func (*BusinessTransaction) CreateDish(ctx *iris.Context) (bool){
	dish_obj := dish.NewDish()
	err := json.Unmarshal(ctx.Request.Body(), &dish_obj)
	if err!=nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	err = dish_obj.Insert()
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"dish":dish_obj}))
	return true
}

func (*BusinessTransaction) CreateIngredient(ctx *iris.Context) (bool){
	dish_obj := ingredient.NewIngredient_Local("нет имени", "нет описания")
	err := json.Unmarshal(ctx.Request.Body(), &dish_obj)
	if err!=nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	err = dish_obj.Insert()
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"ingredient":dish_obj}))
	return true
}

func (*BusinessTransaction) CreateKitchen(ctx *iris.Context) (bool){
	dish_obj := kitchen.NewKitchen_Local("нет имени","нет описания")
	err := json.Unmarshal(ctx.Request.Body(), &dish_obj)
	if err!=nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	err = dish_obj.Insert()
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"kitchen":dish_obj}))
	return true
}

func (*BusinessTransaction) ReadDish(ctx *iris.Context) (bool){
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	dish_obj, err := dish.Find(id_int)
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err.Error()}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"dish":dish_obj}))
	return true
}

func (*BusinessTransaction) ReadIngredient(ctx *iris.Context) (bool){
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	dish_obj, err := ingredient.Find(id_int)

	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err.Error()}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"ingredient":dish_obj}))
	return true
}

func (*BusinessTransaction) ReadKitchen(ctx *iris.Context) (bool){
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	dish_obj, err := kitchen.Find(id_int)
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err.Error()}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"kitchen":dish_obj}))
	return true
}

func (*BusinessTransaction) UpdateDish(ctx *iris.Context) (bool){
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	dish_obj, err := dish.Find(id_int)
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err.Error()}))
		return false
	}
	err = json.Unmarshal(ctx.Request.Body(), &dish_obj)
	if err!=nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	err = dish_obj.Update()
	if err!=nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"dish":dish_obj}))
	return true
}

func (*BusinessTransaction) UpdateIngredient(ctx *iris.Context) (bool){
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	dish_obj, err := ingredient.Find(id_int)

	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err.Error()}))
		return false
	}
	err = json.Unmarshal(ctx.Request.Body(), &dish_obj)
	if err!=nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	err = dish_obj.Update()
	if err!=nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"ingredient":dish_obj}))
	return true
}

func (*BusinessTransaction) UpdateKitchen(ctx *iris.Context) (bool){
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	dish_obj, err := kitchen.Find(id_int)
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err.Error()}))
		return false
	}
	err = json.Unmarshal(ctx.Request.Body(), &dish_obj)
	if err!=nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	err = dish_obj.Update()
	if err!=nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"kitchen":dish_obj}))
	return true
}


func (*BusinessTransaction) RemoveDish(ctx *iris.Context) (bool){
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	dish_obj, err := dish.Find(id_int)
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err.Error()}))
		return false
	}
	err = json.Unmarshal(ctx.Request.Body(), &dish_obj)
	if err!=nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	err = dish_obj.Remove()
	if err!=nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"status":"ok"}))
	return true
}

func (*BusinessTransaction) RemoveIngredient(ctx *iris.Context) (bool){
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	dish_obj, err := ingredient.Find(id_int)

	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err.Error()}))
		return false
	}
	err = json.Unmarshal(ctx.Request.Body(), &dish_obj)
	if err!=nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	err = dish_obj.Remove()
	if err!=nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"status":"ok"}))
	return true
}

func (*BusinessTransaction) RemoveKitchen(ctx *iris.Context) (bool){
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	dish_obj, err := kitchen.Find(id_int)
	if err!=nil{
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err.Error()}))
		return false
	}
	err = json.Unmarshal(ctx.Request.Body(), &dish_obj)
	if err!=nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	err = dish_obj.Remove()
	if err!=nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error":err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"status":"ok"}))
	return true
}