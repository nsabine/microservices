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
	"github.com/nsabine/microservices/controller/controllerlib"
)

var GameState [][]controllerlib.UpdateRequest

func main() {
	fmt.Println("Starting Controller")
	reset()
	evaluate()
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
	k.HandleFunc("getState", getState)
	k.HandleFunc("update", update)
	k.Run()
}

func startMessaging() {
	fmt.Println("Controller starting NSQ")
	config := nsq.NewConfig()

  	w, _ := nsq.NewProducer(os.Getenv("MESSAGING_SERVICE_HOST") + ":4150", config)

	for {
		tick(w)
		time.Sleep(time.Second * 10)
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

        // Unmarshal method arguments.
        var params controllerlib.UpdateRequest
        if err := r.Args.One().Unmarshal(&params); err != nil {
                return nil, err
        }

        fmt.Printf("Update received from %s: %s'\n", params.Type, params.MyName)

        // Print a log on remote Kite.
        // This message will be printed on client's console.
        r.Client.Go("kite.log", fmt.Sprintf("Message from %s: Update received", r.LocalKite.Kite().Name))

	x,y := controllerlib.WhereAmI(GameState, params)
	GameState[x][y] = controllerlib.UpdateRequest{"Empty","Empty",x,y}
	GameState[params.XPos][params.YPos] = params

        return nil, nil
}

func getState(r *kite.Request) (interface{}, error) {
        fmt.Println("Controller received state request")
        return GameState, nil
}

func reset() {
	fmt.Println("Resetting Game State")
	// Allocate the top-level slice.
	GameState = make([][]controllerlib.UpdateRequest, controllerlib.YSize) // One row per unit of y.
	// Loop over the rows, allocating the slice for each row.
	for i := range GameState {
		GameState[i] = make([]controllerlib.UpdateRequest, controllerlib.XSize)
		for j := range GameState[i] {
			GameState[i][j] = controllerlib.UpdateRequest{"Empty","Empty",i,j}
		}
	}
}

func evaluate() {
	fmt.Println("Evaluting Game State")
	for i := 0; i<controllerlib.XSize+2; i++ {
		fmt.Print("-")
	}
	fmt.Println()
	for i := range GameState {
		fmt.Print("|")
		for j := range GameState[i] {
			fmt.Print(controllerlib.GetGridCode(GameState[i][j]) + " ")
		}
		fmt.Println("|")
	}
	for i := 0; i<controllerlib.XSize+2; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

