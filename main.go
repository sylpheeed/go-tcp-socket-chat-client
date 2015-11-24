package main
import "github.com/sylpheeed/go-tcp-socket-chat-client/client"

func main() {
	connection := client.Init("localhost:9999")
	connection.Listen()
}
