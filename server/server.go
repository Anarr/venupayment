package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"log"
	"github.com/Anarr/venupayment/proto/customer"
)

type server struct {

}

func main() {
	addr := ":50001"

	lis, err := net.Listen("tco", addr)

	check(err)

	fmt.Printf("Server is listening on %s\n", addr)

	s := grpc.NewServer()

	customer.RegisterCustomerServiceServer(s, &server{})

	err = s.Serve(lis)

	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}