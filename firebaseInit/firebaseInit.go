package firebaseAdmin

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"

	"google.golang.org/api/option"
)

type A struct {
	IDToken string `json:"idToken"`
}

func FirebaseInit() *firebase.App {
	ctx := context.Background()
	opt := option.WithCredentialsFile("firebase-admin-keys.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
	return app
}

func Login(c *gin.Context) {
	ctx := context.Background()

	var accessToken A
	c.BindJSON(&accessToken)
	app := FirebaseInit()
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	expiresIn := time.Hour * 24 * 5
	cookie, err := client.SessionCookie(ctx, accessToken.IDToken, expiresIn)
	c.SetCookie("auth", cookie, int(expiresIn.Seconds()), "/", "http://localhost:3000", true, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Set Cookies Success",
	})
}

func VerifyIDToken(c *gin.Context) {
	ctx := context.Background()

	cookie, err := c.Cookie("auth")

	fmt.Println("cookies ==> " + cookie)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Session Expires",
		})
	}

	app := FirebaseInit()
	client, err := app.Auth(ctx)

	token, err := client.VerifySessionCookie(ctx, cookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
	}
	fmt.Println(token)
	c.Next()
}
