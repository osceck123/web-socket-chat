package main

import (
    "web-socket-chat/auth/middleware"
    "web-socket-chat/handlers"
    "context"
    "fmt"
    "log"
    "net/http"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/gorilla/mux"
)

var client *mongo.Client

func main() {
    // Conexión a MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    clientOptions := options.Client().ApplyURI("mongodb://admin:password@localhost:27017")
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)

    // Configuración de rutas
    router := mux.NewRouter()
    router.HandleFunc("/register", handlers.Register).Methods("POST")
    router.HandleFunc("/login", handlers.Login).Methods("POST")
    
    // Rutas protegidas
    secure := router.PathPrefix("/").Subrouter()
    secure.Use(middleware.JwtAuthMiddleware)
    secure.HandleFunc("/chat", handlers.Chat).Methods("GET")
    secure.HandleFunc("/video", handlers.VideoCall).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", router))
}
