package cloudbucket

import (
	"encoding/json"
	"fmt"
	"io"
	"kade-jessa/mongoMethod"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

var (
	storageClient *storage.Client
)

var GCSBucket = "kade-jessa"

func uploadByFile(c *gin.Context, file *multipart.FileHeader) string {

	var err error

	ctx := appengine.NewContext(c.Request)
	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))

	if err != nil {
		fmt.Println("Key error")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return err.Error()
	}

	f, err := file.Open()

	if err != nil {
		fmt.Println("File error")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return err.Error()
	}

	defer f.Close()

	dateNow := strconv.FormatInt(int64(time.Now().Nanosecond()), 10)

	sw := storageClient.Bucket(GCSBucket).Object(dateNow).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		fmt.Println("Storage Error")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return err.Error()
	}

	if err := sw.Close(); err != nil {
		fmt.Println("Upload Error" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return err.Error()
	}

	u, err := url.Parse("https://storage.googleapis.com" + "/" + GCSBucket + "/" + sw.Attrs().Name)
	errAcls := storageClient.Bucket(GCSBucket).Object(sw.Attrs().Name).ACL().Set(ctx, storage.AllUsers, storage.RoleReader)

	if errAcls != nil {
		fmt.Println("Acls Error")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errAcls.Error(),
			"Error":   true,
		})
		return err.Error()
	}

	if err != nil {
		fmt.Println("Parse Url Error")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"Error":   true,
		})
		return err.Error()
	}

	return u.EscapedPath()
}

func UploadToBucket(c *gin.Context) {
	type Form struct {
		Title       string   `bson:"title"`
		Colors      []string `bson:"colors"`
		Description string   `bson:"description"`
		Hashtags    []string `bson:"hashtags"`
		Images      []string `bson:"images"`
	}

	form, err := c.MultipartForm()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var assign Form

	data := form.Value["data"]

	json.Unmarshal([]byte(data[0]), &assign)

	var res []string

	if len(form.File) <= 0 {
		return
	}

	for _, files := range form.File {
		for _, file := range files {
			url := uploadByFile(c, file)
			res = append(res, url)
		}
	}

	setCreate := bson.D{
		{Key: "title", Value: assign.Title},
		{Key: "colors", Value: assign.Colors},
		{Key: "descriptions", Value: assign.Description},
		{Key: "hashtags", Value: assign.Hashtags},
		{Key: "images", Value: res},
	}

	ret := mongoMethod.Create(c, setCreate)

	log.Printf("inserted ==> %v", ret)

	c.JSON(http.StatusOK, gin.H{
		"message": "Created Success",
	})
}
