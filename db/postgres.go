package db

import (
	"database/sql"
	"fmt" // Added this

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	// Using your "vegeta" db name and password
	// Inside ConnectDB
	connStr := "host=localhost port=5432 user=postgres password=password123 dbname=Vegeta sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func FetchAll() error {
	// Call the function you already wrote!
	database, err := ConnectDB()
	if err != nil {
		return err
	}
	defer database.Close()

	rows, err := database.Query("SELECT id, content, created_at FROM resource_table")
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Printf("%-5s | %-30s | %-20s\n", "ID", "CONTENT", "CREATED_AT")
	fmt.Println("------------------------------------------------------------------")

	for rows.Next() {
		var id int
		var content string
		var createdAt string
		if err := rows.Scan(&id, &content, &createdAt); err != nil {
			return err
		}
		fmt.Printf("%-5d | %-30s | %-20s\n", id, content, createdAt)
	}
	return nil
}
func DeleteRecord(id string) error {
	database, err := ConnectDB()
	if err != nil {
		return err
	}
	defer database.Close()

	// Execute the delete query
	result, err := database.Exec("DELETE FROM resource_table WHERE id = $1", id)
	if err != nil {
		return err
	}

	// Check if any row was actually deleted
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no record found with ID %s", id)
	}

	return nil
}
func UpdateRecord(id string, newData string) error {
	database, err := ConnectDB()
	if err != nil {
		return err
	}
	defer database.Close()

	result, err := database.Exec("UPDATE resource_table SET content = $1 WHERE id = $2", newData, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no record found with ID %s to update", id)
	}

	return nil
}
