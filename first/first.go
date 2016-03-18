package main

import (
	"net/url"
        "os"
	"github.com/koding/kite"
)

func main() {
	k := kite.New("first", "1.0.0")
	k.Config.Port = 6001
	k.HandleFunc("square", func(r *kite.Request) (interface{}, error) {
		a := r.Args.One().MustFloat64()
		return a * a, nil
	})

	k.Register(&url.URL{Scheme: "http", Host: os.Getenv("KONTROL_SERVICE_HOST")+":6000/kite"})
	k.Run()
}
