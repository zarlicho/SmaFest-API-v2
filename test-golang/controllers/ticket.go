package controllers

import (
	"net/http"
	"test-golang/config/mongdb"
	"test-golang/models"
	"test-golang/paymenthandler"

	"github.com/gin-gonic/gin"

	// "myProject/model"
	// "myProject/initializers"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	// "crypto/rand"
	// "golang.org/x/crypto/bcrypt"
	// "github.com/golang-jwt/jwt/v4"
	// "time"
)

var datasis = []models.PreTicket{}
var ticketData = []models.Ticket{}
var Callbackss = []models.Callbacks{}

func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func PostDataRegis(c *gin.Context) {
	var newTicket models.PreTicket
	if err := c.BindJSON(&newTicket); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf(e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
		}
		return
	}
	datasis = append(datasis, newTicket) // this function for adding data
	orderids := paymenthandler.GenerateOrderID()
	totalamount, err := strconv.ParseInt(newTicket.Amount, 10, 64)
	if err != nil {
	}
	snap := paymenthandler.CreateInvoice(newTicket.Name, newTicket.Email, newTicket.PhoneNumber, orderids, totalamount)
	mongdb.InsertData(newTicket, "preTicket")
	mongdb.UpdateOrderID(newTicket.Email, orderids)
	c.JSON(http.StatusOK, gin.H{
		"newdata": newTicket,
		"snap":    snap,
	})
}

func extract(filds, id string) string {
	datas, _ := mongdb.GetMongoData(bson.M{"orderid": id}, "preTicket", filds)
	output := fmt.Sprint(datas)
	value := output[5 : len(output)-1]
	splitStr := strings.Split(value, filds+":")
	values := splitStr[1]
	return values
}
func getDataFromid(filds, id string) string {
	datas, _ := mongdb.GetMongoData(bson.M{"orderid": id}, "Ticket", filds)
	output := fmt.Sprint(datas)
	value := output[5 : len(output)-1]
	splitStr := strings.Split(value, filds+":")
	values := splitStr[1]
	return values
}

func srchData(querys, filds string) string {
	datas, _ := mongdb.SearchMongoData(querys, "Ticket", filds)
	output := fmt.Sprint(datas)
	value := output[5 : len(output)-1]
	splitStr := strings.Split(value, filds+":")
	values := splitStr[1]
	return values
}

func SearchData(c *gin.Context) {
	query := c.Query("qury")
	qrcode := srchData(query, "qrcode")
	name := srchData(query, "name")
	email := srchData(query, "email")
	ticketid := srchData(query, "ticketid")
	orderId := srchData(query, "orderid")
	phone := srchData(query, "phonenumber")
	c.JSON(http.StatusOK, gin.H{
		"qrcode":       qrcode,
		"name":         name,
		"email":        email,
		"ticket id":    ticketid,
		"order id":     orderId,
		"phone number": phone,
	})
}

func CheckOrderID(c *gin.Context) {
	orderId := c.Query("orderid")
	qrcode := getDataFromid("qrcode", orderId)
	transactionid := getDataFromid("transactionid", orderId)
	name := getDataFromid("name", orderId)
	email := getDataFromid("email", orderId)
	ticketid := getDataFromid("ticketid", orderId)
	c.JSON(http.StatusOK, gin.H{
		"qrcode":         qrcode,
		"transaction id": transactionid,
		"name":           name,
		"email":          email,
		"ticket id":      ticketid,
		"order id":       orderId,
	})
}

func Callbacksxen(c *gin.Context) {
	type tempData struct {
		OrderID           string
		Name              string
		Email             string
		PhoneNumber       string
		TicketID          string
		TransactionID     string
		Amount            string
		TransactionStatus string
		QRCODE            string
	}
	var payload map[string]interface{}
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// fmt.Println(payload["external_id"])
	orderID := payload["external_id"]
	// Check if the payment is successful
	if payload["status"] == "PAID" {
		// Process the payment
		upTempData := tempData{
			OrderID:     orderID.(string),
			Name:        extract("name", orderID.(string)),
			Email:       extract("email", orderID.(string)),
			PhoneNumber: extract("phonenumber", orderID.(string)),
			Amount:      extract("amount", orderID.(string)),
			TicketID:    paymenthandler.GenerateTicketID(),
			QRCODE:      paymenthandler.GenerateRandomQRCode(),
		}
		fmt.Println(upTempData)
		mongdb.InsertDataTicket(upTempData, bson.M{"orderid": orderID.(string)}, "Ticket")
		// You can perform other actions here, such as updating your database, sending email notifications, etc.
	} else {
		fmt.Println("Payment failed!")
		// You can handle the failed payment here
	}

	c.Status(http.StatusOK)
}
