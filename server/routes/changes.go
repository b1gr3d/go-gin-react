package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"server/models"
)

var validate = validator.New()
var orderCollection *mongo.Collection = OpenCollection(Client, "changes")

func AddChange(c *gin.Context) {
	//TODO maybe do a check that Alpha, Prod, etc are spelled correctly? Return error if they aren't
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var change models.Change

	if err := c.BindJSON(&change); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(change)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	change.ID = primitive.NewObjectID()

	result, insertErr := orderCollection.InsertOne(ctx, change)
	if insertErr != nil {
		msg := fmt.Sprintf("change item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result)
}

func GetChanges(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var changes []bson.M
	cursor, err := orderCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = cursor.All(ctx, &changes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()

	fmt.Println(changes)
	c.JSON(http.StatusOK, changes)
}

func GetChangesByEnv(c *gin.Context) {
	env := c.Params.ByName("env")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var orders []bson.M
	cursor, err := orderCollection.Find(ctx, bson.M{"env": env})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = cursor.All(ctx, &orders); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cancel()
	fmt.Println(orders)
	c.JSON(http.StatusOK, orders)
}

func GetChangeById(c *gin.Context) {
	changeID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(changeID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var order bson.M
	if err := orderCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cancel()
	fmt.Println(order)
	c.JSON(http.StatusOK, order)
}

func UpdateEnv(c *gin.Context) {
	changeID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(changeID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	type Env struct {
		Env *string `json:"env"`
	}
	var env Env
	if err := c.BindJSON(&env); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := orderCollection.UpdateOne(ctx, bson.M{"_id": docID},
		bson.D{
			{"$set", bson.D{{"server", env.Env}}},
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)
}

func UpdateChange(c *gin.Context) {
	changeID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(changeID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var change models.Change
	if err := c.BindJSON(&change); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErr := validate.Struct(change)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	result, err := orderCollection.ReplaceOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{
			"user": change.User,
			"env":  change.Env,
			"app":  change.App,
			"desc": change.Desc,
			"date": change.Date,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)
}

func DeleteChange(c *gin.Context) {
	changeID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(changeID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	result, err := orderCollection.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.DeletedCount)
}
