package grpc

import (
	userpb "github.com/arkad0912/project-protos/proto/user"
)

type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}
