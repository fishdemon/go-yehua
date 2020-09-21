package main

import (
	"fmt"
	"github.com/fishdemon/go-yehua/grpc/proto"
	"github.com/fishdemon/go-yehua/grpc/server/service"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":12345"
)

// 非流式拦截器
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("before handling unary...", info)
	resp, err = handler(ctx, req)
	log.Println("end handling unary..." , resp)
	return resp, err
}

// 流式拦截器 ，只有流式传送数据才生效
func StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("before handling stream...", info)
	err := handler(srv, ss)
	log.Println("end handling stream...")
	return err
}

// 恢复拦截器
func RecoverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		// 发生宕机时，强制恢复
		err := recover()
		fmt.Println(err)

		resp = &proto.CommonResponse{
			Code: -1,
			Msg: "fail",
		}
	}()
	resp, err = handler(ctx, req)
	return resp, err
}

// 容器拦截器(另外一种写法)
func ContextInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// context 中绑定日志组件
		sugerLogger := logger.Sugar()
		context.WithValue(ctx, "logger", sugerLogger)

		resp, err = handler(ctx, req)
		return resp, err
	}
}

// 日志拦截器


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 初始化日志组件
	logger, _ := zap.NewProduction()

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			UnaryServerInterceptor,
			RecoverInterceptor,
			ContextInterceptor(logger)),
		grpc.StreamInterceptor(StreamServerInterceptor))
	proto.RegisterUserServiceServer(s, service.NewUserService())
	fmt.Println("Server is starting!")
	s.Serve(lis)

}
