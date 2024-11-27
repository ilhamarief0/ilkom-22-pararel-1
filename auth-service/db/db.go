package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	var err error
	// Ubah `username`, `password`, `dbname`, dan `localhost` sesuai dengan konfigurasi MySQL Anda
	DB, err = sql.Open("mysql", "root:111999@tcp(localhost:3306)/ecommerce")
	if err != nil {
		log.Fatal(err)
	}

	// Membuat tabel users jika belum ada
	createUserTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL
    );`
	_, err = DB.Exec(createUserTable)
	if err != nil {
		log.Fatal(err)
	}
}
