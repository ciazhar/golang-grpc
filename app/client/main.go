package main

import (
	"context"
	"fmt"
	"github.com/ciazhar/golang-grpc/grpc/generated/golang"
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

	client := golang.NewRecipesServiceClient(cc)
	request := &golang.ListAllRecipesRequest{}

	resp, _ := client.ListAllRecipes(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp)
}
