// app.go

import "github.com/gorilla/websocket"

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r , nil)

	if err != nul {
		log.Println(err)
		return
	}

	// use conn to send and receive messages.
}