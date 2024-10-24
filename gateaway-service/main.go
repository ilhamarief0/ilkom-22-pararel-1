package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("pass1234") // Secret key for signing JWTs

// Function for proxying requests
func proxyHandler(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow your frontend address
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			return // End the preflight request here
		}

		url, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(url)
		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = url.Host
		proxy.ServeHTTP(w, r)
	}
}

// Function to validate JWTs
// Function to validate JWTs
func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
}

// JWT Middleware for protected routes
func jwtMiddleware(next http.Handler, unprotectedRoutes []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins (or specify your frontend origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			return // End preflight request
		}

		// Bypass JWT Validation for specific unprotected routes
		for _, route := range unprotectedRoutes {
			if strings.HasPrefix(r.URL.Path, route) {
				next.ServeHTTP(w, r)
				return
			}
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer ", "", 1))
		token, err := validateJWT(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	serviceMap := map[string]string{
		"/api/product":       "http://localhost:3010",
		"/api/gamberproduct": "http://localhost:3010",
		"/api/product:id":    "http://localhost:3010/:id",
		"/api/auth/login":    "http://localhost:3012",
	}

	unprotectedRoutes := []string{"/api/auth/login"}

	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path) // Log incoming requests

		for path, target := range serviceMap {
			if strings.HasPrefix(r.URL.Path, path) {
				log.Printf("Proxying to: %s", target) // Log the target for proxying
				proxyHandler(target).ServeHTTP(w, r)
				return
			}
		}
		http.NotFound(w, r) // Send 404 if no routes match
	})

	log.Println("Starting server at :3000")
	if err := http.ListenAndServe(":3000", jwtMiddleware(handler, unprotectedRoutes)); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
