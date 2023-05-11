package mongdb

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
    "time"
	"os"
	log "github.com/sirupsen/logrus"
    // "test-golang/models"
    "go.mongodb.org/mongo-driver/bson"
)

var client *mongo.Client
var db *mongo.Database

func ConnectToDB() {
	// Set up MongoDB connection optionss
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI"))
	// Connect to MongoDB
	var err error
	client, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection was successful
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Connected to MongoDB!")

	// Set the default database
	db = client.Database("SmaFest")
}
// Example of how to use the db variable to interact with MongoDB
func InsertData(data interface{}, collections string) {
	collection := db.Collection(collections)
	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Data inserted successfully!")
	
}

func InsertDataTicket(data interface{},filter bson.M, collections string) {
	collection := db.Collection(collections)
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		// log.Fatal(err)
		log.Warn("error or data already exist")
	}
	defer cursor.Close(context.Background())
	var results []interface{}
	err = cursor.All(context.Background(), &results)
	if err != nil {
		log.Fatal(err)
	}
	if results != nil{
		log.Info("data already exist")
		// return "data already exist"
	}else{
		_, err := collection.InsertOne(context.Background(), data)
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Data inserted successfully!")
	}
}

func GetMongoData(filter bson.M, collections string, fields string) ([]bson.M, error) {
    ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
    collection := db.Collection(collections)
    cur, err := collection.Find(ctx, filter, options.Find().SetProjection(bson.M{fields: 1}))
    if err != nil {
        return nil, err
    }
    var result []bson.M
    if err := cur.All(ctx, &result); err != nil {
        return nil, err
    }
    return result, nil
}

// Example of how to use the db variable to interact with MongoDB
func UpdateData(filter interface{}, update interface{}, collections string) {
	collection := db.Collection(collections)
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Data updated successfully!")
}
//update order id
func UpdateOrderID(email, orderID string) error {
    // filter untuk mencari data berdasarkan nama
	collection := db.Collection("preTicket")
    filter := bson.M{"email": email}

    // update untuk mengubah nilai orderID
    update := bson.M{
        "$set": bson.M{
            "orderid": orderID,
        },
    }

    // konfigurasi options untuk operasi update
    opts := options.Update().SetUpsert(false)

    // panggil fungsi UpdateOne untuk melakukan operasi update pada database
    result, err := collection.UpdateOne(context.Background(), filter, update, opts)
    if err != nil {
        return fmt.Errorf("failed to update orderID for %s: %v", email, err)
    }

    // cek apakah ada data yang berhasil diupdate
    if result.ModifiedCount == 0 {
        return fmt.Errorf("no data found for %s", email)
    }

    fmt.Printf("orderID updated for %s\n", email)
    return nil
}

// Example of how to use the db variable to interact with MongoDB
func DeleteData(filter interface{}, collections string) {
	collection := db.Collection(collections)
	_, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Data deleted successfully!")
}

func CloseDBConnection() {
	// Close the MongoDB client
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Disconnected from MongoDB!")
}

    
    