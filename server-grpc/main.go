package main

import (
	"flag"
	"fmt"
	"net"

	"langriver_service/internal/config"
	"langriver_service/internal/server"
	"langriver_service/internal/svc"
	pb "langriver_service/proto"

	"github.com/zeromicro/go-zero/core/conf"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/chatbot.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	grpcServer := grpc.NewServer()
	pb.RegisterChatbotServiceServer(grpcServer, server.NewChatbotServer(ctx))

	lis, err := net.Listen("tcp", c.ListenOn)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Starting gRPC server at %s...\n", c.ListenOn)
	grpcServer.Serve(lis)
}
