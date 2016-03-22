package main

import (
	"github.com/koding/kite"
	"fmt"
	"os"
	"sync"
	"log"
	"github.com/bitly/go-nsq"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Starting Controller")
	runtime.GOMAXPROCS(2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go startKite()
	go startMessaging()
	wg.Wait()
}

func startKite() {
	fmt.Println("Controller starting kite")
	k := kite.New("controller", "1.0.0")
	k.Config.Port = 6001
	k.Config.DisableAuthentication = true
	k.HandleFunc("hello", hello)
	k.Run()
}

func startMessaging() {
	fmt.Println("Controller starting NSQ")
	config := nsq.NewConfig()

  	w, _ := nsq.NewProducer(os.Getenv("MESSAGING_SERVICE_HOST") + ":4150", config)

	for {
		err := w.Publish("tick", []byte("test"))
		if err != nil {
			log.Panic("Could not connect")
		}
		time.Sleep(time.Second * 1)

	}	
	w.Stop()

}

func hello(r *kite.Request) (interface{}, error) {

	fmt.Println("Controller got hello")

	// You can return anything as result, as long as it is JSON marshalable.
	return nil, nil
}
