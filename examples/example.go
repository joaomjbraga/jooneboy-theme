package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

type UserService interface {
	FindUserByID(ctx context.Context, id int64) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	ListActiveUsers(ctx context.Context) ([]*User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) FindUserByID(ctx context.Context, id int64) (*User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return user, nil
}

func (s *userService) CreateUser(ctx context.Context, user *User) (*User, error) {
	if err := user.Validate(); err != nil {
		return nil, fmt.Errorf("invalid user: %w", err)
	}
	
	return s.repo.Create(ctx, user)
}

func (s *userService) ListActiveUsers(ctx context.Context) ([]*User, error) {
	users, err := s.repo.FindActive(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	return users, nil
}

func main() {
	userRepo := NewUserRepository()
	userService := NewUserService(userRepo)
	
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			users, err := userService.ListActiveUsers(r.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users)
			
		case http.MethodPost:
			var user User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}
			
			createdUser, err := userService.CreateUser(r.Context(), &user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(createdUser)
			
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	
	// Goroutine for background processing
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				log.Println("Performing background cleanup...")
				// Background cleanup logic
			case <-ctx.Done():
				return
			}
		}
	}()
	
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}