package auth

import (
    "context"
    "fmt"
    "net/http"
    "strings"
    "github.com/dgrijalva/jwt-go"
)

func JwtAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
        if tokenString == "" {
            http.Error(w, "Authorization header required", http.StatusUnauthorized)
            return
        }
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })
        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        ctx := context.WithValue(r.Context(), "userId", claims.UserID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
