package notification

import "fmt"

type Gateway struct{}

func NewGateway() *Gateway {
	return &Gateway{}
}

func (g Gateway) send(userId, message string) {
	fmt.Printf("Sending message '%s' to user %s\n", message, userId)
}
