package store

import (
	"context"

	"github.com/rshulabs/micro-frame/internal/demo/store/pb"
)

func (i *Impl) Get(c context.Context, in *pb.Request) (*pb.Response, error) {
	return &pb.Response{Value: []byte("test")}, nil
}
