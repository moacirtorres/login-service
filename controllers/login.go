package controllers

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/models"
	"github.com/hyperyuri/webapi-with-go/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Login(c *gin.Context) {

	var login models.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	// Verifying if the user sends the name and password
	if login.Name == "" {
		c.AbortWithError(400, err)
	}
	if login.Password == "" {
		c.AbortWithError(400, err)
	}

	// Getting token from authentication service....
	token := services.AuthMiddleware(login, c)

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://moatorres:abacate123@restful-api.l5qph.mongodb.net/mongodb?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Selecting our database...

	collection := client.Database("mongodb").Collection("mongo_db")

	// Preparing to connect to mongo...
	result, err := collection.InsertOne(ctx, login)
	if err != nil {
		log.Fatal("Error:", err)
	}

	c.JSON(200, gin.H{
		"registro": result,
		"mensagem": "Login efetuado com sucesso.",
		"token":    token,
	})

}
