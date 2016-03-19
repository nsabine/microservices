package main

import (
	"fmt"
	"github.com/koding/kite"
	"github.com/koding/kite/protocol"
	"os"
)

func main() {
	k := kite.New("client", "1.0.0")
	k.Config.KontrolURL=os.Getenv("KITE_KONTROL_URL")
	k.Config.Username=os.Getenv("KITE_USERNAME")
        k.Config.DisableAuthentication = true
	
	fmt.Println(k.Config)

	// search a kite that has the same username and environment as us, but the
	// kite name should be "square"
	kites, _ := k.GetKites(&protocol.KontrolQuery{
		Username:    k.Config.Username,
		Environment: k.Config.Environment,
		Name:        "square",
	})

	// there might be several kites that matches our query
	for _, fetchedKite := range kites {
		fetchedKite.Dial()

		response, err := fetchedKite.Tell("square", 4)
		if err != nil {
			panic(err)
		}
		fmt.Println(response.MustFloat64())
	}
}
