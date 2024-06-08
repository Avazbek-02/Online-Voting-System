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
	votingconection, err := grpc.NewClient(":50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error while voting :50050", err)
		return
	}
	defer votingconection.Close()

	publicConection, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":50051"), grpc.WithTransportCredentials(insecure.NewCredentials()))//conection
	if err != nil {
		log.Fatal("error whike :50051", err)
		return
	}
	defer publicConection.Close()

	public := pb.NewPublicServiceClient(publicConection)
	party := pb.NewPartyServiceClient(publicConection)
	canditate := pb.NewCandidateServiceClient(votingconection)
	election := pb.NewElectionServiceClient(votingconection)
	votes := pb.NewPublicVoteServiceClient(votingconection)


	h := handler.NewHandler(party, public,canditate,election,votes)
	router := api.NewEngine(h)

	fmt.Println("Server is run on :8080")

	err = router.Run()
	if err != nil{
		log.Fatal("Error while router",err)
		return
	}
}
