package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Projects/ComunityService/config"
	pb "github.com/Projects/ComunityService/genproto/CommunityService"
	user "github.com/Projects/ComunityService/genproto/UserManagementService"

	"github.com/Projects/ComunityService/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetDB(path string) (*sqlx.DB, error) {
	cfg := config.Load(path)

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.DbHost,
		cfg.Postgres.DbPort,
		cfg.Postgres.DbUser,
		cfg.Postgres.DbPassword,
		cfg.Postgres.DbName,
	)

	db, err := sqlx.Connect("postgres", psqlUrl)
	return db, err
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50055")
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	db, err := GetDB(".")

	if err != nil {
		log.Fatalf("Connecting to database failed: %v", err)
	}

	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to User Management Service: %v", err)
	}
	defer conn.Close()
	userClient := user.NewUserManagementServiceClient(conn)

	communityService := services.NewCommunityService(db, userClient)
	pb.RegisterCommunityServiceServer(grpcServer, communityService)

	log.Println("gRPC server is running on port 50055")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
