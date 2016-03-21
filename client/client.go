package main

import (
	"fmt"
	"github.com/koding/kite"
	"os"
	"time"
	"github.com/nsabine/microservices/square/squarelib"
)

func main() {
	k := kite.New("client", "1.0.0")
        k.Config.DisableAuthentication = true

	fmt.Println(k.Config)

	client := k.NewClient("http://" + os.Getenv("SQUARE_SERVICE_HOST")  + ":" + os.Getenv("SQUARE_SERVICE_PORT") + "/kite")
	connected, err := client.DialForever()
	if err != nil {
		k.Log.Fatal(err.Error())
	}

	// Wait until connected
	<-connected

	requestNum := 0
	for {
		response, err := client.Tell("square", &SquareRequest{
			Number: 4,
			Name: "square-client",
		})
		if err != nil {
			panic(err)
		}
		requestNum++
		fmt.Printf("Request: %d, Response: %f\n", requestNum , response.MustFloat64())
		time.Sleep(time.Second * 1)
	}
}
