package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	grpcPkg "todo-app/pkg/grpc"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcPkg.NewTodoServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Create Todo")
		fmt.Println("2. Get Todo")
		fmt.Println("3. List Todos")
		fmt.Println("4. Update Todo")
		fmt.Println("5. Delete Todo")
		fmt.Println("6. Exit")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
		defer cancel()

		switch choice {
		case "1":
			fmt.Print("Enter title: ")
			title, _ := reader.ReadString('\n')
			fmt.Print("Enter description: ")
			description, _ := reader.ReadString('\n')
			createReq := &grpcPkg.CreateTodoRequest{
				Title:       strings.TrimSpace(title),
				Description: strings.TrimSpace(description),
			}
			createRes, err := client.CreateTodo(ctx, createReq)
			if err != nil {
				log.Fatalf("could not create todo: %v", err)
			}
			fmt.Printf("Created Todo: %v\n", createRes.Todo)
		case "2":
			fmt.Print("Enter Todo ID: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64)
			if err != nil {
				log.Fatalf("invalid ID: %v", err)
			}
			getReq := &grpcPkg.GetTodoRequest{Id: id}
			getRes, err := client.GetTodo(ctx, getReq)
			if err != nil {
				log.Fatalf("could not get todo: %v", err)
			}
			fmt.Printf("Todo: %v\n", getRes.Todo)
		case "3":
			listReq := &grpcPkg.ListTodosRequest{}
			listRes, err := client.ListTodos(ctx, listReq)
			if err != nil {
				log.Fatalf("could not list todos: %v", err)
			}
			fmt.Printf("List of Todos: %v\n", listRes.Todos)
		case "4":
			fmt.Print("Enter Todo ID: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64)
			if err != nil {
				log.Fatalf("invalid ID: %v", err)
			}
			fmt.Print("Enter new title: ")
			title, _ := reader.ReadString('\n')
			fmt.Print("Enter new description: ")
			description, _ := reader.ReadString('\n')
			fmt.Print("Is it completed? (true/false): ")
			completedStr, _ := reader.ReadString('\n')
			completed, err := strconv.ParseBool(strings.TrimSpace(completedStr))
			if err != nil {
				log.Fatalf("invalid completed value: %v", err)
			}
			updateReq := &grpcPkg.UpdateTodoRequest{
				Id:          id,
				Title:       strings.TrimSpace(title),
				Description: strings.TrimSpace(description),
				Completed:   completed,
			}
			updateRes, err := client.UpdateTodo(ctx, updateReq)
			if err != nil {
				log.Fatalf("could not update todo: %v", err)
			}
			fmt.Printf("Updated Todo: %v\n", updateRes.Todo)
		case "5":
			fmt.Print("Enter Todo ID: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64)
			if err != nil {
				log.Fatalf("invalid ID: %v", err)
			}
			deleteReq := &grpcPkg.DeleteTodoRequest{Id: id}
			_, err = client.DeleteTodo(ctx, deleteReq)
			if err != nil {
				log.Fatalf("could not delete todo: %v", err)
			}
			fmt.Println("Deleted Todo")
		case "6":
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
