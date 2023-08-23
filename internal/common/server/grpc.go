package server

import (
	"fmt"
	"net"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus" // TODO github.com/grpc-ecosystem/go-grpc-middleware/v2 https://github.com/grpc-ecosystem/go-grpc-middleware/blob/main/interceptors/logging/examples/logrus/example_test.go
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func RunGRPCServer(registerServer func(server *grpc.Server)) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	grpcEndpoint := fmt.Sprintf(":%s", port)

	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
		),
		grpc.ChainStreamInterceptor(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(logrusEntry),
		),
	)
	registerServer(grpcServer)

	listen, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		logrus.Fatal(err) // TODO ??? fatal?!!!
	}
	logrus.WithField("grpcEndpoint", grpcEndpoint).Info("Starting: gRPC Listener")
	logrus.Fatal(grpcServer.Serve(listen)) // TODO ??? em... fatal again. what about returning error
}