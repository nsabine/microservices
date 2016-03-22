package main

import (
	"github.com/koding/kite"
	"fmt"
	"os"
	"sync"
	"log"
	"github.com/bitly/go-nsq"
	"runtime"
)

func main() {
	fmt.Println("Starting Robber")
	runtime.GOMAXPROCS(2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go startKite()
	go startMessaging()
	wg.Wait()
}

func startKite() {
	k := kite.New("robber", "1.0.0")
	k.Config.Port = 6001
	k.Config.DisableAuthentication = true
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
