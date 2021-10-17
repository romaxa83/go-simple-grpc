package main

import (
	"context"
	"fmt"
	pb "github.com/romaxa83/go-simple-grpc/gen/proto"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewTestApiClient(conn)

	res, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Yell"})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(res)
	fmt.Println(res.Msg)
}
