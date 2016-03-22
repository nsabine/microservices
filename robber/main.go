package main

import (
	"github.com/koding/kite"
	"fmt"
	"os"
	"sync"
	"log"
	"github.com/bitly/go-nsq"
	"runtime"
        "github.com/nsabine/microservices/controller/controllerlib"

)

var client *kite.Client
var k 	   *kite.Kite

func main() {
	fmt.Println("Starting Robber")

	k = kite.New("robber", "1.0.0")
	k.Config.Port = 6001
	k.Config.DisableAuthentication = true

        client = k.NewClient("http://" + os.Getenv("CONTROLLER_SERVICE_HOST")  + ":" + os.Getenv("CONTROLLER_SERVICE_PORT") + "/kite")
        connected, err := client.DialForever()
        if err != nil {
                k.Log.Fatal(err.Error())
        }

        // Wait until connected
        <-connected

	runtime.GOMAXPROCS(2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go startKite()
	go startMessaging()
	wg.Wait()
}

func startKite() {
	k.HandleFunc("hello", hello)
	fmt.Println("Robber starting kite")
	k.Run()
}

func startMessaging() {
	fmt.Println("Robber configuring NSQ")

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("tick", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %s", string(message.Body[:]))
		tick()
		return nil
	}))

	fmt.Println("Trying to connect to NSQ: " + os.Getenv("MESSAGING_SERVICE_HOST") + ":4150")

	err := q.ConnectToNSQD(os.Getenv("MESSAGING_SERVICE_HOST") + ":4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	fmt.Println("Robber starting NSQ")
}

func hello(r *kite.Request) (interface{}, error) {

	fmt.Println("Robber got hello")

	// You can return anything as result, as long as it is JSON marshalable.
	return nil, nil
}

func tick() {
	result, err := client.Tell("getState")
        var GameState [][]controllerlib.UpdateRequest
        if err := result.Unmarshal(&GameState); err != nil {
                panic(err)
        }

	x,y := controllerlib.WhereAmI(GameState, controllerlib.UpdateRequest{
		MyName: os.Getenv("HOSTNAME"),
		Type: "Robber",
		XPos: -1,
		YPos: -1,
	})
        fmt.Println("I am here:", x, y)
	if (x<controllerlib.XSize-1) {
		x++
	} else {
		x--
	}
	if (y<controllerlib.YSize-1) {
		y++
	} else {
		y--
	}
	_, err = client.Tell("update", &controllerlib.UpdateRequest{
		MyName: os.Getenv("HOSTNAME"),
		Type: "Robber",
		XPos: x,
		YPos: y,
	})
	if err != nil {
		panic(err)
	}
}
