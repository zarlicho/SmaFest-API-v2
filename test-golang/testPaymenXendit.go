package main

import (
	"fmt"
	"log"
	"time"
	"math/rand"
	"context"
	"github.com/xendit/xendit-go/invoice"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/transaction"
)
func GenerateOrderIDs() string {
    rand.Seed(time.Now().UnixNano()) // Set seed dari waktu sekarang agar randomizer selalu berubah

    const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // Set karakter yang mungkin digunakan pada order ID
    const idLength = 10                                       // Panjang order ID yang diinginkan

    var result string
    for i := 0; i < idLength; i++ {
        randomChar := string(charset[rand.Intn(len(charset))]) // Ambil karakter random dari charset
        result += randomChar                                   // Tambahkan karakter tersebut ke dalam order ID
    }

    return result
}

func ExampleCreate() {
	xendit.Opt.SecretKey = "xnd_development_KJqZ5c5Q2skG8FiitUhyFx5kbjM9V9V5SrWd4EkAysPuLzIfZlAuvt5s36qi"

	data := invoice.CreateParams{
		ExternalID:  GenerateOrderIDs(),
		// ForUserID: GenerateOrderIDs(),
		Amount:      100000,
		PayerEmail:  "hitech7261@gmail.com",
		Description: "invoice #5",
		SuccessRedirectURL: "https://www.youtube.com/watch?v=5Bkub_BYZvU",
		Customer: xendit.InvoiceCustomer{
			GivenNames: "John",
			Surname:    "china",
			MobileNumber:  "085771014979",
			Email: "hitech7261@gmail.com",
			
		},
	}


	resp, err := invoice.Create(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created invoice: %+v\n", resp.InvoiceURL)
}

func ExampleGet() {
	xendit.Opt.SecretKey = "xnd_development_KJqZ5c5Q2skG8FiitUhyFx5kbjM9V9V5SrWd4EkAysPuLzIfZlAuvt5s36qi"

	data := invoice.GetParams{
		ID: "4ZUVECDR3H",
	}

	resp, err := invoice.Get(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved invoice: %+v\n", resp.Status)
}

func ExampleGetListTransaction() {
	xendit.Opt.SecretKey = "xnd_development_KJqZ5c5Q2skG8FiitUhyFx5kbjM9V9V5SrWd4EkAysPuLzIfZlAuvt5s36qi"

	refID := "WYQ3NB6F4U"

	resp, err := transaction.GetListTransactionnWithContext(context.Background(), &transaction.GetListTransactionParams{
		ReferenceID: refID,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("transaction detail: %+v\n", resp.Data[0])
}

func ExampleExpire() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := invoice.ExpireParams{
		ID: "invoice-id",
	}

	resp, err := invoice.Expire(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("expired invoice: %+v\n", resp.Status)
}

func ExampleGetAll() {
	xendit.Opt.SecretKey = "examplesecretkey"

	createdAfter, _ := time.Parse(time.RFC3339, "2016-02-24T23:48:36.697Z")

	data := invoice.GetAllParams{
		Statuses:     []string{"SETTLED"},
		Limit:        5,
		CreatedAfter: createdAfter,
	}

	resps, err := invoice.GetAll(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("invoices: %+v\n", resps)
}

func main(){
	ExampleCreate()
	// ExampleGet()
	// ExampleGetListTransaction()
}