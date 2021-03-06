package server

import (
	"log"
	"net"

	pb "github.com/KatsuyaAkasaka/grpc-gateway-sample/proto"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

type helloService struct{}

func (hs *helloService) Echo(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "hello, " + req.UserName,
	}, nil
}

func Start(port string) {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("server listen: " + listen.Addr().String())
	server := grpc.NewServer()
	pb.RegisterSayHelloServer(server, &helloService{})

	if err := server.Serve(listen); err != nil {
		log.Fatalln(err)
	}
	return
}
