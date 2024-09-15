package main

import (
    "github.com/dgrijalva/jwt-go"
    "time"
    "net/http"
    "log"
    "encoding/json"
)

var jwtKey = []byte("my_secret_key")

type Credentials struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

func generateJWT(email string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Email: email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    _, err = db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", creds.Email, creds.Password)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    var storedPassword string
    err = db.QueryRow("SELECT password FROM users WHERE email=$1", creds.Email).Scan(&storedPassword)
    if err != nil || storedPassword != creds.Password {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    tokenString, err := generateJWT(creds.Email)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name:    "token",
        Value:   tokenString,
        Expires: time.Now().Add(24 * time.Hour),
    })
}
