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
	fmt.Println("Starting Cop")

	k = kite.New("cop", "1.0.0")
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
	fmt.Println("Cop starting kite")
	k.Run()
}

func startMessaging() {
	fmt.Println("Cop configuring NSQ")

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
	fmt.Println("Cop starting NSQ")
}

func hello(r *kite.Request) (interface{}, error) {

	fmt.Println("Cop got hello")

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
		Type: "Cop",
		XPos: -1,
		YPos: -1,
	})

	me := controllerlib.UpdateRequest{
                MyName: os.Getenv("HOSTNAME"),
                Type: "Cop",
                XPos: x,
                YPos: y,
        }

        fmt.Println("I am here:", x, y)
	nearestRobber := controllerlib.WhereNearest(GameState, me, "Robber")
	if nearestRobber.XPos < me.XPos {
		x--
	} else {
		x++
	}
	if nearestRobber.YPos < me.YPos {
		y--
	} else {
		y++
	}

	// make sure we didn't leave the board
	if x<0 {
		x=0
	}
	if y<0 {
		y=0
	}
	if x==controllerlib.XSize-1 {
		x--
	}
	if y==controllerlib.YSize-1 {
		y--
	}
	_, err = client.Tell("update", &controllerlib.UpdateRequest{
		MyName: os.Getenv("HOSTNAME"),
		Type: "Cop",
		XPos: x,
		YPos: y,
	})
	if err != nil {
		panic(err)
	}
}
