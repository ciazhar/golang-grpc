package main

import (
	"context"
	"fmt"
	"github.com/ciazhar/golang-grpc/grpc/generated/go/recipe"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello client ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := recipe.NewRecipesServiceClient(cc)
	request := &recipe.ListAllRecipesRequest{}

	resp, _ := client.ListAllRecipes(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp)
}
