package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	username := os.Getenv("mysqlUsername")
	password := os.Getenv("mysqlPassword")
	host := os.Getenv("mysqlHost")
	port := os.Getenv("mysqlPort")
	schema := os.Getenv("mysqlSchema")
	timeoutSecs, err := strconv.Atoi(os.Getenv("timeout"))

	if err != nil {
		log.Fatalf("Unable to get timeout %v", err)
	}

	timeout := time.Duration(timeoutSecs) * time.Second
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, schema)

	db, err := sql.Open("mysql", connectionStr)
	defer db.Close()

	if err != nil {
		fmt.Printf("Failed to open sql connection: %s\n", err.Error())
		os.Exit(1)
	}

	startTime := time.Now()

	fmt.Print("trying to connect")
	for {
		fmt.Println("trying to connect...")
		if err = db.Ping(); err == nil {
			break
		}

		if time.Now().Sub(startTime) > timeout {
			log.Fatal("Unable to connect to the database!")
		}

		time.Sleep(2 * time.Second)
	}

	fmt.Println("connection established...")
}
