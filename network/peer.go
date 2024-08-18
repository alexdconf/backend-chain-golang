package network

import (
    "github.com/gorilla/websocket"
)

type Peer struct {
    Conn *websocket.Conn
}

func addPeer(conn *websocket.Conn) *Peer {
    peer := &Peer{Conn: conn}
    peers = append(peers, peer)
    return peer
}
