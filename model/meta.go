package model

const(
	CategoryTable = "category"
	C_ID = "id"
	C_Name = "name"

	DishTable = "dish"

	D_ID = "id"
	D_Name = "name"
	D_Description = "description"
	D_Price = "price"
	D_Category = "category_id"

	IngredientTable = "ingredient"
	I_ID = "id"
	I_name = "name"
	I_description = "description"

	KitchenTable = "kitchen"
	K_ID = "id"
	K_name = "name"
	K_description = "description"

	DITable = "dish_ingredient"
	DI_ID = "id"
	DI_dish_ID = "dish_id"
	DI_ingredient_ID = "ingredient_id"

	DKTable = "dish_kitchen"
	DK_ID = "id"
	DK_dish_ID = "dish_id"
	DK_kitchen_ID = "kitchen_id"
)
