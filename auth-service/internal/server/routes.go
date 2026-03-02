package server

import (
	"fmt"
	"net/http"

	"github.com/Keotex/devops-lecture-project/shared/pkg/token"
)

func authLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	username := r.FormValue("username")
	password := r.FormValue("password")
	// For simplicity, we'll use a hardcoded username and password
	// This should be replaced with a proper authentication mechanism
	if username == "user" && password == "pass" {
		token, err := token.CreateToken(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error generating the token"}`))
			return
		}
		w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, token)))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "Invalid credentials"}`))
	}
}

func authLogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// In this simple example, we'll just return a success message
	w.Write([]byte(`{"message": "Logout successful"}`))
}
