package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

type User struct {
	ID 		 int 	`json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Channel struct {
	ID 		 int 	`json:"id"`
	Name string `json:"name"`
}

type Message struct {
	ID 		  int 	 `json:"id"`
	ChannelID int 	 `json:"channel_id"`
	UserID	  int 	 `json:"user_id"`
	Username  string `json:"user_name"`
	Text  	  string `json:"text"`
}

func main()  {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Working directory:", wd)

	//Open the SQLite database file

	db, err := sql.Open("sqlite", wd+"/database.db")

	defer func (db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	r := gin.Default()

	if err != nil {
		log.Fatal(err)
	}

	r.POST("/users", func(ctx *gin.Context) { createUser(ctx, db) })
	r.POST("/channels", func(ctx *gin.Context) { createChannel(ctx, db) })
	r.POST("/messages", func(ctx *gin.Context) { createMessage(ctx, db) })

	r.GET("/channels", func(ctx *gin.Context) { listChannels(ctx, db) })
	r.GET("/message", func(ctx *gin.Context) { listMessages(ctx, db) })

	r.POST("/login", func(ctx *gin.Context) { login(ctx, db) })

	err = r.Run(":8080")

	if err != nil {
		log.Fatal(err)
	}
}