package connection

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// The function `ConnectionRequest` establishes a WebSocket connection, reads and writes messages, and
// broadcasts messages to all connected clients.
func ConnectionRequest(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error Upgrading to Websockets",
		})
	}

	defer conn.Close()

	var clients = make(map[*websocket.Conn]bool) // Conjunto de conexões ativas
	clients[conn] = true

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Error Reading Message",
			})
			break
		}
		err = conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Error Writing Message",
			})
			break
		}
		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}

}
