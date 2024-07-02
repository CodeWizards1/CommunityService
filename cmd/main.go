package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Projects/ComunityService/config"
	pb "github.com/Projects/ComunityService/genproto/CommunityService"
	"github.com/Projects/ComunityService/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
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
	lis, err := net.Listen("tcp", "localhost:50054")
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	gprcServer := grpc.NewServer()
	db, err := GetDB(".")

	if err != nil {
		log.Fatalf("Connecting to database failed: %v", err)
	}

	communityService := services.NewCommunityService(db)
	pb.RegisterCommunityServiceServer(gprcServer, communityService)

	log.Println("gRPC server is running on port 50054")
	if err := gprcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
