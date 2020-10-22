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
	username := os.Getenv("dbUsername")
	password := os.Getenv("dbPassword")
	host := os.Getenv("dbHost")
	port := os.Getenv("dbPort")
	schema := os.Getenv("dbSchema")
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

	for {
		fmt.Println("trying to connect...(START)")
		if err = db.Ping(); err == nil {
			break
		}

		if time.Now().Sub(startTime) > timeout {
			log.Fatal("trying to connect...(FAILED)")
		}

		time.Sleep(2 * time.Second)
	}

	fmt.Println("trying to connect...(SUCCEEDED)")
}
