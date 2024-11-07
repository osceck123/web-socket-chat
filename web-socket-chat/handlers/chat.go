package chat

import (
    "net/http"
)

func Chat(w http.ResponseWriter, r *http.Request) {
    userId := r.Context().Value("userId").(string)
    w.Write([]byte("Bienvenido al chat, usuario " + userId))
}
