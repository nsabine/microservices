package main

import (
	"net/url"
        "os"
	"github.com/koding/kite"
)

func main() {
	k := kite.New("square", "1.0.0")
	k.Config.Port = 6001
        k.Config.KontrolURL=os.Getenv("KITE_KONTROL_URL")
        k.Config.Username=os.Getenv("KITE_USERNAME")
        k.Config.Environment=os.Getenv("KITE_ENVIRONMENT")

	k.HandleFunc("square", func(r *kite.Request) (interface{}, error) {
		a := r.Args.One().MustFloat64()
		return a * a, nil
	})

	k.Register(&url.URL{Scheme: "http", Host: os.Getenv("SQUARE_SERVICE_HOST")+":6001/kite"})
	k.Run()
}
