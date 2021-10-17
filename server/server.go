package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/romaxa83/go-simple-grpc/gen/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func main()  {

	go func() {
		// mux (Mультипле́ксор)
		mux := runtime.NewServeMux()
		// register
		pb.RegisterTestApiHandlerServer(context.Background(), mux, &testApiServer{})
		// http server
		log.Fatalln(http.ListenAndServe("localhost:8081", mux))
	}()

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}