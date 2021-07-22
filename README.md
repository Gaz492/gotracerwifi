# Go Tracer WiFi

A Go package to communicate with an EPSolar Tracer A/AN charge controller equipped with the eBox-Wifi-01 or HiFlying Elfin EW11

Only tested with EPSolar Tracer3210AN equipped with a eBox-Wifi-01 and HiFlying Elfin EW11

## Example

### TCP example
```go
package main

import (
	"encoding/json"
	"github.com/pterm/pterm"
	"github.com/Gaz492/gotracerwifi"
	"time"
)

func main()  {
	tracer, err := goTracerWiFi.Status("192.168.1.181", "8088", 5 * time.Second, "TCP")
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

### RTU over TCP example
```go
package main

import (
	"encoding/json"
	"github.com/pterm/pterm"
	"github.com/Gaz492/gotracerwifi"
	"time"
)

func main()  {
	tracer, err := goTracerWiFi.Status("192.168.1.181", "8088", 5 * time.Second, "RTU_TCP")
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