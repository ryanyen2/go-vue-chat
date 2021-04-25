package main

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"net/http"

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
	// create a simple file server
	fs := http.FileServer(http.Dir("C:\\Users\\ryanm\\GolandProjects\\go-vue-chat\\public\\vue-chat\\dist"))
	http.Handle("/", fs)

	// configure websocket route
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	// start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8010")
	err := http.ListenAndServe(":8010", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	// setup firestore from firebase app
	sa := option.WithCredentialsFile("./credentials/wizardofoz-b2c61-firebase-adminsdk-hi62x-7bc9782fc7.json")
	app, err := firebase.NewApp(ctx, nil, sa)

	fireStoreClient, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {

		}
	}(fireStoreClient)
}

func storeDataToFirebase(ctx context.Context, client *firestore.Client, message Message) {
	result, err := client.Collection("audio-context").Doc("contentHistory").Set(ctx, message)
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

		storeDataToFirebase(ctx, fireStoreClient, msg)
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
