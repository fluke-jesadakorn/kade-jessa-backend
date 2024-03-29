package main

import (
	"log"
	"os"

	"kade-jessa/cloudbucket"
	firebaseAdmin "kade-jessa/firebaseInit"
	"kade-jessa/mongoMethod"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000", "https://www.kadejessa.com", "https://kadejessa.com"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:  []string{"Access-Control-Allow-Headers"},
		AllowWildcard: true,
		// AllowAllOrigins:  true,
		AllowCredentials: true,
	}))

	//Public
	router.GET("/", mongoMethod.GetProducts)
	router.GET("/products/name/:title", mongoMethod.GetProductByName)
	router.GET("/products/id/:id", mongoMethod.GetProductByID)
	router.POST("/admin/login", firebaseAdmin.Login)

	//Admin Protect
	adminRoute := router.Group("/admin", firebaseAdmin.VerifyIDToken)
	{
		adminRoute.GET("/post", mongoMethod.GetProducts)
		adminRoute.POST("/upload", cloudbucket.UploadToBucket)
	}

	//User Protect
	userRoute := router.Group("/user")
	{
		userRoute.POST("/login", firebaseAdmin.Login)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}

}
