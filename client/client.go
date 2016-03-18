package main

import (
	"fmt"
	"github.com/koding/kite"
	"github.com/koding/kite/protocol"
)

func main() {
	k := kite.New("client", "1.0.0")

	// search a kite that has the same username and environment as us, but the
	// kite name should be "square"
	kites, _ := k.GetKites(&protocol.KontrolQuery{
		Username:    k.Config.Username,
		Environment: k.Config.Environment,
		Name:        "square",
	})

	// there might be several kites that matches our query
	client := kites[0]
	client.Dial()

	response, _ := client.Tell("square", 4)
	fmt.Println(response.MustFloat64())
}
