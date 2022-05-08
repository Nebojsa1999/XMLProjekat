package services

import (
	postingPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/posting_service"
	userPb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewUserClient(address string) userPb.UserServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}

	return userPb.NewUserServiceClient(conn)
}

func NewPostingClient(address string) postingPb.PostingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Posting service: %v", err)
	}

	return postingPb.NewPostingServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
