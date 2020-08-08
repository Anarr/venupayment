package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Anarr/venupayment/proto/customer"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type server struct {
}

func (s *server) New(cxt context.Context, req *customer.CustomerReq) (*customer.CustomerRes, error) {

	rand.Seed(time.Now().UnixNano())

	id := rand.Intn(9999999)

	idStr := strconv.Itoa(id)
	if req.Description != "" {
		return &customer.CustomerRes{
			Id: idStr,
		}, nil
	}

	return nil, errors.New("can not create new customer")
}

func main() {
	addr := ":50001"

	lis, err := net.Listen("tcp", addr)

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
