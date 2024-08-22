package grpc

import (
    "context"
    "todo-app/pkg/repository"

    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type TodoServiceServer1 struct {
    repo repository.Repository
    UnimplementedTodoServiceServer
}

func NewTodoServiceServer(repo repository.Repository) *TodoServiceServer1 {
    return &TodoServiceServer1{repo: repo}
}

func (s *TodoServiceServer1) CreateTodo(ctx context.Context, req *CreateTodoRequest) (*TodoResponse, error) {
    todo := &repository.Todo{
        Title:       req.Title,
        Description: req.Description,
    }
    if err := s.repo.Create(todo); err != nil {
        return nil, status.Errorf(codes.Internal, "failed to create todo: %v", err)
    }
    return &TodoResponse{Todo: &Todo{
        Id:          todo.ID,
        Title:       todo.Title,
        Description: todo.Description,
        Completed:   todo.Completed,
    }}, nil
}

func (s *TodoServiceServer1) GetTodo(ctx context.Context, req *GetTodoRequest) (*TodoResponse, error) {
    todo, err := s.repo.Get(req.Id)
    if err != nil {
        return nil, status.Errorf(codes.NotFound, "todo not found: %v", err)
    }
    return &TodoResponse{Todo: &Todo{
        Id:          todo.ID,
        Title:       todo.Title,
        Description: todo.Description,
        Completed:   todo.Completed,
    }}, nil
}

func (s *TodoServiceServer1) ListTodos(ctx context.Context, req *ListTodosRequest) (*ListTodosResponse, error) {
    todos, err := s.repo.List()
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to list todos: %v", err)
    }
    var grpcTodos []*Todo
    for _, todo := range todos {
        grpcTodos = append(grpcTodos, &Todo{
            Id:          todo.ID,
            Title:       todo.Title,
            Description: todo.Description,
            Completed:   todo.Completed,
        })
    }
    return &ListTodosResponse{Todos: grpcTodos}, nil
}

func (s *TodoServiceServer1) UpdateTodo(ctx context.Context, req *UpdateTodoRequest) (*TodoResponse, error) {
    todo := &repository.Todo{
        ID:          req.Id,
        Title:       req.Title,
        Description: req.Description,
        Completed:   req.Completed,
    }
    if err := s.repo.Update(todo); err != nil {
        return nil, status.Errorf(codes.Internal, "failed to update todo: %v", err)
    }
    return &TodoResponse{Todo: &Todo{
        Id:          todo.ID,
        Title:       todo.Title,
        Description: todo.Description,
        Completed:   todo.Completed,
    }}, nil
}

func (s *TodoServiceServer1) DeleteTodo(ctx context.Context, req *DeleteTodoRequest) (*EmptyResponse, error) {
    if err := s.repo.Delete(req.Id); err != nil {
        return nil, status.Errorf(codes.Internal, "failed to delete todo: %v", err)
    }
    return &EmptyResponse{}, nil
}
