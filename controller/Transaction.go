package controller

import (
	"database/sql"
	"fmt"
	"github.com/iHelos/VinoHomework/model/dish"
	"github.com/iHelos/VinoHomework/model/ingredient"
	"github.com/iHelos/VinoHomework/model/kitchen"
	"github.com/iHelos/VinoHomework/model/link_DI"
	"github.com/iHelos/VinoHomework/model/link_DK"
	logsmodel "github.com/iHelos/VinoHomework/model/logs"
	systemlogs "github.com/iHelos/VinoHomework/system"
	decorators "github.com/iHelos/VinoHomework/system/concreteDecorator"
	_ "github.com/lib/pq"
	"gopkg.in/kataras/iris.v6"
	"log"
	"strconv"
)

type BusinessTransaction struct {
	connection_pool *sql.DB
	logger systemlogs.LogComponent
}

func (b *BusinessTransaction) Start() bool {

	b.logger = systemlogs.Decorate(
		systemlogs.LogStr(
			func(s string) {
				fmt.Println(s)
			},
		),
		decorators.DBLogDecorator,
		decorators.TimeLogDecorator,		//Зависит от расположения
	)

	db, err := sql.Open("postgres", "postgresql://localhost:5432/kitchendb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	ingredient.Start(db)
	kitchen.Start(db)
	dish.Start(db)
	link_DI.Start(db)
	link_DK.Start(db)
	logsmodel.Start(db)
	return true
}

func (b *BusinessTransaction) CreateDish(ctx *iris.Context) bool {
	b.logger.Log("CreateDish")
	dish_obj := dish.NewDish()
	err := ctx.ReadJSON(&dish_obj)
	log.Print(dish_obj)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	err = dish_obj.Insert()
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"dish": dish_obj}))
	return true
}

func (b *BusinessTransaction) CreateIngredient(ctx *iris.Context) bool {
	b.logger.Log("CreateIngredient")
	dish_obj := ingredient.NewIngredient_Local("нет имени", "нет описания")
	err := ctx.ReadJSON(&dish_obj)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	err = dish_obj.Insert()
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"ingredient": dish_obj}))
	return true
}

func (b *BusinessTransaction) CreateKitchen(ctx *iris.Context) bool {
	b.logger.Log("CreateKitchen")
	dish_obj := kitchen.NewKitchen_Local("нет имени", "нет описания")
	err := ctx.ReadJSON(&dish_obj)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	err = dish_obj.Insert()
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"kitchen": dish_obj}))
	return true
}

func (b *BusinessTransaction) ReadDish(ctx *iris.Context) bool {
	b.logger.Log("ReadDish")
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	dish_obj, err := dish.Find(id_int)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err.Error()}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"dish": dish_obj}))
	return true
}

func (b *BusinessTransaction) ReadIngredient(ctx *iris.Context) bool {
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	dish_obj, err := ingredient.Find(id_int)

	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err.Error()}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"ingredient": dish_obj}))
	return true
}

func (b *BusinessTransaction) ReadKitchen(ctx *iris.Context) bool {
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	dish_obj, err := kitchen.Find(id_int)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err.Error()}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"kitchen": dish_obj}))
	return true
}

func (b *BusinessTransaction) UpdateDish(ctx *iris.Context) bool {
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	dish_obj, err := dish.Find(id_int)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err.Error()}))
		return false
	}
	err = ctx.ReadJSON(&dish_obj)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	err = dish_obj.Update()
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"dish": dish_obj}))
	return true
}

func (b *BusinessTransaction) UpdateIngredient(ctx *iris.Context) bool {
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	dish_obj, err := ingredient.Find(id_int)

	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err.Error()}))
		return false
	}
	err = ctx.ReadJSON(&dish_obj)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	err = dish_obj.Update()
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"ingredient": dish_obj}))
	return true
}

func (b *BusinessTransaction) UpdateKitchen(ctx *iris.Context) bool {
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	dish_obj, err := kitchen.Find(id_int)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err.Error()}))
		return false
	}
	err = ctx.ReadJSON(&dish_obj)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	err = dish_obj.Update()
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"kitchen": dish_obj}))
	return true
}

func (b *BusinessTransaction) RemoveDish(ctx *iris.Context) bool {
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	dish_obj, err := dish.Find(id_int)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err.Error()}))
		return false
	}
	err = ctx.ReadJSON(&dish_obj)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	err = dish_obj.Remove()
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"status": "ok"}))
	return true
}

func (b *BusinessTransaction) RemoveIngredient(ctx *iris.Context) bool {
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	dish_obj, err := ingredient.Find(id_int)

	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err.Error()}))
		return false
	}
	err = ctx.ReadJSON(&dish_obj)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	err = dish_obj.Remove()
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"status": "ok"}))
	return true
}

func (b *BusinessTransaction) RemoveKitchen(ctx *iris.Context) bool {
	id := ctx.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	dish_obj, err := kitchen.Find(id_int)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err.Error()}))
		return false
	}
	err = ctx.ReadJSON(&dish_obj)
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	err = dish_obj.Remove()
	if err != nil {
		log.Print(err)
		ctx.JSON(iris.StatusOK, GetResponse(0, map[string]interface{}{"error": err}))
		return false
	}
	ctx.JSON(iris.StatusOK, GetResponse(1, map[string]interface{}{"status": "ok"}))
	return true
}
