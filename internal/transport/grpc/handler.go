package grpc

import (
	"context"

	userpb "github.com/arkad0912/project-protos/proto/user"
	"github.com/arkad0912/user-service/internal/userService"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	svc *userService.UserService
	userpb.UnimplementedUserServiceServer
}

func NewUserHandlers(svc *userService.UserService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user := &userService.User{
		Email: req.GetEmail(),
	}

	createdUser, err := h.svc.CreateUser(user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    uint32(createdUser.ID),
			Email: createdUser.Email,
		},
	}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.User) (*userpb.User, error) {
	user, err := h.svc.GetUserByID(uint(req.GetId()))
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &userpb.User{
		Id:    uint32(user.ID),
		Email: user.Email,
	}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	user := &userService.User{
		Email: req.GetNewEmail(),
	}

	updatedUser, err := h.svc.UpdateUser(uint(req.GetId()), user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:    uint32(updatedUser.ID),
			Email: updatedUser.Email,
		},
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.UpdateUserResponse, error) {
	err := h.svc.DeleteUser(uint(req.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id: req.GetId(),
		},
	}, nil
}

func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.GetUsers()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var pbUsers []*userpb.User
	for _, u := range users {
		pbUsers = append(pbUsers, &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
		})
	}

	return &userpb.ListUsersResponse{
		Users:      pbUsers,
		TotalCount: uint32(len(pbUsers)),
	}, nil
}
