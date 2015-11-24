package client
import (
	"bufio"
	"os"
)

func (c *Client) InitConsole() {

	in := bufio.NewReader(os.Stdin)
	for {
		message, err := in.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if message != ""{
			c.Send(message)
			message = ""
		}
	}
}
