package main

import (
	v "all/genproto/genproto/public"
	"all/service"
	"all/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	db, err := postgres.DBConnect()
	if err != nil {
		log.Fatalf("Error while connecting to Database! %v", err)
	}

	newServe := grpc.NewServer()

	v.RegisterPartyServiceServer(newServe, service.NewPartyService(db))
	v.RegisterPublicServiceServer(newServe, service.NewPublicService(db))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error while listening Server! %v", err)
	}

	if err := newServe.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
