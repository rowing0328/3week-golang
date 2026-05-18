package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type userRequest struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type Store struct {
	users  []User
	nextID int
}

func NewStore() *Store {
	return &Store{nextID: 1}
}

func (s *Store) List(minAge *int) []User {
	result := make([]User, 0, len(s.users))
	for _, user := range s.users {
		if minAge != nil && user.Age < *minAge {
			continue
		}
		result = append(result, user)
	}
	return result
}

func (s *Store) Get(id int) (User, bool) {
	for _, user := range s.users {
		if user.ID == id {
			return user, true
		}
	}
	return User{}, false
}

func (s *Store) Create(req userRequest) User {
	user := User{
		ID:    s.nextID,
		Name:  req.Name,
		Age:   req.Age,
		Email: req.Email,
	}
	s.nextID++
	s.users = append(s.users, user)
	return user
}

func (s *Store) Update(id int, req userRequest) (User, bool) {
	for i := range s.users {
		if s.users[i].ID == id {
			s.users[i].Name = req.Name
			s.users[i].Age = req.Age
			s.users[i].Email = req.Email
			return s.users[i], true
		}
	}
	return User{}, false
}

func (s *Store) Delete(id int) (User, bool) {
	for i := range s.users {
		if s.users[i].ID == id {
			deleted := s.users[i]
			s.users = append(s.users[:i], s.users[i+1:]...)
			return deleted, true
		}
	}
	return User{}, false
}

func newRouter(store *Store) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", usersHandler(store))
	mux.HandleFunc("/users/", userByIDHandler(store))
	return mux
}

func main() {
	store := NewStore()

	fmt.Println("server started: http://localhost:8080")
	if err := http.ListenAndServe(":8080", newRouter(store)); err != nil {
		fmt.Println("server error:", err)
	}
}

func usersHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUsers(w, r, store)
		case http.MethodPost:
			createUser(w, r, store)
		default:
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	}
}

func userByIDHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, ok := idFromPath(r.URL.Path)
		if !ok {
			writeError(w, http.StatusBadRequest, "invalid user id")
			return
		}

		switch r.Method {
		case http.MethodGet:
			getUser(w, store, id)
		case http.MethodPut:
			updateUser(w, r, store, id)
		case http.MethodDelete:
			deleteUser(w, store, id)
		default:
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	}
}

func getUsers(w http.ResponseWriter, r *http.Request, store *Store) {
	minAge, ok := minAgeFromQuery(w, r)
	if !ok {
		return
	}

	writeJSON(w, http.StatusOK, store.List(minAge))
}

func getUser(w http.ResponseWriter, store *Store, id int) {
	user, ok := store.Get(id)
	if ok {
		writeJSON(w, http.StatusOK, user)
		return
	}

	writeError(w, http.StatusNotFound, "user not found")
}

func createUser(w http.ResponseWriter, r *http.Request, store *Store) {
	req, ok := decodeUserRequest(w, r)
	if !ok {
		return
	}

	writeJSON(w, http.StatusCreated, store.Create(req))
}

func updateUser(w http.ResponseWriter, r *http.Request, store *Store, id int) {
	req, ok := decodeUserRequest(w, r)
	if !ok {
		return
	}

	user, ok := store.Update(id, req)
	if ok {
		writeJSON(w, http.StatusOK, user)
		return
	}

	writeError(w, http.StatusNotFound, "user not found")
}

func deleteUser(w http.ResponseWriter, store *Store, id int) {
	user, ok := store.Delete(id)
	if ok {
		writeJSON(w, http.StatusOK, user)
		return
	}

	writeError(w, http.StatusNotFound, "user not found")
}

func decodeUserRequest(w http.ResponseWriter, r *http.Request) (userRequest, bool) {
	defer r.Body.Close()

	var req userRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json body")
		return userRequest{}, false
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)

	if req.Name == "" {
		writeError(w, http.StatusBadRequest, "name is required")
		return userRequest{}, false
	}

	if req.Age < 0 {
		writeError(w, http.StatusBadRequest, "age must be greater than or equal to 0")
		return userRequest{}, false
	}

	if req.Email == "" {
		writeError(w, http.StatusBadRequest, "email is required")
		return userRequest{}, false
	}

	return req, true
}

func minAgeFromQuery(w http.ResponseWriter, r *http.Request) (*int, bool) {
	minAgeText := strings.TrimSpace(r.URL.Query().Get("minAge"))
	if minAgeText == "" {
		return nil, true
	}

	minAge, err := strconv.Atoi(minAgeText)
	if err != nil || minAge < 0 {
		writeError(w, http.StatusBadRequest, "minAge must be greater than or equal to 0")
		return nil, false
	}

	return &minAge, true
}

func idFromPath(path string) (int, bool) {
	idText := strings.TrimPrefix(path, "/users/")
	if idText == "" || strings.Contains(idText, "/") {
		return 0, false
	}

	id, err := strconv.Atoi(idText)
	if err != nil || id <= 0 {
		return 0, false
	}

	return id, true
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
