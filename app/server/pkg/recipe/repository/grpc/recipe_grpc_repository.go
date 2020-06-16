package grpc

import (
	"context"
	"github.com/ciazhar/golang-grpc/grpc/generated/go/recipe"
	"google.golang.org/grpc"
	"io"
	"log"
)

type repository struct {
}

func (r repository) AddRecipe(ctx context.Context, request *recipe.AddRecipeRequest) (*recipe.AddRecipeResponse, error) {
	log.Printf("Adding new recipe\t\tName: %v, Cuisine: %v", request.GetRecipe().GetName(), request.GetRecipe().GetCuisine())

	response := &recipe.AddRecipeResponse{Success: true}

	return response, nil
}

func (r repository) ListAllRecipesArray(request *recipe.ListAllRecipesRequest, server recipe.RecipesService_ListAllRecipesArrayServer) error {
	return server.Send(&recipe.ListAllRecipesResponseArray{Recipe: Recipes})
}

func (r repository) ListAllRecipes(request *recipe.ListAllRecipesRequest, server recipe.RecipesService_ListAllRecipesServer) error {
	log.Printf("Listing all available recipes")

	for _, r := range Recipes {
		err := server.Send(&recipe.ListAllRecipesResponse{Recipe: r})

		if err != nil {

		}
	}

	return nil
}

func (r repository) ListAllIngredientsAtHome(server recipe.RecipesService_ListAllIngredientsAtHomeServer) error {
	log.Printf("Noting all the ingredients that you have at home:")

	for {
		data, err := server.Recv()
		if err == io.EOF {
			return server.SendAndClose(&recipe.ListAllIngredientsAtHomeResponse{Success: true})
		}
		if err != nil {
			return err
		}

		log.Printf("You have %v quantity of %v", data.GetIngredient().GetQuantity(), data.GetIngredient().GetName())
	}
}

func (r repository) GetIngredientsForAllRecipes(server recipe.RecipesService_GetIngredientsForAllRecipesServer) error {
	log.Printf("For all recipes sent, I will reply back with a list of ingredients")

	for {
		r, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("You have requested ingredients for %v", r.GetRecipe().GetName())

		RecipeToIngredients := RecipeToIngredientsMap()
		ingredients := RecipeToIngredients[r.GetRecipe().GetName()]
		for _, item := range ingredients {
			server.Send(&recipe.GetIngredientsForAllRecipesResponse{Ingredient: &item})
		}
	}

	return nil
}

func NewSocialGRPCRepository(server *grpc.Server) {
	recipe.RegisterRecipesServiceServer(server, &repository{})
}
