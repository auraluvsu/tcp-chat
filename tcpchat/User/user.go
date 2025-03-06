package user

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"net"
	"os"

	"auraluvsu.com/Chat"
)

type User struct {
	name   string
	userid []byte
	room   chatroom.Chatroom
}

type ChatMsg struct {
	message []byte
	user    User
}

/*
This function will be called later to create a new user everytime a new client joins the server
Not sure how exactly Im gonna do this project as im still pretty new to programming concepts such as buffers and the semantics of how data is sent
over a connection like TCP. I understand how the request/response model works but this kind of conenction is new to me, especially because i dont fully understand
all of the imported functions, as well as which ones to use and when...
*/
func CreateUser() (user User) {
	var username string
	fmt.Println("Create a username:")
	fmt.Scan(&username) //This will allow the user to create a custom username that will then be logged into a struct called newUser.
	UserID := CreateUserID(8)
	newUser := User{
		name:   username,
		userid: UserID,
	}
	return newUser
}

func CreateUserID(b int) []byte {
	name, err := randBytes(8)
	if err != nil {
		log.Fatal("Error processing bytes:", err)
	}
	newhash := sha256.New()
	newhash.Write([]byte(name))
	shaname := newhash.Sum(nil)
	return shaname
}

func randBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (u *User) SendMessage(conn net.Conn, message string) {
	messageChannnel := make(chan string)

	go func() {
		for msg := range messageChannnel {
			_, err := conn.Write([]byte(msg + "\n"))
			if err != nil {
				fmt.Println("Error writing to server:", err)
				return
			}
		}
	}()

	fmt.Println("Connected to chatroom. Enter message:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		messageChannnel <- message
	}
	close(messageChannnel)
}
