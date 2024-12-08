package handlers

import (
	"context"
	"database/sql"
	"log"

	pb "user-service/proto"
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
		log.Println(err)
		return nil, err
	}

	user.Role = roleName // Assign role name
	return &pb.UserResponse{User: &user}, nil
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
