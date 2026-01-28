package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "customer/api/customer/v1"
)

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printJSON(label string, v any) {
	fmt.Println("==== " + label + " ====")
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}

func main() {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	must(err)
	defer conn.Close()

	client := pb.NewCustomerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()


	// Create Customers

	alice, err := client.CreateCustomer(ctx, &pb.CreateCustomerReq{
		Name:        "Alice",
		DateOfBirth: "2000-01-01",
	})
	must(err)
	printJSON("CreateCustomer Alice", alice)

	bob, err := client.CreateCustomer(ctx, &pb.CreateCustomerReq{
		Name:        "Bob",
		DateOfBirth: "1995-06-15",
	})
	must(err)
	printJSON("CreateCustomer Bob", bob)

	charlie, err := client.CreateCustomer(ctx, &pb.CreateCustomerReq{
		Name:        "Charlie",
		DateOfBirth: "1988-12-09",
	})
	must(err)
	printJSON("CreateCustomer Charlie", charlie)


	// Add Emails

	mustAddEmail := func(cid int64, email string) {
		_, err := client.AddEmail(ctx, &pb.AddEmailReq{
			CustomerId: cid,
			Email:      email,
		})
		must(err)
	}

	mustAddEmail(alice.Id, "alice@example.com")
	mustAddEmail(alice.Id, "alice.work@example.com")
	mustAddEmail(bob.Id, "bob@example.com")
	mustAddEmail(charlie.Id, "charlie@example.com")

	// Add Phone Numbers

	mustAddPhone := func(cid int64, phone string) {
		_, err := client.AddPhoneNumber(ctx, &pb.AddPhoneNumberReq{
			CustomerId:  cid,
			PhoneNumber: phone,
		})
		must(err)
	}

	mustAddPhone(alice.Id, "111111111")
	mustAddPhone(bob.Id, "222222222")
	mustAddPhone(charlie.Id, "333333333")
	mustAddPhone(charlie.Id, "444444444")


	// Add Addresses

	mustAddAddress := func(cid int64, addr string) {
		_, err := client.AddAddress(ctx, &pb.AddAddressReq{
			CustomerId: cid,
			Address:    addr,
		})
		must(err)
	}

	mustAddAddress(alice.Id, "123 Main St")
	mustAddAddress(bob.Id, "456 King Rd")
	mustAddAddress(charlie.Id, "789 Queen Ave")


	// Query: GetCustomer
	
	aliceGet, err := client.GetCustomer(ctx, &pb.GetCustomerReq{
		Id: alice.Id,
	})
	must(err)
	printJSON("GetCustomer Alice", aliceGet)

	
	// Query: GetCustomerByEmail
	
	byEmail, err := client.GetCustomerByEmail(ctx, &pb.GetCustomerByEmailReq{
		Email: "alice.work@example.com",
	})
	must(err)
	printJSON("GetCustomerByEmail alice.work@example.com", byEmail)

	
	// Query: GetCustomerByPhoneNumber
	
	byPhone, err := client.GetCustomerByPhoneNumber(ctx, &pb.GetCustomerByPhoneNumberReq{
		PhoneNumber: "333333333",
	})
	must(err)
	printJSON("GetCustomerByPhoneNumber 333333333", byPhone)

	
	// List per-customer details

	aliceEmails, err := client.ListEmail(ctx, &pb.ListEmailReq{
		CustomerId: alice.Id,
	})
	must(err)
	printJSON("ListEmails Alice", aliceEmails)

	charliePhones, err := client.ListPhoneNumber(ctx, &pb.ListPhoneNumberReq{
		CustomerId: charlie.Id,
	})
	must(err)
	printJSON("ListPhoneNumbers Charlie", charliePhones)

	
	// List all customers
	
	allCustomers, err := client.ListCustomer(ctx, &pb.ListCustomerReq{})
	must(err)
	printJSON("ListCustomer (ALL)", allCustomers)

	fmt.Println("==== MULTI-CUSTOMER QUERY TEST COMPLETED SUCCESSFULLY ====")
}
