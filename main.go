package main

import (
	"net/http"

	"kade-jessa/cloudbucket"
	firebaseAdmin "kade-jessa/firebaseInit"
	"kade-jessa/mongoMethod"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

func getBooks(c *gin.Context) {

	var books = []book{
		{
			ID:     "1",
			Name:   "Harry Potter",
			Author: "J.K. Rowling",
			Price:  15.9,
		},
		{
			ID:     "2",
			Name:   "One Piece",
			Author: "Oda Eiichirō",
			Price:  2.99,
		},
		{
			ID:     "3",
			Name:   "demon slayer",
			Author: "koyoharu gotouge",
			Price:  2.99,
		},
	}
	c.JSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		// AllowOrigins:  []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers"},
		AllowWildcard:    true,
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))

	router.GET("/", mongoMethod.Get)
	router.POST("/admin/login", firebaseAdmin.Login)

	//Admin Protect
	adminRoute := router.Group("/admin", firebaseAdmin.VerifyIDToken)
	{
		adminRoute.GET("/post", mongoMethod.Get)
		adminRoute.POST("/upload", cloudbucket.UploadToBucket)
	}

	//User Protect
	userRoute := router.Group("/user")
	{
		userRoute.POST("/login", firebaseAdmin.Login)
	}

	router.Run(":8080")

}
