package main

import (
	"database/sql"
	"log"
	"net"

	"user-service/db/migrations"
	"user-service/handlers"
	pb "user-service/proto"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/ecommerce")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	migrations.Migrate(db)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &handlers.UserServiceServer{DB: db})

	log.Println("User Service is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
