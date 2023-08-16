package store

import "github.com/rshulabs/micro-frame/internal/demo/store/pb"

type Impl struct {
	pb.UnimplementedDemoServer
}
