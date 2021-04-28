package main

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel
var fireStoreClient *firestore.Client
var ctx = context.Background()

// configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Message define our message object
type Message struct {
	Type string `json:"type"`
	Username string `json:"username"`
	Message  string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type User struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Role string `json:"role"`
	CreatedAt string `json:"createdAt"`
}

func main() {

	// setup firestore from firebase app
	sa := option.WithCredentialsFile("./credentials/wizardofoz-b2c61-firebase-adminsdk-hi62x-7bc9782fc7.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	fireStoreClient, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer fireStoreClient.Close()
	fmt.Println("firebase app is initialized.")

	// create a simple file server
	fs := http.FileServer(http.Dir("./front-dist"))
	http.Handle("/", fs)

	// configure websocket route
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	// start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func storeDataToFirebase(message Message) {
	documentName := fmt.Sprintf("content - %d", time.Now().Unix())
	result, err := fireStoreClient.Collection("audio-context").Doc(documentName).Set(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// make sure we close the connection when the function returns
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {

		}
	}(ws)

	// register our new client
	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object

		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		storeDataToFirebase(msg)
		log.Printf("Message: %v", msg)
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
