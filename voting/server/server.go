package main

import (
	"log"
	"net"

	"voting/storage/postgres"
	v "voting/genproto/genproto/voting"
	"voting/service"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.DBConnect()
	if err != nil {
		log.Fatalf("Error while connecting to Database! %v", err)
	}

	newServe := grpc.NewServer()

	v.RegisterCandidateServiceServer(newServe, service.NewCandidateService(db))
	v.RegisterElectionServiceServer(newServe, service.NewElectionService(db))
	v.RegisterPublicVoteServiceServer(newServe, service.NewPublicVoteService(db))

	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("Error while listening Server! %v", err)
	}

	if err := newServe.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
