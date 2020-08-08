package main

import (
	"context"
	"fmt"
	"github.com/Anarr/venupayment/proto/customer"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Customer Client.....")

	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())

	check(err)

	defer conn.Close()

	client := customer.NewCustomerServiceClient(conn)

	req := &customer.CustomerReq{
		Description: "My new customer",
	}

	resp, err := client.New(context.Background(), req)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Resp from server", resp)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
