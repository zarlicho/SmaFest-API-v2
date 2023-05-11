package main

import (
	"fmt"
	"test-golang/config/mongdb"
	// "test-golang/models"
	// "context"
	// "log"
	"os"
    "go.mongodb.org/mongo-driver/bson"
	"github.com/joho/godotenv"
	"test-golang/seeders"
	"strings"
)
// var datasis = []models.Product{}
type data struct{
	OrderID	string
	Name string
	Email string
}


func extract(filds,id string)string{
	datas,_ := mongdb.GetMongoData(bson.M{"orderid": id},"preTicket",filds)
	output := fmt.Sprint(datas)
	value := output[5 : len(output)-1]
	splitStr := strings.Split(value, filds+":")
	values := splitStr[1]
	return values

}

func init(){
	// mongdb.ConnectToDB()
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	fmt.Println(os.Getenv("DB_URI"))
	seeders.Test()
}

func main(){

	// mongdb.InsertData(data{Name: "reyhan"}) //inserting data to DB
	// filter := bson.M{"orderid": "S9V75WVSV5"}
	// fields := `{"name": 1, "email": 1}`
	// results:= mongdb.FindData(filter, "preTicket", fields)
	// fmt.Println("Results:", results)
	// fmt.Println(extract("name","S9V75WVSV5"))
	fmt.Println("test-database")
	// update data:
	// dataFilter := bson.M{"name": "fatih"}
	// mongdb.UpdateOrderID("reyhan","ZKCULFPN5I","preTicket")
	// updateData := bson.M{"$set": bson.M{"name": "zeta"}}
	// mongdb.UpdateData(dataFilter,updateData,"ticket")

	// mongdb.DeleteData(data{Name:"reyhan"},"ticket") //deleting data
	// mongdb.CloseDBConnection() //close mongodb connection
}