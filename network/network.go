package network

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
    "encoding/json"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

var peers = make([]*Peer, 0)

func handleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer ws.Close()

    peer := &Peer{Conn: ws}
    peers = append(peers, peer)

    for {
        _, msg, err := ws.ReadMessage()
        if err != nil {
            log.Printf("error: %v", err)
            break
        }
        handleMessage(msg)
    }
}

func handleMessage(msg []byte) {
    var message Message
    json.Unmarshal(msg, &message)

    switch message.Type {
    case "new_block":
        // Handle new block message
    case "new_transaction":
        // Handle new transaction message
    case "request_blockchain":
        // Handle blockchain request message
    case "blockchain":
        // Handle received blockchain message
    }
}

func broadcastMessage(msg []byte) {
    for _, peer := range peers {
        err := peer.Conn.WriteMessage(websocket.TextMessage, msg)
        if err != nil {
            log.Printf("error: %v", err)
            peer.Conn.Close()
        }
    }
}
