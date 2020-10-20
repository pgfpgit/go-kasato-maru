package handlers

import (
	"context"
	"errors"
	"log"

	pb "todo"
	repository "todo/db"
)

var (
	ERROR_EMPTY_MESSAGE = errors.New("Texto não pode ser vazio")
	ERROR_FAIL_INSERT   = errors.New("Falha ao inserir item")
	ERROR_FAIL_LIST     = errors.New("Falha ao lista itens")
	ERROR_FAIL_REMOVE   = errors.New("Falha ao remover item")
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.TodoServer {
	return todoService{
		repository: repository.NewMongoDBRepository(),
	}
}

type todoService struct {
	repository repository.Repository
}

func (s todoService) AddItem(ctx context.Context, in *pb.AddItemRequest) (*pb.ItemResponse, error) {
	var resp pb.ItemResponse

	if in.Text == "" {
		return nil, ERROR_EMPTY_MESSAGE
	}

	err := s.repository.Add(*in)
	if err != nil {
		log.Printf("failed to insert item %s", err)
		return nil, ERROR_FAIL_INSERT
	}

	resp.Text = "Item inserido com sucesso"

	return &resp, nil
}

func (s todoService) DeleteItem(ctx context.Context, in *pb.DeleteItemRequest) (*pb.ItemResponse, error) {
	var resp pb.ItemResponse

	if err := s.repository.Delete(in.Id); err != nil {
		return &resp, ERROR_FAIL_REMOVE
	}

	resp.Text = "Item removido com sucesso"
	return &resp, nil
}

func (s todoService) ListAll(ctx context.Context, in *pb.EmptyRequest) (*pb.ListResponse, error) {
	elem, err := s.repository.FindAll()
	if err != nil {
		return nil, ERROR_FAIL_LIST
	}

	return &elem, nil
}
