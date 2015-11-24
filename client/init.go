package client
import (
	"fmt"
	"net"
	"bufio"
)

type Client struct {
	connection net.Conn
}

func Init(addr string) *Client {
	fmt.Println("Connect to", addr)


	c, err := net.Dial("tcp", addr)
	result := &Client{connection: c}
	if err != nil {
		panic(err)
	}

	return result
}

func (c *Client) Listen() {
	scanner := bufio.NewReader(c.connection)

	for {
		str, err := scanner.ReadString('\n')
		if len(str) > 0 {
			fmt.Print(str)
		}
		if err != nil {
			break
		}
	}
}


func (c *Client) Send(message string) {
	c.connection.Write([]byte(message))
}
