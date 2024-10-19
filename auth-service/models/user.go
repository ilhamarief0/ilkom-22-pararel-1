package models

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
	Password string
}

func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	user := &User{}
	row := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(db *sql.DB, userID int) (*User, error) {
	user := &User{}
	row := db.QueryRow("SELECT id, username FROM users WHERE id = ?", userID)
	err := row.Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Fungsi untuk membuat user baru
func CreateUser(db *sql.DB, username, password string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	// Mengecek apakah username sudah ada
	var existingUser User
	err = db.QueryRow("SELECT id, username FROM users WHERE username = ?", username).Scan(&existingUser.ID, &existingUser.Username)
	if err != sql.ErrNoRows {
		return errors.New("username already exists")
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

// Fungsi untuk hashing password saat pendaftaran (jika Anda ingin menambahkan endpoint pendaftaran)
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Fungsi untuk verifikasi password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
