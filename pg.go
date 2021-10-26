package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var pgPassword string

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = pgPassword
	dbname   = "go-postgres"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("could not load env file to run the rest of the program.", err)
	}
	pgPassword = os.Getenv("USER_PASSWORD")
}

// insertIntoDB - queries db, uses sql statement to add users.
func insertIntoDB(db *sql.DB) {
	// Sql statements
	sqlStatement := `
      INSERT INTO users (age, email, first_name, last_name)
      VALUES ($1, $2, $3, $4)
      RETURNING id
      `
	// since pg doesnt return the last insert id, we get around this.
	id := 0
	err := db.QueryRow(sqlStatement, 25, "brendi@brendi.co", "brendi", "prendi").Scan(&id)
	if err != nil {
		log.Println("could not exec sqlStatement a few line above this error: ", err)
	}
	fmt.Printf("New record ID is: %d", id)
}

func main() {

	// connection string
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open the database
	// simply validates the arguments given, doesnt actually check if the connection exists.
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	// Ping the db to make sure it is there.
	err = db.Ping()
	if err != nil {
		log.Println("could not ping the db: ", err)
	}
	defer db.Close()

	fmt.Println("connected to db")
}
