package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Model User
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

// Credentials for login
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// JWT claims struct
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Global variables
var (
	DB     *gorm.DB
	jwtKey = []byte("pass1234") // Ganti dengan kunci rahasia Anda
)

// Koneksi ke Database
func connectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local" // Sesuaikan dsn Anda
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Println("Database connected")
}

// Fungsi untuk Login
func login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var user User
	result := DB.Where("email = ?", creds.Email).First(&user)
	if result.Error != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Verifikasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Membuat token JWT
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	// Kirim token sebagai response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// Fungsi untuk mengambil semua pengguna (opsional)
func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	result := DB.Find(&users)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	// Kirim data pengguna dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode users to JSON", http.StatusInternalServerError)
	}
}

func main() {
	// Koneksi ke database
	connectDatabase()

	// Secara otomatis migrasi model User
	DB.AutoMigrate(&User{})

	// Rute untuk login dan mengambil pengguna
	http.HandleFunc("/login", login)
	http.HandleFunc("/users", getUsers)

	// Menggunakan middleware CORS
	cors := handlers.AllowedOrigins([]string{"http://localhost:3001"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	// Mulai server dengan middleware CORS
	log.Fatal(http.ListenAndServe(":8082", handlers.CORS(cors, corsMethods, corsHeaders)(http.DefaultServeMux)))
}
