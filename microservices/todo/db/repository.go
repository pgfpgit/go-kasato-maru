package repository

import pb "todo"

type Repository interface {
	Add(element pb.AddItemRequest) error
	Delete(id string) error
	FindAll() (pb.ListResponse, error)
}
