package main

import (
	"context"
	authpb "github.com/Alexx1088/authservice/proto"
	userpb "github.com/Alexx1088/userservice/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Connect to AuthService
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("failed to connect to AuthService: %v", err)
	}

	// Connect to UserService
	err = userpb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:50052", opts)
	if err != nil {
		log.Fatalf("failed to connect to UserService: %v", err)
	}

	log.Println("gateway server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
