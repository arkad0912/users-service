package grpc

import (
	"log"
	"net"

	userpb "github.com/arkad0912/project-protos/proto/user"
	"google.golang.org/grpc"
)

func RunServer(handler *Handler, addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, handler)

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
