package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sqlStatement := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
		err = db.QueryRow(sqlStatement, user.Username, hashedPassword).Scan(&user.ID)
		if err != nil {
			fmt.Println("error db...")
			http.Error(w, err.Error()+"Error insert", http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("class handlers in method LoginHandler...")
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		var dbUser User
		sqlStatement := `SELECT id, password FROM users WHERE username=$1`
		err := db.QueryRow(sqlStatement, user.Username).Scan(&dbUser.ID, &dbUser.Password)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Authentication successful, return the user ID and username
		response := struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		}{
			ID:       dbUser.ID,
			Username: user.Username,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func GetAllUsersHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, username FROM users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Username)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		if err = rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(users)
	}
}
