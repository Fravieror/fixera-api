package service

import (
	"encoding/json"
	"net/http"

	"fixera-api/domain"

	"github.com/gorilla/mux"
)

var users []domain.User

// Constructor
type UserService struct{}

// TODO
func NewUserService() *UserService {
	return &UserService{}
}

// Get all users
func (u UserService) GetUsers(w http.ResponseWriter, r *http.Request) {

	users = append(users, domain.User{Name: "Friend_1", Phone: "98xxx-xxxxx", Email: "person1@mail.com"})
	users = append(users, domain.User{Name: "Friend_2", Phone: "96xxx-xxxxx", Email: "person2@mail.com"})
	users = append(users, domain.User{Name: "Friend_3", Phone: "97xxx-xxxxx", Email: "person3@mail.com"})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Get single user
func (u UserService) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Looping through contacts and find one with the id from the params
	for _, item := range users {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&domain.User{})
}

// Add new user
func (u UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user domain.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// Delete user
func (u UserService) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range users {
		if item.Name == params["name"] {
			users = append(users[:idx], users[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

// Update users
func (u UserService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range users {
		if item.Name == params["name"] {
			users = append(users[:idx], users[idx+1:]...)
			var user domain.User
			_ = json.NewDecoder(r.Body).Decode(user)
			user.Name = params["name"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}
