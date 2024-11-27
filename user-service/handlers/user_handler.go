package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	pb "user-service/proto"
)

type UserServiceServer struct {
	DB *sql.DB
	pb.UnimplementedUserServiceServer
}

// CreateUser handles the creation of a new user.
func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	stmt, err := s.DB.Prepare("INSERT INTO users (username, email, password, role_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Printf("Failed to prepare statement: %v\n", err)
		return &pb.CreateUserResponse{Success: false, Message: "Failed to prepare query"}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Username, req.Email, req.Password, req.RoleId)
	if err != nil {
		log.Printf("Failed to execute query: %v\n", err)
		return &pb.CreateUserResponse{Success: false, Message: "Failed to create user"}, err
	}

	return &pb.CreateUserResponse{Success: true, Message: "User created successfully"}, nil
}

// GetUser handles fetching a user by ID.
func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	var user pb.User
	err := s.DB.QueryRow("SELECT id, username, email, role_id FROM users WHERE id = ?", req.Id).
		Scan(&user.Id, &user.Username, &user.Email, &user.RoleId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User with ID %d not found\n", req.Id)
			return nil, fmt.Errorf("user not found")
		}
		log.Printf("Failed to query user: %v\n", err)
		return nil, err
	}

	return &pb.UserResponse{User: &user}, nil
}

// UpdateUser handles updating an existing user.
func (s *UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	stmt, err := s.DB.Prepare("UPDATE users SET username = ?, email = ?, password = ?, role_id = ? WHERE id = ?")
	if err != nil {
		log.Printf("Failed to prepare statement: %v\n", err)
		return &pb.UpdateUserResponse{Success: false, Message: "Failed to prepare query"}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(req.Username, req.Email, req.Password, req.RoleId, req.Id)
	if err != nil {
		log.Printf("Failed to execute query: %v\n", err)
		return &pb.UpdateUserResponse{Success: false, Message: "Failed to update user"}, err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		log.Printf("No user found with ID %d to update\n", req.Id)
		return &pb.UpdateUserResponse{Success: false, Message: "User not found"}, nil
	}

	return &pb.UpdateUserResponse{Success: true, Message: "User updated successfully"}, nil
}

// DeleteUser handles deleting a user by ID.
func (s *UserServiceServer) DeleteUser(ctx context.Context, req *pb.UserRequest) (*pb.DeleteUserResponse, error) {
	stmt, err := s.DB.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		log.Printf("Failed to prepare statement: %v\n", err)
		return &pb.DeleteUserResponse{Success: false, Message: "Failed to prepare query"}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(req.Id)
	if err != nil {
		log.Printf("Failed to execute query: %v\n", err)
		return &pb.DeleteUserResponse{Success: false, Message: "Failed to delete user"}, err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		log.Printf("No user found with ID %d to delete\n", req.Id)
		return &pb.DeleteUserResponse{Success: false, Message: "User not found"}, nil
	}

	return &pb.DeleteUserResponse{Success: true, Message: "User deleted successfully"}, nil
}
