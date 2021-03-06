// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: ef2331b7e2
// Version Date: 2020-10-07T23:22:38Z

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	pb "todo"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	AddItemEndpoint    endpoint.Endpoint
	DeleteItemEndpoint endpoint.Endpoint
	ListAllEndpoint    endpoint.Endpoint
}

// Endpoints

func (e Endpoints) AddItem(ctx context.Context, in *pb.AddItemRequest) (*pb.ItemResponse, error) {
	response, err := e.AddItemEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.ItemResponse), nil
}

func (e Endpoints) DeleteItem(ctx context.Context, in *pb.DeleteItemRequest) (*pb.ItemResponse, error) {
	response, err := e.DeleteItemEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.ItemResponse), nil
}

func (e Endpoints) ListAll(ctx context.Context, in *pb.EmptyRequest) (*pb.ListResponse, error) {
	response, err := e.ListAllEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.ListResponse), nil
}

// Make Endpoints

func MakeAddItemEndpoint(s pb.TodoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.AddItemRequest)
		v, err := s.AddItem(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeDeleteItemEndpoint(s pb.TodoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.DeleteItemRequest)
		v, err := s.DeleteItem(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeListAllEndpoint(s pb.TodoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.EmptyRequest)
		v, err := s.ListAll(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"AddItem":    {},
		"DeleteItem": {},
		"ListAll":    {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "AddItem" {
			e.AddItemEndpoint = middleware(e.AddItemEndpoint)
		}
		if inc == "DeleteItem" {
			e.DeleteItemEndpoint = middleware(e.DeleteItemEndpoint)
		}
		if inc == "ListAll" {
			e.ListAllEndpoint = middleware(e.ListAllEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"AddItem":    {},
		"DeleteItem": {},
		"ListAll":    {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "AddItem" {
			e.AddItemEndpoint = middleware("AddItem", e.AddItemEndpoint)
		}
		if inc == "DeleteItem" {
			e.DeleteItemEndpoint = middleware("DeleteItem", e.DeleteItemEndpoint)
		}
		if inc == "ListAll" {
			e.ListAllEndpoint = middleware("ListAll", e.ListAllEndpoint)
		}
	}
}
