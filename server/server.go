package main

import (
	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"net/http"
	"../core/store"
	
)

func ApiMiddleware(db *leveldb.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("leveldbconnc", db)
		ctx.Next()
	}
}

type KVPair struct {
	Key   string
	Value string
}

func main() {
	router := gin.Default()
	db, err := leveldb.OpenFile("Node/LevelDB", nil)

	if err != nil {
		log.Fatal("Unable to Create Key Value Store")
	}

	defer db.Close()

	//router.Use(ApiMiddleware(db))
	// Test API
	router.GET("/v1/Ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Ping Successful", "Status": "OK"})
	})

	router.GET("/v1/get/:key", func(ctx *gin.Context) {
		key := ctx.Param("key")
		value, err := db.Get([]byte(key), nil)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"Key": key, "Value": value})
	})

	router.PUT("v1/put", func(ctx *gin.Context) {
		var data KVPair
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		err := db.Put([]byte(data.Key), []byte(data.Value), nil)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"Message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"Message": "Insertion Successful"})
	})
	router.DELETE("v1/delete/:key", func(ctx *gin.Context) {
		key := ctx.Param("key")
		err := db.Delete([]byte(key), nil)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"Message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"Message": "Deletion Successful"})
	})
	router.Run(":7990")
}
