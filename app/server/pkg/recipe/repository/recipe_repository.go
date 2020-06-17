package repository

import "github.com/ciazhar/golang-grpc/grpc/generated/golang"

// Recipes - Mock data of all the recipes to return for ListAllRecipes endpoint
var Recipes = []*golang.Recipe{
	{
		Name:    "Pizza",
		Cuisine: "Italian",
	},
	{
		Name:    "Beef Burgers",
		Cuisine: "American",
	},
	{
		Name:    "Croissants",
		Cuisine: "French",
	},
	{
		Name:    "Katsu Curry",
		Cuisine: "Japanese",
	},
	{
		Name:    "Butter Chicken",
		Cuisine: "Indian",
	},
	{
		Name:    "Raviolli",
		Cuisine: "Italian",
	},
	{
		Name:    "Sunday Roast",
		Cuisine: "British",
	},
	{
		Name:    "Kebabs",
		Cuisine: "Lebanese",
	},
	{
		Name:    "Bread",
		Cuisine: "International",
	},
	{
		Name:    "Duck stir fry",
		Cuisine: "Chinese",
	},
	{
		Name:    "Ceviche",
		Cuisine: "Peruvian",
	},
	{
		Name:    "Burritos",
		Cuisine: "Mexican",
	},
	{
		Name:    "Pho",
		Cuisine: "Vietnamese",
	},
	{
		Name:    "Dim Sums",
		Cuisine: "Hong Kong",
	},
	{
		Name:    "Tagines",
		Cuisine: "Moroccon",
	},
	{
		Name:    "Fish and chips",
		Cuisine: "British",
	},
}

// RecipeToIngredientsMap - mock data for mapping golangs to ingredients for GetIngredientsForAllRecipes endpoint
func RecipeToIngredientsMap() map[string][]golang.Ingredient {
	golangToIngredientsMap := make(map[string][]golang.Ingredient)
	golangToIngredientsMap["Bread"] = []golang.Ingredient{
		{Name: "golang:Bread", Quantity: ""},
		{Name: "Flour", Quantity: "10kg"},
		{Name: "Yeast", Quantity: "1 packet"},
	}
	golangToIngredientsMap["Nachos"] = []golang.Ingredient{
		{Name: "golang:Nachos", Quantity: ""},
		{Name: "Tortilla chips", Quantity: "1 packett"},
		{Name: "Sour cream", Quantity: "1 packet"},
		{Name: "Cheese", Quantity: "200 gms"},
	}
	golangToIngredientsMap["Croissants"] = []golang.Ingredient{
		{Name: "golang:Croissants", Quantity: ""},
		{Name: "500g strong white flour, plus extra for dusting", Quantity: ""},
		{Name: "1½ tsp salt", Quantity: ""},
		{Name: "50g sugar", Quantity: ""},
		{Name: "2 x 7g sachets fast-action dried yeast", Quantity: ""},
		{Name: "oil, for greasing", Quantity: ""},
		{Name: "300g butter, at room temperature", Quantity: ""},
		{Name: "1 egg, beaten", Quantity: ""},
	}
	golangToIngredientsMap["Chicken pasta bake"] = []golang.Ingredient{
		{Name: "golang:Chicken pasta bake", Quantity: ""},
		{Name: "4 tbsp olive oil", Quantity: ""},
		{Name: "1 onion, finely chopped", Quantity: ""},
		{Name: "2 garlic cloves, crushed", Quantity: ""},
		{Name: "¼ tsp chilli flakes", Quantity: ""},
		{Name: "2 x 400g cans chopped tomatoes", Quantity: ""},
		{Name: "1 tsp caster sugar", Quantity: ""},
		{Name: "6 tbsp mascarpone", Quantity: ""},
		{Name: "4 skinless chicken breasts, sliced into strips", Quantity: ""},
		{Name: "300g penne", Quantity: ""},
		{Name: "70g mature cheddar, grated", Quantity: ""},
		{Name: "50g grated mozzarella", Quantity: ""},
		{Name: "½ small bunch of parsley, finely chopped", Quantity: ""},
	}
	golangToIngredientsMap["Roast salmon with preserved lemon"] = []golang.Ingredient{
		{Name: "golang:Roast salmon with preserved lemon", Quantity: ""},
		{Name: "40g preserved lemon, flesh and pith removed", Quantity: ""},
		{Name: "100ml gin", Quantity: ""},
		{Name: "1kg side organic farmed or wild salmon (tail end)", Quantity: ""},
		{Name: "50g sea salt", Quantity: ""},
		{Name: "50g golden caster sugar", Quantity: ""},
		{Name: "1 tsp thyme leaves", Quantity: ""},
		{Name: "1 tsp chilli flakes", Quantity: ""},
		{Name: "½ small bunch dill, washed", Quantity: ""},
		{Name: "30g preserved lemons, seeds removed", Quantity: ""},
		{Name: "4 tbsp olive oil", Quantity: ""},
	}

	return golangToIngredientsMap
}
