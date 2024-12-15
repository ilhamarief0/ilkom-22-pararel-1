package migrations

import (
	"database/sql"
	"log"
)

func Migrate(db *sql.DB) {
	// Membuat tabel product
	createProductsTable := `
	CREATE TABLE IF NOT EXISTS product (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		content VARCHAR(255) NOT NULL,
		image VARCHAR(255) NOT NULL,
		price INT NOT NULL,
		stock INT NOT NULL,
		user_id INT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	// Eksekusi query pembuatan tabel
	if _, err := db.Exec(createProductsTable); err != nil {
		log.Fatalf("Failed to create product table: %v", err)
	}
	log.Println("Product table migrated successfully.")
}
