package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "customer/api/customer/v1" // adjust import path if needed
)

func printJSON(v any) {
    b, _ := json.MarshalIndent(v, "", "  ")
    fmt.Println(string(b))
}

func main() {
    conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    client := pb.NewCustomerClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // 1. Create customer
    createResp, err := client.CreateCustomer(ctx, &pb.CreateCustomerReq{
        Name:         "Alice",
        DateOfBirth:  "2000-01-01",
    })
    if err != nil {
        log.Fatal("CreateCustomer:", err)
    }
    fmt.Println("CreateCustomer response:")
    printJSON(createResp)

    // 2. Get customer
    getResp, err := client.GetCustomer(ctx, &pb.GetCustomerReq{
        Id: createResp.Id,
    })
    if err != nil {
        log.Fatal("GetCustomer:", err)
    }
    fmt.Println("GetCustomer response:")
    printJSON(getResp)
}
