package main

import (
	"fmt"
	"log"
	"net"
	"userManagement/config"
	pb "userManagement/genproto/UserManagementService"

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
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	gprcServer := grpc.NewServer()

	db, err := GetDB(".")
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	userManaementService := service.NewuserManagementService(db)
	pb.RegisterUserManagementServiceServer(gprcServer, userManaementService)

	log.Println("gRPC server is running on port 50051")
	if err := gprcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
