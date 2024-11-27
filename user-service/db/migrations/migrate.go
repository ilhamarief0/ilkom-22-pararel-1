package migrations

import (
	"database/sql"
	"log"
)

func Migrate(db *sql.DB) {
	createRolesTable := `
	CREATE TABLE IF NOT EXISTS roles (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL
	);`

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		role_id INT NOT NULL,
		FOREIGN KEY (role_id) REFERENCES roles(id)
	);`

	if _, err := db.Exec(createRolesTable); err != nil {
		log.Fatalf("Failed to create roles table: %v", err)
	}
	if _, err := db.Exec(createUsersTable); err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}
}
