package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/ptodev/book-exercises/Building_Distributed_Applications_in_Gin/ch2/recipes-api/handlers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Recipe struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name"`
	Tags         []string           `json:"tags"`
	Ingredients  []string           `json:"ingredients"`
	Instructions []string           `json:"instructions"`
	PublishedAt  time.Time          `json:"publishedAt"`
}

// var recipes []Recipe
var recipesHandler *handlers.RecipesHandler

func init() {
	// recipes := make([]models.Recipe, 0)
	// file, err := os.ReadFile("recipes.json")
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("File contents: ", string(file))
	//
	// err = json.Unmarshal([]byte(file), &recipes)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("Parsed recipes: ", len(recipes))
	//
	// var listOfRecipes []interface{}
	// for _, recipe := range recipes {
	// 	listOfRecipes = append(listOfRecipes, recipe)
	// }
	// log.Println("Converted recipes: ", len(listOfRecipes))

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")

	// insertManyResult, err := collection.InsertMany(ctx, listOfRecipes)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Inserted recipes: ", len(insertManyResult.InsertedIDs))

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       9,
	})
	status := redisClient.Ping()
	fmt.Println(status)

	recipesHandler = handlers.NewRecipesHandler(ctx, collection, redisClient)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipesHandler)
	router.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	router.GET("/recipes/search", recipesHandler.SearchRecipesHandler)
	router.Run()
}
