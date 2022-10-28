package mongoMethod

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Form struct {
	ID          string   `bson:"_id"`
	Title       string   `bson:"title"`
	Colors      []string `bson:"colors"`
	Description string   `bson:"description"`
	Hashtags    []string `bson:"hashtags"`
	Images      []string `bson:"images"`
}

func Get(c *gin.Context) {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var db string = os.Getenv("MONGO_DATABASE")
	var collction string = os.Getenv("MONGO_COLLECTION")

	coll := client.Database(db).Collection(collction)

	findOptions := options.Find()
	// findOptions.SetLimit(5)

	cursor, err := coll.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	var results []Form
	for cursor.Next(context.TODO()) {
		var elem Form
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	c.JSON(http.StatusOK, results)
}

func GetProduct(c *gin.Context) {

	id, has := c.Params.Get("id")
	if has != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}

	convertedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatalln(err)
	}

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var db string = os.Getenv("MONGO_DATABASE")
	var collction string = os.Getenv("MONGO_COLLECTION")

	coll := client.Database(db).Collection(collction)

	findOptions := options.FindOne()
	// findOptions.SetLimit(5)

	var result Form

	coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: convertedID}}, findOptions).Decode(&result)

	fmt.Println(result)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, result)
}

func Delete(receiveID string) {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var db string = os.Getenv("MONGO_DATABASE")
	var collction string = os.Getenv("MONGO_COLLECTION")
	coll := client.Database(db).Collection(collction)

	id, _ := primitive.ObjectIDFromHex(receiveID)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "deleted", Value: true}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(result)
	return
}

func Create(c *gin.Context, insert bson.D) *mongo.InsertOneResult {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var db string = os.Getenv("MONGO_DATABASE")
	var collction string = os.Getenv("MONGO_COLLECTION")
	coll := client.Database(db).Collection(collction)

	// doc := bson.D{{Key: "title", Value: "Record of a Shriveled Datum"}, {Key: "text", Value: "No bytes, no problem. Just insert a document, in MongoDB"}}
	result, err := coll.InsertOne(context.TODO(), insert)

	if err != nil {
		fmt.Println("err" + err.Error())
		log.Fatalln(err.Error())
	}

	return result
}

func Update(receiveID string) {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var db string = os.Getenv("MONGO_DATABASE")
	var collction string = os.Getenv("MONGO_COLLECTION")
	coll := client.Database(db).Collection(collction)

	id, _ := primitive.ObjectIDFromHex(receiveID)

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "title", Value: "hi"}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(result)
	return
}
