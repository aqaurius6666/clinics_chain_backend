package authservice

import (
	"context"
	"fmt"
	"time"

	"github.com/minh1611/clinics_chain_management/apiservice/src/pb/authpb"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type AuthServiceAddr string

var (
	_       Service = ServiceGRPC{}
	timeout         = 5 * time.Second
)

type ServiceGRPC struct {
	Ctx    context.Context
	Client authpb.AuthServiceClient
}

func ConnectClient(ctx context.Context, addr AuthServiceAddr) (authpb.AuthServiceClient, error) {
	nctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	conn, err := grpc.DialContext(nctx, string(addr), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	} else {
		fmt.Println("CONNECT AUTHSERVICE SUCCESSFULLY")
	}
	return authpb.NewAuthServiceClient(conn), nil
}

func (s ServiceGRPC) CreateNewUser(ctx context.Context, name string, age int) (*authpb.User, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	res, err := s.Client.CreateNewUser(ctx, &authpb.NewUser{
		Name: name,
		Age:  int32(age),
	})
	if err != nil {
		stt, ok := status.FromError(err)
		if ok {
			err = xerrors.New(stt.Message())
		}
		return nil, err
	}
	return res, nil
}
