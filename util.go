package gotracerwifi

import (
	"github.com/grid-x/modbus"
	"github.com/pterm/pterm"
	"time"
)

// Credit: https://github.com/spagettikod/gotracer
func unpack(slice []byte) float32 {
	var v uint32
	for i, b := range slice {
		shift := uint((len(slice) - 1 - i) * 8)
		v += uint32(b) << shift
	}
	return float32(v)
}

func requestInputRegister(client modbus.Client, address uint16, quantity uint16) (results []byte) {
	results, err := client.ReadInputRegisters(address, quantity)
	if err != nil {
		pterm.Error.Println(time.Now().UTC().Format(time.RFC1123Z), err)
	}
	return
}

func (c Client) readInputRegister(address uint16, quantity uint16) ([]byte, error) {
	results, err := c.ReadInputRegisters(address, quantity)
	if err != nil {
		return nil, err
	}
	return results, nil
}
