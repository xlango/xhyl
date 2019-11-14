package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "xhyl/proto/model"
)

const (
	address = "192.168.10.250:50001"
)

func main() {
	http.HandleFunc("/grpc/client/test", testClient)
	http.ListenAndServe(fmt.Sprintf(":%d", 7001), nil)
}

func testClient(writer http.ResponseWriter, request *http.Request) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	name := "xhyl"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(r.Message)
	writer.Write([]byte(r.Message))
}
