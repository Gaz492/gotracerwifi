package main

import (
	"encoding/json"
	"github.com/Gaz492/gotracerwifi"
	"github.com/pterm/pterm"
	"time"
)

func main()  {
	tracer, err := gotracerwifi.Status("192.168.1.127", "8899", 5 * time.Second, "TCP")
	if err != nil {
		pterm.Fatal.Println(err)
	}

	response, err := json.Marshal(tracer)
	if err != nil {
		pterm.Error.Println(err)
		return
	}
	pterm.Info.Println(string(response))
}
