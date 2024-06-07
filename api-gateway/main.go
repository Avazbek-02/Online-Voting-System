package main

import (
	"api/api/handler"
	"api/api"
	pb "api/genproto"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// voting, err := grpc.NewClient(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatal("error while voting :9090", err)
	// 	return
	// }

	publicserer, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":9099"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error whike :9099", err)
		return
	}
	defer publicserer.Close()

	public := pb.NewPublicServiceClient(publicserer)
	party := pb.NewPartyServiceClient(publicserer)


	h := handler.NewHandler(party, public)
	router := api.NewEngine(h)

	fmt.Println("Server is run on :8080")
	
	err = router.Run()
	if err != nil{
		log.Fatal("Error while router",err)
		return
	}
}
