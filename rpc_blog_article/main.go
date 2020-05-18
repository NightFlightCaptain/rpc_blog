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
	"rpc_blog_article/models"
	"rpc_blog_article/rpc"
	"rpc_blog_article/rpc/out"
)

func ZapLogger() *zap.Logger {
	file, _ := os.Create("test.log")
	writer := zapcore.AddSync(file)

	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewCore(encoder, writer, zapcore.InfoLevel)
	logger := zap.New(core)
	return logger
}

func RegisterMethod() {
	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(ZapLogger()),
		),
	}

	s := grpc.NewServer(opts...)
	out.RegisterArticleServiceServer(s, new(rpc.ArticleServer))

	lis, err := net.Listen("tcp", ":8082")
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
