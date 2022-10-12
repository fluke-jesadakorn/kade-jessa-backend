package firebaseAdmin

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"

	// "firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

func FirebaseInit() *firebase.App {
	ctx := context.Background()
	opt := option.WithCredentialsFile("firebase-admin-keys.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
	}
	return app
}

func VerifyIDToken(c *gin.Context) {
	// ctx := context.Background()

	// type A struct {
	// 	IDToken string `json:"idToken"`
	// }
	// var accessToken A
	// c.BindJSON(&accessToken)

	// app := FirebaseInit()
	// client, err := app.Auth(ctx)

	// if err != nil {
	// 	fmt.Errorf("error getting Auth client: %v\n", err)
	// }

	// cookies, err := client.SessionCookie(ctx, accessToken.IDToken, 10000)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(cookies)

	// subSlice := strings.Split(accessToken.IDToken, "Bearer ")
	// fmt.Println(subSlice)

	// token, err := client.VerifyIDToken(ctx, subSlice[1])
	// if err != nil {
	// 	client.RevokeRefreshTokens(ctx, accessToken.IDToken)
	// 	c.Abort()
	// }
	// fmt.Println(token)
	c.Next()
}

// params := (&auth.UserToCreate{}).
//         Email("user@example.com").
//         EmailVerified(false).
//         PhoneNumber("+15555550100").
//         Password("secretPassword").
//         DisplayName("John Doe").
//         PhotoURL("http://www.example.com/12345678/photo.png").
//         Disabled(false)
// u, err := client.CreateUser(ctx, params)
// if err != nil {
//         log.Fatalf("error creating user: %v\n", err)
// }
// log.Printf("Successfully created user: %v\n", u)
