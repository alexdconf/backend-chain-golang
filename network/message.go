package network

type Message struct {
    Type    string      `json:"type"`
    Payload interface{} `json:"payload"`
}
