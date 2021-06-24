# Go Tracer WiFi

A Go package to communicate with an EPSolar Tracer A/AN charge controller equipped with the eBox-Wifi-01

Only tested with EPSolar Tracer3210AN equipped with the eBox-Wifi-01

## Example

```go
package main

import (
	"encoding/json"
	"github.com/pterm/pterm"
	"github.com/Gaz492/gotracerwifi"
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
```