package main

import (
	"github.com/koding/kite"
)

func main() {
	k := kite.New("square", "1.0.0")
	k.Config.Port = 6001
	k.Config.DisableAuthentication = true

	k.HandleFunc("square", func(r *kite.Request) (interface{}, error) {
		a := r.Args.One().MustFloat64()
		return a * a, nil
	})

	k.Run()
}
