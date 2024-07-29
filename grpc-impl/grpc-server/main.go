package main

import (
	"log"
	"net"

	pb "kodinggo/pb/story"
	grpcsvc "kodinggo/service"

	"google.golang.org/grpc"
)

const (
	host = "localhost"
	port = 50051
)

func main() {
	// Init grpc service & servce
	s := grpc.NewServer()
	svc := grpcsvc.NewService()
	pb.RegisterStoryServiceServer(s, svc)

	// Buat network listener
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Panicf("failed to create listener, error: %v", err)
	}

	log.Printf("listening on %d", port)

	// Run grpc server
	err = s.Serve(lis)
	if err != nil {
		log.Panicf("failed to run grpc server, error: %v", err)
	}
}
