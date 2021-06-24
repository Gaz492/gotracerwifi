package main

import (
	"encoding/json"
	"github.com/pterm/pterm"
	"goTracerWiFi"
	"time"
)

func main()  {
	tracer, err := goTracerWiFi.Status("192.168.1.181", "8088", 5 * time.Second)
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
