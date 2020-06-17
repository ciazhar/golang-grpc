package grpc

import (
	"context"
	"github.com/ciazhar/golang-grpc/grpc/generated/golang"
	"google.golang.org/grpc"
	"io"
	"log"
)

type repository struct {
}

func (r repository) AddRecipe(ctx context.Context, request *golang.AddRecipeRequest) (*golang.AddRecipeResponse, error) {
	log.Printf("Adding new golang\t\tName: %v, Cuisine: %v", request.GetRecipe().GetName(), request.GetRecipe().GetCuisine())

	response := &golang.AddRecipeResponse{Success: true}

	return response, nil
}

func (r repository) ListAllRecipesArray(request *golang.ListAllRecipesRequest, server golang.RecipesService_ListAllRecipesArrayServer) error {
	return server.Send(&golang.ListAllRecipesResponseArray{Recipe: Recipes})
}

func (r repository) ListAllRecipes(request *golang.ListAllRecipesRequest, server golang.RecipesService_ListAllRecipesServer) error {
	log.Printf("Listing all available golangs")

	for _, r := range Recipes {
		err := server.Send(&golang.ListAllRecipesResponse{Recipe: r})

		if err != nil {

		}
	}

	return nil
}

func (r repository) ListAllIngredientsAtHome(server golang.RecipesService_ListAllIngredientsAtHomeServer) error {
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

func (r repository) GetIngredientsForAllRecipes(server golang.RecipesService_GetIngredientsForAllRecipesServer) error {
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

		RecipeToIngredients := RecipeToIngredientsMap()
		ingredients := RecipeToIngredients[r.GetRecipe().GetName()]
		for _, item := range ingredients {
			server.Send(&golang.GetIngredientsForAllRecipesResponse{Ingredient: &item})
		}
	}

	return nil
}

func NewSocialGRPCRepository(server *grpc.Server) {
	golang.RegisterRecipesServiceServer(server, &repository{})
}
