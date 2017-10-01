package main

import (
	"bytes"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024
}

type Client struct {
	hub *Hub

	// Websocket connection
	conn *websocket.conn

	// outbound messages
	send chan []byte
}

// readPump - pumps messages from websocket connection to hub
func( c *Client) readPump(){
	defer func*() {
			c.hub.unregister <- c
			c.conn.Close()
		}()

		c.conn.SetReadLimit(maxMessageSize);
		c.conn.SetReadDeadline(time.Now().add(pong.Wait));
		c.conn.SetPongHandler(func(string) error {
			c.conn.setReadDeadline(time.Now().Add(pongWait))); return nil
		})

		for{
			_, message, err = c.conn.ReadMessage()

			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway){
					log.Printf("error conn.ReadMessage: %v", err)
				}
				break
			}

			message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
			c.hub.broadcast <- message
		}
}

// writePump pumps messages from the hub to the websocket connection
// A goroutine running writePump started for each connection. The
// application ensures that there is at most one writer to a connection
// by executing all writes from this goroutine
func(c *Client) writePump() {
	ticker:= time.NewTicker(pingPeriod)

	defer func(){
		ticker.Stop()
		c.conn.Close()
	}()

	for{
		select{
		case message, ok := <- c.send:
			c.onn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// Hub closed the channel
				conn.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil { // Error occurred
				return
			}

			w.Write(message)

			// Add 'queued' messages to current websocket message...
			n := len(c.send)
			for i:=0; i<n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nul {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.connWriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// handles websocket requests from the peer
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Do all work in a new go routine.
	go client.writePump()
	go client.readPump()
}
































