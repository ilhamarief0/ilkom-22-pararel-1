package handlers

import (
	"context"
	"database/sql"
	"log"

	pb "user-service/proto"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceServer struct {
	DB *sql.DB
	pb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	var user pb.User
	query := `
		SELECT u.id, u.username, u.email, u.password, r.name as role_name
		FROM users u
		LEFT JOIN roles r ON u.role_id = r.id
		WHERE u.id = ?
	`
	row := s.DB.QueryRow(query, req.Id)

	var roleName string
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &roleName)
	if err != nil {
		log.Printf("Error scanning GetUser row: %v", err)
		return nil, err
	}

	user.Role = roleName
	return &pb.UserResponse{User: &user}, nil
}

func (s *UserServiceServer) GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameRequest) (*pb.GetUserByUsernameResponse, error) {
	var user pb.User
	query := `
		SELECT u.id, u.username, u.email, u.password, r.name as role_name
		FROM users u
		LEFT JOIN roles r ON u.role_id = r.id
		WHERE u.username = ?
	`
	row := s.DB.QueryRow(query, req.Username)

	var roleName string
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &roleName)
	if err != nil {
		log.Printf("Error scanning GetUserByUsername row: %v", err)
		return nil, err
	}

	user.Role = roleName
	return &pb.GetUserByUsernameResponse{User: &user}, nil
}

func (s *UserServiceServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	query := `
		SELECT u.id, u.username, u.email, u.password, r.name as role_name
		FROM users u
		LEFT JOIN roles r ON u.role_id = r.id
	`
	rows, err := s.DB.Query(query)
	if err != nil {
		log.Printf("Failed to execute ListUsers query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []*pb.User
	for rows.Next() {
		var user pb.User
		var roleName string
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &roleName)
		if err != nil {
			log.Printf("Failed to scan ListUsers row: %v", err)
			return nil, err
		}
		user.Role = roleName
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error encountered during ListUsers row iteration: %v", err)
		return nil, err
	}

	log.Printf("Successfully listed %d users", len(users))
	return &pb.ListUsersResponse{Users: users}, nil
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return &pb.CreateUserResponse{Success: false, Message: "Failed to hash password"}, err
	}

	stmt, err := s.DB.Prepare("INSERT INTO users (username, email, password, role_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Printf("Failed to prepare CreateUser statement: %v", err)
		return &pb.CreateUserResponse{Success: false, Message: "Failed to prepare query"}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Username, req.Email, hashedPassword, req.RoleId)
	if err != nil {
		log.Printf("Failed to execute CreateUser query: %v", err)
		return &pb.CreateUserResponse{Success: false, Message: "Failed to create user"}, err
	}

	return &pb.CreateUserResponse{Success: true, Message: "User created successfully"}, nil
}
