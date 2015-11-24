package main
import (
	"github.com/sylpheeed/go-tcp-socket-chat-client/client"
	"fmt"
)



func main() {
	var host string = ""
	fmt.Print("Specify the host: ")
	_, err := fmt.Scanln(&host)
	if err != nil {
		panic(err)
	}
	connection := client.Init(host)
	go connection.InitConsole()
	connection.Listen()
}
