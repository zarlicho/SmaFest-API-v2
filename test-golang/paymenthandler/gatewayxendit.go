package paymenthandler

import (
	"log"
	"os"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

func CreateInvoice(name, email, phone, orderid string, amount int64) string {
	xendit.Opt.SecretKey = os.Getenv("XND_DEVELOPMENT_API")
	data := invoice.CreateParams{
		ExternalID:  orderid,
		Amount:      float64(amount),
		PayerEmail:  email,
		Description: "ticket-SmaFest",
		// SuccessRedirectURL: "https://www.youtube.com/watch?v=5Bkub_BYZvU",
		SuccessRedirectURL: "https://af0f-2001-448a-2020-7d04-a18c-382c-22ef-136b.ngrok-free.app/printTicket?orderid=" + orderid,
		Customer: xendit.InvoiceCustomer{
			GivenNames:   name,
			Surname:      name,
			MobileNumber: phone,
			Email:        email,
		},
	}
	resp, err := invoice.Create(&data)
	if err != nil {
		log.Fatal(err)
	}
	return resp.InvoiceURL
}
