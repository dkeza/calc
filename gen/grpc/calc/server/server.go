// Code generated by goa v3.0.8, DO NOT EDIT.
//
// calc gRPC server
//
// Command:
// $ goa gen calc/design

package server

import (
	calc "calc/gen/calc"
	calcpb "calc/gen/grpc/calc/pb"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the calcpb.CalcServer interface.
type Server struct {
	AddH goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the calc service endpoints.
func New(e *calc.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		AddH: NewAddHandler(e.Add, uh),
	}
}

// NewAddHandler creates a gRPC handler which serves the "calc" service "add"
// endpoint.
func NewAddHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAddRequest, EncodeAddResponse)
	}
	return h
}

// Add implements the "Add" method in calcpb.CalcServer interface.
func (s *Server) Add(ctx context.Context, message *calcpb.AddRequest) (*calcpb.AddResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "add")
	ctx = context.WithValue(ctx, goa.ServiceKey, "calc")
	resp, err := s.AddH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*calcpb.AddResponse), nil
}
