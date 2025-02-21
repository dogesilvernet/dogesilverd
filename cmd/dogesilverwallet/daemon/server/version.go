package server

import (
	"context"
	"github.com/dogesilvernet/dogesilverd/cmd/dogesilverwallet/daemon/pb"
	"github.com/dogesilvernet/dogesilverd/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
