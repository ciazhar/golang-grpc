package grpc

import (
	"context"
	repository2 "github.com/ciazhar/golang-grpc/app/server/pkg/recipe/repository"
	"github.com/ciazhar/golang-grpc/grpc/generated/golang"
	"google.golang.org/grpc"
	"io"
	"log"
)

type recipeController struct {
}

func (r recipeController) AddRecipe(ctx context.Context, request *golang.AddRecipeRequest) (*golang.AddRecipeResponse, error) {
	log.Printf("Adding new golang\t\tName: %v, Cuisine: %v", request.GetRecipe().GetName(), request.GetRecipe().GetCuisine())

	response := &golang.AddRecipeResponse{Success: true}

	return response, nil
}

func (r recipeController) ListAllRecipesArray(request *golang.ListAllRecipesRequest, server golang.RecipesService_ListAllRecipesArrayServer) error {
	return server.Send(&golang.ListAllRecipesResponseArray{Recipe: repository2.Recipes})
}

func (r recipeController) ListAllRecipes(request *golang.ListAllRecipesRequest, server golang.RecipesService_ListAllRecipesServer) error {
	log.Printf("Listing all available golangs")

	for _, r := range repository2.Recipes {
		err := server.Send(&golang.ListAllRecipesResponse{Recipe: r})

		if err != nil {

		}
	}

	return nil
}

func (r recipeController) ListAllIngredientsAtHome(server golang.RecipesService_ListAllIngredientsAtHomeServer) error {
	log.Printf("Noting all the ingredients that you have at home:")

	for {
		data, err := server.Recv()
		if err == io.EOF {
			return server.SendAndClose(&golang.ListAllIngredientsAtHomeResponse{Success: true})
		}
		if err != nil {
			return err
		}

		log.Printf("You have %v quantity of %v", data.GetIngredient().GetQuantity(), data.GetIngredient().GetName())
	}
}

func (r recipeController) GetIngredientsForAllRecipes(server golang.RecipesService_GetIngredientsForAllRecipesServer) error {
	log.Printf("For all golangs sent, I will reply back with a list of ingredients")

	for {
		r, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("You have requested ingredients for %v", r.GetRecipe().GetName())

		RecipeToIngredients := repository2.RecipeToIngredientsMap()
		ingredients := RecipeToIngredients[r.GetRecipe().GetName()]
		for _, item := range ingredients {
			server.Send(&golang.GetIngredientsForAllRecipesResponse{Ingredient: &item})
		}
	}

	return nil
}

func NewSocialGRPCController(server *grpc.Server) {
	golang.RegisterRecipesServiceServer(server, &recipeController{})
}
