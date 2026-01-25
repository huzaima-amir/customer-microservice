package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "time"

    "google.golang.org/grpc"
    pb "customer/api/customer/v1"
)
// FIX biz layer logic - issues with calling repo layer - emails, addresses and phone numbers are being treated as embedded fields instead of separate tables leading to errors in transaction
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

func writeJSON(filename string, v any) {
    b, _ := json.MarshalIndent(v, "", "  ")
    _ = os.WriteFile(filename, b, 0644)
}

func main() {
    // Connect to gRPC server
    conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())  // strikethrough - FIX
    must(err)
    defer conn.Close()

    client := pb.NewCustomerClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Create Customer
    createResp, err := client.CreateCustomer(ctx, &pb.CreateCustomerReq{
        Name:         "Alice",
        DateOfBirth:  "2000-01-01",
    })
    must(err)
    printJSON("CreateCustomer", createResp)
    writeJSON("create_customer.json", createResp)

    customerID := createResp.Id

    // Add Email
    emailResp, err := client.AddEmail(ctx, &pb.AddEmailReq{
        CustomerId: customerID,
        Email:      "alice@example.com",
    })
    must(err)
    printJSON("AddEmail", emailResp)
    writeJSON("add_email.json", emailResp)

    //Add Phone Number
    phoneResp, err := client.AddPhoneNumber(ctx, &pb.AddPhoneNumberReq{
        CustomerId:  customerID,
        PhoneNumber: "123456789",
    })
    must(err)
    printJSON("AddPhoneNumber", phoneResp)
    writeJSON("add_phone.json", phoneResp)

    // add Address
    addrResp, err := client.AddAddress(ctx, &pb.AddAddressReq{
        CustomerId: customerID,
        Address:    "123 Main St",
    })
    must(err)
    printJSON("AddAddress", addrResp)
    writeJSON("add_address.json", addrResp)

    // get Customer (with nested data)
    getResp, err := client.GetCustomer(ctx, &pb.GetCustomerReq{
        Id: customerID,
    })
    must(err)
    printJSON("GetCustomer", getResp)
    writeJSON("get_customer.json", getResp)

    // list Emails
    listEmailsResp, err := client.ListEmail(ctx, &pb.ListEmailReq{
        CustomerId: customerID,
    })
    must(err)
    printJSON("ListEmails", listEmailsResp)
    writeJSON("list_emails.json", listEmailsResp)

    //lit Phone Numbers
    listPhonesResp, err := client.ListPhoneNumber(ctx, &pb.ListPhoneNumberReq{
        CustomerId: customerID,
    })
    must(err)
    printJSON("ListPhoneNumbers", listPhonesResp)
    writeJSON("list_phones.json", listPhonesResp)

    //List Addresses
    listAddrResp, err := client.ListAddress(ctx, &pb.ListAddressReq{
        CustomerId: customerID,
    })
    must(err)
    printJSON("ListAddresses", listAddrResp)
    writeJSON("list_addresses.json", listAddrResp)

    // List All Customers
    listCustResp, err := client.ListCustomer(ctx, &pb.ListCustomerReq{})
    must(err)
    printJSON("ListCustomer", listCustResp)
    writeJSON("list_customers.json", listCustResp)

    // Delete Email
    delEmailResp, err := client.DeleteEmail(ctx, &pb.DeleteEmailReq{
        CustomerId: customerID,
        Email:      "alice@example.com",
    })
    must(err)
    printJSON("DeleteEmail", delEmailResp)
    writeJSON("delete_email.json", delEmailResp)

    // Delete Phone Number
    delPhoneResp, err := client.DeletePhoneNumber(ctx, &pb.DeletePhoneNumberReq{
        CustomerId: customerID,
        PhoneNumber: "123456789",
    })
    must(err)
    printJSON("DeletePhoneNumber", delPhoneResp)
    writeJSON("delete_phone.json", delPhoneResp)

    ///delete Address
    delAddrResp, err := client.DeleteAddress(ctx, &pb.DeleteAddressReq{
        CustomerId: customerID,
        Address:    "123 Main St",
    })
    must(err)
    printJSON("DeleteAddress", delAddrResp)
    writeJSON("delete_address.json", delAddrResp)

    //Delete Customer
    delCustResp, err := client.DeleteCustomer(ctx, &pb.DeleteCustomerReq{
        Id: customerID,
    })
    must(err)
    printJSON("DeleteCustomer", delCustResp)
    writeJSON("delete_customer.json", delCustResp)

    fmt.Println("==== ALL TESTS COMPLETED SUCCESSFULLY ====")
}
