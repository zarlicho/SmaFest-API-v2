package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type PreTicket struct {
	ID       			primitive.ObjectID 	`bson:"_id,omitempty"`
	Name     			string             	`bson:"name"`
	Email    			string       		`bson:"email"`
	PhoneNumber			string				`bson:"phonenumber"`
	Amount				string				`bson:"amount"`
	OrderID  			string				`bson:"orderid"`	
	CreatedAt   		time.Time	
	UpdatedAt   		time.Time
}

type Ticket struct {
	ID       			primitive.ObjectID 	`bson:"_id,omitempty"`
	Name     			string             	`bson:"name"`
	Email    			string       		`bson:"email"`
	PhoneNumber			string				`bson:"phonenumber"`
	TicketID 			string				`bson:"ticketid"`
	OrderID  			string				`bson:"orderid"`
	TransactionID 		string				`bson:"transactionid"`
	Amount		 		string				`bson:"amount"`
	TransactionStatus 	string				`bson:"status"`
	VaNumber 			string				`bson:"vanumber"`	
	QRCODE   			string				`bson:"qrcode"`
	CreatedAt   		time.Time	
	UpdatedAt   		time.Time
}
