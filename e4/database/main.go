package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	db, err := sql.Open("mysql", "root:alifiandra@tcp(127.0.0.1:3306)/sik")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Execute the database query with the custom context
	rows, err := db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	// Process query results
	fmt.Println("Process response")
}
