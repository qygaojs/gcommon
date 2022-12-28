package tools

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
)

/**
 * 获取grpc的客户端addr
 */
func GrpcClientAddr(ctx context.Context) (string, error) {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("get grpc client address failed")
	}
	if pr.Addr == net.Addr(nil) {
		return "", fmt.Errorf("get grpc client address result is invalid")
	}
	return pr.Addr.String(), nil
}

/**
 * 获取grpc的客户端ip
 */
func GrpcClientIP(ctx context.Context) (string, error) {
	addr, err := GrpcClientAddr(ctx)
	if err != nil {
		return "", err
	}
	addrFields := strings.Split(addr, ":")
	return addrFields[0], nil
}

/**
 * 获取grpc的客户端port
 */
func GrpcClietPort(ctx context.Context) (int, error) {
	addr, err := GrpcClientAddr(ctx)
	if err != nil {
		return -1, err
	}
	addrFields := strings.Split(addr, ":")
	if len(addrFields) < 2 {
		return -1, fmt.Errorf("get grpc client address result is invalid")
	}
	return strconv.Atoi(addrFields[1])
}
