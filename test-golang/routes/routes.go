package routes

import (
	"test-golang/controllers"
	// "test-golang/middleware"

	// "github.com/gin-contrib/static"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// inisialisasi Gin
	r := gin.Default()
	router := gin.New()

	// Konfigurasi CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH"}
	config.AllowHeaders = []string{"Origin", "authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "X-Timestamp", "X-Source", "X-Signature", "withCredentials"} // Tambahkan "withCredentials" ke dalam AllowHeaders
	config.ExposeHeaders = []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"}
	config.AllowCredentials = true
	config.MaxAge = 86400

	router.Use(cors.New(config))

	v1 := r.Group("/api")
	{
		v1.OPTIONS("/regis", func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "https://1da7-2001-448a-2020-52c1-381e-7a92-6008-674b.ngrok-free.app")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, withCredentials") // Tambahkan "withCredentials" ke dalam Allow-Headers
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")

			c.AbortWithStatus(200)
		})

		v1.POST("/regis", controllers.PostDataRegis)
		v1.POST("/callbacks", controllers.Callbacksxen)
		v1.GET("/printTicket", controllers.CheckOrderID)
		v1.GET("/ping", controllers.GetPing)

		// PRODUCT
		// v1.GET("products/:page", middleware.Auth, controllers.ProductIndex)
		// v1.GET("product/:id", middleware.Auth, controllers.ProductShow)
		// v1.POST("product", middleware.Auth, controllers.ProductCreate)
		// v1.PUT("product/:id", middleware.Auth, controllers.ProductUpdate)
		// v1.DELETE("product/:id", middleware.Auth, controllers.ProductDelete)
	}

	return r
}
