// handlers/video.go
package video

import (
    "net/http"
    "github.com/gorilla/websocket"
    "sync"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

var rooms = make(map[string][]*websocket.Conn)
var roomsMutex = sync.Mutex{}

func VideoCall(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
        return
    }

    defer conn.Close()

    roomID := r.URL.Query().Get("room")
    if roomID == "" {
        roomID = "default"
    }

    roomsMutex.Lock()
    rooms[roomID] = append(rooms[roomID], conn)
    currentUsers := rooms[roomID]
    roomsMutex.Unlock()

    defer func() {
        roomsMutex.Lock()
        for i, c := range rooms[roomID] {
            if c == conn {
                rooms[roomID] = append(rooms[roomID][:i], rooms[roomID][i+1:]...)
                break
            }
        }
        roomsMutex.Unlock()
    }()

    for {
        var msg map[string]interface{}
        err := conn.ReadJSON(&msg)
        if err != nil {
            break
        }

        // Enviar mensaje a los demás usuarios en la sala
        roomsMutex.Lock()
        for _, user := range currentUsers {
            if user != conn {
                user.WriteJSON(msg)
            }
        }
        roomsMutex.Unlock()
    }
}
