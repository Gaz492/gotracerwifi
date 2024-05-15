package main

import (
	"github.com/Gaz492/gotracerwifi"
	"github.com/pterm/pterm"
	"time"
)

func main() {
	c, err := gotracerwifi.NewTCPClient("192.168.1.127", "8899", 5*time.Second, "TCP")
	if err != nil {
		pterm.Fatal.Println(err)
	}
	pterm.Info.Println(c.GetBatterySOC())
}

//func main() {
//	tracer, err := gotracerwifi.Status("192.168.1.127", "8899", 5*time.Second, "TCP")
//	if err != nil {
//		pterm.Fatal.Println(err)
//	}
//
//	response, err := json.Marshal(tracer)
//	if err != nil {
//		pterm.Error.Println(err)
//		return
//	}
//	pterm.Info.Println(string(response))
//
//	err = gotracerwifi.ToggleLoad("192.168.1.127", "8899", 5*time.Second, "TCP")
//	if err != nil {
//		pterm.Fatal.Println(err)
//	}
//}
