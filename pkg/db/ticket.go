package db

import (
	"WixChallengeBE2023/pkg/models"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

func CreateTable() {
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal("Failed to read JSON file:", err)
	}

	var tickets []models.Ticket
	err = json.Unmarshal(file, &tickets)
	if err != nil {
		log.Fatal("Failed to parse JSON data:", err)
	}

	db := GetDB()

	db.Exec(`
		CREATE TABLE IF NOT EXISTS tickets (
			id TEXT PRIMARY KEY,
			title TEXT,
			content TEXT,
			userEmail TEXT,
			creationTime INTEGER,
			labels TEXT
		)
	`)

	for _, ticket := range tickets {
		_, err := db.Exec("INSERT INTO tickets (id, title, content, userEmail, creationTime, labels ) VALUES (?, ?, ?, ?, ?, ?)",
			ticket.Id, ticket.Title, ticket.Content, ticket.UserEmail, ticket.CreationTime, strings.Join(ticket.Labels, " "))
		if err != nil {
			log.Fatal("Failed to insert ticket:", err)
		}
	}
}

func GetAllTicketsFromDB() []models.Ticket {
	db := GetDB()

	rows, err := db.Query("SELECT * FROM tickets")
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	checkErr(err)

	return ConvertToTicketArray(rows)
}

func GetTicketsByTitleFromDB(title string) *models.Ticket {
	db := GetDB()
	rows, err := db.Query("SELECT * FROM tickets WHERE title = $1 ", title)
	defer rows.Close()
	checkErr(err)

	var ticket *models.Ticket = nil

	if rows.Next() {
		ticket = ConvertToTicket(rows)
	}
	return ticket
}

func GetTicketsByCreationTimeFromDB(creationTime int) *models.Ticket {
	db := GetDB()
	rows, err := db.Query("SELECT * FROM tickets WHERE creationTime < $1 ", creationTime)
	defer rows.Close()
	checkErr(err)

	var ticket *models.Ticket = nil

	if rows.Next() {
		ticket = ConvertToTicket(rows)
	}
	return ticket
}

/* Convert to album model- array */

func ConvertToTicketArray(rows *sql.Rows) []models.Ticket {
	tickets := make([]models.Ticket, 0)
	for rows.Next() {
		ticket := ConvertToTicket(rows)
		tickets = append(tickets, *ticket)
	}
	return tickets
}

/* Convert to album model */

func ConvertToTicket(row *sql.Rows) *models.Ticket {

	ticket := &models.Ticket{}
	var temp string
	err := row.Scan(&ticket.Id, &ticket.Title, &ticket.Content, &ticket.UserEmail, &ticket.CreationTime, &temp)
	ticket.Labels = strings.Split(temp, " ")
	checkErr(err)

	return ticket
}
