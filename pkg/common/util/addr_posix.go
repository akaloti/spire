//go:build !windows

package util

import (
	"context"
	"errors"
	"net"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GRPCDialContext(ctx context.Context, target string, options ...grpc.DialOption) (*grpc.ClientConn, error) {
	options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpc.DialContext(ctx, target, options...) //nolint: staticcheck // It is going to be resolved on #5152
}

func GetWorkloadAPIClientOption(addr net.Addr) (workloadapi.ClientOption, error) {
	if _, ok := addr.(*net.UnixAddr); !ok {
		return nil, errors.New("address does not represent a Unix domain socket endpoint")
	}
	target, err := GetTargetName(addr)
	if err != nil {
		return nil, err
	}
	return workloadapi.WithAddr(target), nil
}
