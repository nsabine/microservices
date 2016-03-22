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

var GameState [][]uint8
var XSize = 25
var YSize = 25

func main() {
	fmt.Println("Starting Controller")
	reset()
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
		tick(w)
		time.Sleep(time.Second * 1)
	}	
	w.Stop()

}

func tick(w *nsq.Producer) {
	evaluate()	
	err := w.Publish("tick", []byte("test"))
	if err != nil {
		log.Panic("Could not connect")
	}
	
}

func hello(r *kite.Request) (interface{}, error) {

	fmt.Println("Controller got hello")

	// You can return anything as result, as long as it is JSON marshalable.
	return nil, nil
}

func update(r *kite.Request) (interface{}, error) {
        fmt.Println("Controller received state update")
        return nil, nil
}

func getState(r *kite.Request) (interface{}, error) {
        fmt.Println("Controller received state request")
        return GameState, nil
}

func reset() {
	// Allocate the top-level slice.
	GameState := make([][]uint8, YSize) // One row per unit of y.
	// Loop over the rows, allocating the slice for each row.
	for i := range GameState {
		GameState[i] = make([]uint8, XSize)
	}
}

func evaluate() {
	for i := range GameState {
		for j := range GameState[i] {
			fmt.Print(GameState[i][j])
		}
		fmt.Println()
	}
}

