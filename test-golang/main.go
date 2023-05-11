package main

import (
	// "test-golang/config/database"
	// "test-golang/config/mysql"
	"test-golang/routes"
	// "test-golang/seeders"
	"test-golang/config/mongdb"
	"github.com/joho/godotenv"
)

func main() {

	// Load Env
	
	// Init database
	// mysql.DatabaseInit()
	
	// Run Migration
	// database.RunMigration()
	
	// WARNING , SEEDENG 100 DATA KETIKA MENJALANKAN PROGRAM PERTAMA KALI
	// MOHON COMENT JIKA TIDAK DIGUNAKAN
	// seeders.Seed()
	
	// Router
	r := routes.SetupRouter()
	
	// http://localhost:8080/api/products/1?sort=price&order=desc
	
	// var port = os.Getenv("PORT")
	r.Run()
}

func init(){
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	mongdb.ConnectToDB()
}