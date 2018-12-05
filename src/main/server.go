// grpc project main.go
package main

import (
	pb "github.com/Yq2/grpc_demo/src/greeter"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"runtime"
	"strconv"
)

const (
	PORT_CLIENT = "41005"
)

type Data struct{}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 起服务
	lis, err := net.Listen("tcp", ":"+PORT_CLIENT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDataServer(s, &Data{})
	serverErr := s.Serve(lis)
	if err != nil {
		log.Fatal(serverErr)
	}
	log.Println("grpc server in: %s", PORT_CLIENT)
}

// 定义方法
func (t *Data) GetUser(ctx context.Context, request *pb.UserRq) (response *pb.UserRp, err error) {
	//fmt.Println("request.Id:",request.Id)
	response = &pb.UserRp{
		Name: strconv.Itoa(int(request.Id)) + ":rpc_test",
	}
	return response, err
}
