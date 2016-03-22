package main

import (
	"github.com/koding/kite"
	"fmt"
	"os"
	"sync"
	"log"
	"github.com/bitly/go-nsq"
)

func main() {
	k := kite.New("robber", "1.0.0")
	k.Config.Port = 6002
	k.Config.DisableAuthentication = true
	k.HandleFunc("hello", hello)
	
	fmt.Println("Robber staring kite")
	k.Run()

	
	fmt.Println("Robber configuring NSQ")
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("tick", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v", message)
		wg.Done()
		return nil
	}))

	err := q.ConnectToNSQD(os.Getenv("MESSAGING_SERVICE_HOST" + ":4150"))
	if err != nil {
		log.Panic("Could not connect")
	}
	fmt.Println("Robber starting NSQ")
	wg.Wait()


}

func hello(r *kite.Request) (interface{}, error) {

	fmt.Println("Robber got hello")

	// You can return anything as result, as long as it is JSON marshalable.
	return nil, nil
}
