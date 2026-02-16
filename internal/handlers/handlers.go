package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "strings"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/robinlieb/devops-lecture-project-2026/pkg/products"
)

var secretKey = []byte("secret-key")

func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    username := r.FormValue("username")
    password := r.FormValue("password")
    if username == "user" && password == "pass" {
        token, err := createToken(username)
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

func AuthLogoutHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"message": "Logout successful"}`))
}

func ProductListHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    response, err := json.Marshal(products.Products)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"error": "Internal Server Error"}`))
        return
    }
    w.Write(response)
}

func ProductDetailHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    idStr := r.PathValue("id")

    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, `{"error":"Product ID has wrong format"}`, http.StatusBadRequest)
        return
    }

    product := products.FindByID(id)
    if product == nil {
        http.Error(w, `{"error":"Product not found"}`, http.StatusNotFound)
        return
    }

    resp, err := json.Marshal(product)
    if err != nil {
        http.Error(w, `{"error":"Internal Server Error"}`, http.StatusInternalServerError)
        return
    }

    w.Write(resp)
}

func CheckoutPlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte(`{"error":"Missing Authorization header"}`))
        return
    }

    const bearerPrefix = "Bearer "
    if !strings.HasPrefix(authHeader, bearerPrefix) {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte(`{"error":"Authorization header must use Bearer scheme"}`))
        return
    }

    tokenString := strings.TrimPrefix(authHeader, bearerPrefix)

    if !verifyToken(tokenString) {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte(`{"error":"Invalid token"}`))
        return
    }

    w.Write([]byte(`{"message":"Order placed successfully"}`))
}

func createToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256,
        jwt.MapClaims{
            "username": username,
            "exp":      time.Now().Add(time.Hour * 24).Unix(),
        })
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

func verifyToken(tokenString string) bool {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method")
        }
        return secretKey, nil
    })

    return err == nil && token.Valid
}
