package main

import (
	"WixChallengeBE2023/pkg/db"
	"WixChallengeBE2023/pkg/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db.ConnectDB()
	db.CreateTable()

	router := gin.Default() // init a new router
	router.GET("/tickets", handlers.GetTickets)
	router.GET("/tickets/:title", handlers.GetTicketsByTitle)
	router.GET("/tickets/time/:creationTime", handlers.GetTicketsByCreationTime)

	log.Println("API is running!")
	router.Run("localhost:9091") // port for application
}
