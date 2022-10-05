package authservice

import (
	"context"
	"time"

	//"github.com/minh1611/go_structure/apiservice/src/pb"
)

type AuthServiceAddr string

var (
	_       Service = ServiceGRPC{}
	timeout         = 5 * time.Second
)

type ServiceGRPC struct {
	Ctx    context.Context
	//Client pb.
}

// func ConnectClient(ctx context.Context, addr AuthServiceAddr) (authpb.AuthServiceClient, error) {
// 	nctx, cancel := context.WithTimeout(ctx, timeout)
// 	defer cancel()
// 	conn, err := grpc.DialContext(nctx, string(addr), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(
// 		otelgrpc.UnaryClientInterceptor(),
// 	))
// 	if err != nil {
// 		return nil, xerrors.Errorf("%w", err)
// 	}
// 	return authpb.NewAuthServiceClient(conn), nil
// }