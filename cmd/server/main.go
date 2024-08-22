/*package main

import (
	"log"
	"net"

	grpcPkg "todo-app/pkg/grpc"
	"todo-app/pkg/repository"

	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&repository.Todo{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	todoRepo := repository.NewGORMRepository(db)
	grpcServer := grpc.NewServer()
	todoService := service.NewTodoService(todoRepo)
	grpcPkg.RegisterTodoServiceServer(grpcServer, todoService)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}*/

package main

import (
	"log"
	"net"

	grpcPkg "todo-app/pkg/grpc"
	"todo-app/pkg/repository"

	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"

	// "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Replace with your SQL Server connection string
	// dsn := "sqlserver://sa:123@localhost:1433?database=playersss"
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&repository.Todo{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	todoRepo := repository.NewGORMRepository(db)
	grpcServer := grpc.NewServer()
	todoService := grpcPkg.NewTodoServiceServer(todoRepo)
	grpcPkg.RegisterTodoServiceServer(grpcServer, todoService)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
