package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "kodinggo/pb/story"
)

const (
	storyHost = "localhost:50051"
)

func main() {
	// Buat koneksi ke grpc server tanpa credential
	conn, err := grpc.NewClient(storyHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	// Init grpc stub
	c := pb.NewStoryServiceClient(conn)

	// Panggil method FindAll untuk mengirim request ke grpc server
	r, err := c.FindAll(context.Background(), &pb.FindAllRequest{Page: 1, Limit: 10})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for _, st := range r.Stories {
		fmt.Println("ID", st.Id)
		fmt.Println("Title", st.Title)
	}
}
