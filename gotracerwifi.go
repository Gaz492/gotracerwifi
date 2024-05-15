package gotracerwifi

import (
	"errors"
	"fmt"
	"github.com/grid-x/modbus"
	"github.com/pterm/pterm"
	"time"
)

type Response struct {
	Solar     Solar     `json:"solar"`
	Battery   Battery   `json:"battery"`
	Load      Load      `json:"load"`
	Stats     Stats     `json:"stats"`
	Timestamp time.Time `json:"timestamp"`
}
type Solar struct {
	Voltage float32 `json:"voltage"`
	Current float32 `json:"current"`
	Power   float32 `json:"power"`
}
type Battery struct {
	Voltage       float32 `json:"voltage"`
	Current       float32 `json:"current"`
	MaxVoltage    float32 `json:"max_voltage"`
	MinVoltage    float32 `json:"min_voltage"`
	Temp          float32 `json:"temp"`
	BatteryStatus string  `json:"battery_status"`
	ChargeStatus  string  `json:"charge_status"`
}
type Load struct {
	Voltage float32 `json:"voltage"`
	Current float32 `json:"current"`
	Power   float32 `json:"power"`
	Status  int     `json:"status"`
}
type Generated struct {
	Day    float32 `json:"day"`
	Month  float32 `json:"month"`
	Annual float32 `json:"annual"`
	Total  float32 `json:"total"`
}
type Consumed struct {
	Day    float32 `json:"day"`
	Month  float32 `json:"month"`
	Annual float32 `json:"annual"`
	Total  float32 `json:"total"`
}
type Energy struct {
	Generated Generated `json:"generated"`
	Consumed  Consumed  `json:"consumed"`
}
type Stats struct {
	Energy Energy `json:"energy"`
}

func Status(ip string, port string, timeout time.Duration, protocol string) (r Response, err error) {
	var client modbus.Client
	if protocol == "TCP" {
		handler := modbus.NewTCPClientHandler(fmt.Sprintf("%s:%s", ip, port))
		handler.SlaveID = 1
		handler.Timeout = timeout
		//handler.Logger = log.New(os.Stdout, "INFO: ", log.LstdFlags)

		err = handler.Connect()
		if err != nil {
			return
		}
		defer handler.Close()
		client = modbus.NewClient(handler)
	} else if protocol == "RTU_TCP" {
		handler := modbus.NewRTUOverTCPClientHandler(fmt.Sprintf("%s:%s", ip, port))
		handler.SlaveID = 1
		handler.Timeout = timeout
		//handler.Logger = log.New(os.Stdout, "INFO: ", log.LstdFlags)

		err = handler.Connect()
		if err != nil {
			return
		}
		defer handler.Close()
		client = modbus.NewClient(handler)
	} else {
		err = errors.New("invalid protocol")
		return
	}

	r.Timestamp = time.Now().UTC()

	r.Solar.Voltage = unpack(requestInputRegister(client, 0x3100, 1)) / 100
	r.Solar.Current = unpack(requestInputRegister(client, 0x3101, 1)) / 100
	r.Solar.Power = unpack(requestInputRegister(client, 0x3102, 1)) / 100

	r.Battery.Voltage = unpack(requestInputRegister(client, 0x3104, 1)) / 100
	batCurrent := unpack(requestInputRegister(client, 0x3105, 1)) / 100
	if batCurrent > 32768 {
		batCurrent = batCurrent - 65536
	}
	r.Battery.Current = batCurrent / 100
	r.Battery.MaxVoltage = unpack(requestInputRegister(client, 0x3302, 1)) / 100
	r.Battery.MinVoltage = unpack(requestInputRegister(client, 0x3303, 1)) / 100
	batTemp := unpack(requestInputRegister(client, 0x3110, 1))
	if batTemp > 32768 {
		batTemp = batTemp - 65536
	}
	r.Battery.Temp = batTemp / 100

	r.Load.Voltage = unpack(requestInputRegister(client, 0x310C, 1)) / 100
	r.Load.Current = unpack(requestInputRegister(client, 0x310D, 1)) / 100
	r.Load.Power = unpack(requestInputRegister(client, 0x310E, 1)) / 100
	r.Load.Status = OldGetLoadStatus(client)

	r.Stats.Energy.Generated.Day = unpack(requestInputRegister(client, 0x330C, 1)) / 100
	r.Stats.Energy.Generated.Month = unpack(requestInputRegister(client, 0x330E, 1)) / 100
	r.Stats.Energy.Generated.Annual = unpack(requestInputRegister(client, 0x3310, 1)) / 100
	r.Stats.Energy.Generated.Total = unpack(requestInputRegister(client, 0x3312, 1)) / 100

	r.Stats.Energy.Consumed.Day = unpack(requestInputRegister(client, 0x3304, 1)) / 100
	r.Stats.Energy.Consumed.Month = unpack(requestInputRegister(client, 0x3306, 1)) / 100
	r.Stats.Energy.Consumed.Annual = unpack(requestInputRegister(client, 0x3308, 1)) / 100
	r.Stats.Energy.Consumed.Total = unpack(requestInputRegister(client, 0x330A, 1)) / 100

	return
}

func OldToggleLoad(ip string, port string, timeout time.Duration, protocol string) (err error) {
	var client modbus.Client
	if protocol == "TCP" {
		handler := modbus.NewTCPClientHandler(fmt.Sprintf("%s:%s", ip, port))
		handler.SlaveID = 1
		handler.Timeout = timeout
		//handler.Logger = log.New(os.Stdout, "INFO: ", log.LstdFlags)

		err = handler.Connect()
		if err != nil {
			return
		}
		defer handler.Close()
		client = modbus.NewClient(handler)
	} else if protocol == "RTU_TCP" {
		handler := modbus.NewRTUOverTCPClientHandler(fmt.Sprintf("%s:%s", ip, port))
		handler.SlaveID = 1
		handler.Timeout = timeout
		//handler.Logger = log.New(os.Stdout, "INFO: ", log.LstdFlags)

		err = handler.Connect()
		if err != nil {
			return
		}
		defer handler.Close()
		client = modbus.NewClient(handler)
	} else {
		err = errors.New("invalid protocol")
		return
	}

	loadStatus := OldGetLoadStatus(client)

	var res []byte
	if loadStatus == 0 {
		res, err = client.WriteSingleCoil(0x2, 0xFF00)
	} else {
		res, err = client.WriteSingleCoil(0x2, 0x0000)
	}
	pterm.Info.Println("Response:", res)
	return
}

func OldGetLoadStatus(client modbus.Client) (status int) {
	loadStatus, err := client.ReadCoils(0x2, 1)
	if err != nil {
		pterm.Error.Println(time.Now().UTC().Format(time.RFC1123Z), err)
		status = 0
		return
	}
	status = int(loadStatus[0])
	return
}
