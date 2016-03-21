package square

import (
	"github.com/koding/kite"
	"fmt"
)

type Request struct {
	Number int
	Name   string
}

func main() {
	k := kite.New("square", "1.0.0")
	k.Config.Port = 6001
	k.Config.DisableAuthentication = true
	k.HandleFunc("square", square)
	k.Run()
}

func square(r *kite.Request) (interface{}, error) {
	// Unmarshal method arguments.
	var params Request
	if err := r.Args.One().Unmarshal(&params); err != nil {
		return nil, err
	}

	result := params.Number * params.Number

	fmt.Printf("Call received from '%s', sending result '%.0d' back\n", params.Name, result)

	// Print a log on remote Kite.
	// This message will be printed on client's console.
	r.Client.Go("kite.log", fmt.Sprintf("Message from %s: \"You have requested square of %.0d\"", r.LocalKite.Kite().Name, params.Number))

	// You can return anything as result, as long as it is JSON marshalable.
	return result, nil
}
