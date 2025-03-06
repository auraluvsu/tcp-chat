package chatroom

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"net"
)

type Chatroom struct {
	ID   []byte
	name string
}

func CreateNewRoom(newName string) Chatroom {
	idBytes, err := randBytes(8)
	if err != nil {
		log.Fatalf("Error getting custom ID: %v", err)
	}
	newRoom := Chatroom{
		ID:   idBytes,
		name: newName,
	}
	return newRoom
}

func HandleConnection(port int16) (net.Conn, error) {
	host := "localhost"
	address := fmt.Sprintf("%v:%v", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to server...")
	}
	fmt.Printf("Connected to server: %v\n", port)
	return conn, nil
}

func randBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	h := sha256.New()
	h.Write(b)
	f := h.Sum(nil)
	return f, nil
}
