package handlers

import (
	"WixChallengeBE2023/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {

}

// getTickets responds with the list of all tickets as JSON.

func GetAllTicketsFromDB(w http.ResponseWriter, r *http.Request) {

}

func GetTickets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.GetAllTicketsFromDB())
}

func GetTicketsByTitle(c *gin.Context) {
	title := c.Param("title")

	var ticket = db.GetTicketsByTitleFromDB(title)
	if ticket == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
	} else {
		c.IndentedJSON(http.StatusOK, ticket)
	}
}

func GetTicketsByCreationTime(c *gin.Context) {
	creationTime := c.Param("creationTime")

	time, _ := strconv.Atoi(creationTime)

	var ticket = db.GetTicketsByCreationTimeFromDB(time)
	if ticket == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
	} else {
		c.IndentedJSON(http.StatusOK, creationTime)
	}
}
