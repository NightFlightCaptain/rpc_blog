package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"rpc_blog_article/conf"
	"rpc_blog_article/consul"
	"rpc_blog_article/models"
	"rpc_blog_article/rpc"
	"rpc_blog_article/rpc/out"
	"strconv"
	"time"
)

func ZapLogger() *zap.Logger {
	file, _ := os.Create("test.log")
	writer := zapcore.AddSync(file)

	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewCore(encoder, writer, zapcore.InfoLevel)
	logger := zap.New(core)
	return logger
}

func consulRegister() {
	register := consul.NewConsulRegister("localhost:8500", 15)
	register.Register(consul.RegisterInfo{
		Host:           "localhost",
		Port:           conf.Config.Server.HTTPPort,
		ServiceName:    conf.Config.Server.Name,
		UpdateInterval: time.Second,
	})
}

func RegisterMethod() {
	consulRegister()

	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(ZapLogger()),
		),
	}

	s := grpc.NewServer(opts...)
	out.RegisterArticleServiceServer(s, new(rpc.ArticleServer))

	lis, err := net.Listen("tcp", "localhost:"+strconv.Itoa(conf.Config.Server.HTTPPort))
	if err != nil {
		log.Fatal(err)
	}
	s.Serve(lis)
}

func main() {
	conf.SetUp()
	models.SetUp()
	rpc.SetUp()

	RegisterMethod()
}
