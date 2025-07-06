package main

import (
	"fmt"
	"log"
	"net"

	"github.com/LavaJover/shvark-authz-service/internal/config"
	"github.com/LavaJover/shvark-authz-service/internal/delivery/grpcapi"
	"github.com/LavaJover/shvark-authz-service/internal/rbac"
	authzpb "github.com/LavaJover/shvark-authz-service/proto/gen"
	"google.golang.org/grpc"

	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load .env")
	}

	cfg := config.MustLoad()
	fmt.Println(cfg)

	lis, err := net.Listen("tcp", ":"+cfg.GRPCServer.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	enf := rbac.InitEnforcer(cfg)
	s := grpc.NewServer()
	authzpb.RegisterAuthzServiceServer(s, &grpcapi.AuthzService{Enforcer: enf})

	log.Printf("Authz service is running on :%s\n", cfg.GRPCServer.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}