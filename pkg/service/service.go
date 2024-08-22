package service

import (
    "context"
    "todo-app/pkg/grpc"
    "todo-app/pkg/repository"

    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type TodoService struct {
    repo repository.Repository
    grpc.UnimplementedTodoServiceServer
}

func NewTodoService(repo repository.Repository) *TodoService {
    return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *grpc.CreateTodoRequest) (*grpc.TodoResponse, error) {
    todo := &repository.Todo{
        Title	:       req.Title,
        Description: req.Description,
    }
    if err := s.repo.Create(todo); err != nil {
        return nil, status.Errorf(codes.Internal, "failed to create todo: %v", err)
    }
    return &grpc.TodoResponse{Todo: &grpc.Todo{
        Id:          todo.ID,
        Title:       todo.Title,
        Description: todo.Description,
        Completed:   todo.Completed,
    }}, nil
}

func (s *TodoService) GetTodo(ctx context.Context, req *grpc.GetTodoRequest) (*grpc.TodoResponse, error) {
    todo, err := s.repo.Get(req.Id)
    if err != nil {
        return nil, status.Errorf(codes.NotFound, "todo not found: %v", err)
    }
    return &grpc.TodoResponse{Todo: &grpc.Todo{
        Id:          todo.ID,
        Title:       todo.Title,
        Description: todo.Description,
        Completed:   todo.Completed,
    }}, nil
}

func (s *TodoService) ListTodos(ctx context.Context, req *grpc.ListTodosRequest) (*grpc.ListTodosResponse, error) {
    todos, err := s.repo.List()
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to list todos: %v", err)
    }
    var grpcTodos []*grpc.Todo
    for _, todo := range todos {
        grpcTodos = append(grpcTodos, &grpc.Todo{
            Id:          todo.ID,
            Title:       todo.Title,
            Description: todo.Description,
            Completed:   todo.Completed,
        })
    }
    return &grpc.ListTodosResponse{Todos: grpcTodos}, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, req *grpc.UpdateTodoRequest) (*grpc.TodoResponse, error) {
    todo := &repository.Todo{
        ID:          req.Id,
        Title:       req.Title,
        Description: req.Description,
        Completed:   req.Completed,
    }
    if err := s.repo.Update(todo); err != nil {
        return nil, status.Errorf(codes.Internal, "failed to update todo: %v", err)
    }
    return &grpc.TodoResponse{Todo: &grpc.Todo{
        Id:          todo.ID,
        Title:       todo.Title,
        Description: todo.Description,
        Completed:   todo.Completed,
    }}, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, req *grpc.DeleteTodoRequest) (*grpc.EmptyResponse, error) {
    if err := s.repo.Delete(req.Id); err != nil {
        return nil, status.Errorf(codes.Internal, "failed to delete todo: %v", err)
    }
    return &grpc.EmptyResponse{}, nil
}
