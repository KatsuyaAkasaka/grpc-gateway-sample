module github.com/KatsuyaAkasaka/grpc-gateway-sample

go 1.12

// ローカルファイルを参照する場合
replace github.com/KatsuyaAkasaka/grpc-gateway-sample => ./

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-gateway v1.13.0
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	google.golang.org/genproto v0.0.0-20190927181202-20e1ac93f88c
	google.golang.org/grpc v1.27.1
)
