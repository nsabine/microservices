package main

import (
	"fmt"
	"github.com/koding/kite"
)

func main() {
	k := kite.New("client", "1.0.0")
        k.Config.DisableAuthentication = true

	fmt.Println(k.Config)

	client := k.NewClient("http://square.openshiftapps.com:6001/kite")
	connected, err := client.DialForever()
	if err != nil {
		k.Log.Fatal(err.Error())
	}

	// Wait until connected
	<-connected
	
	response, err := client.Tell("square", 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.MustFloat64())
}
