package gateway

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "github.com/KatsuyaAkasaka/grpc-gateway-sample/proto"
)

func run(serverPort string, gwPort string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf(":" + serverPort)
	err := gw.RegisterSayHelloHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}
	log.Printf("gateway port:" + gwPort)
	log.Printf("server listen: " + serverPort)
	return http.ListenAndServe(":"+gwPort, mux)
}

func Start(serverPort string, gwPort string) {
	flag.Parse()
	defer glog.Flush()
	if err := run(serverPort, gwPort); err != nil {
		glog.Fatal(err)
	}
}
