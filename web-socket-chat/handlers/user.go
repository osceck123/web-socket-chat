package user

import (
    "chat-app/auth"
    "context"
    "encoding/json"
    "net/http"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    _, err := userCollection.InsertOne(ctx, user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    var foundUser User
    err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser)
    if err != nil || foundUser.Password != user.Password {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }
    token, _ := auth.GenerateJWT(foundUser.Username)
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
