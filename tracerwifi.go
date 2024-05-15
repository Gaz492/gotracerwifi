package gotracerwifi

import (
	"errors"
	"fmt"
	"github.com/grid-x/modbus"
	"time"
)

type Client struct {
	modbus.Client
}

func NewTCPClient(ip string, port string, timeout time.Duration, protocol string) (Client, error) {
	var err error
	if protocol == "TCP" {
		handler := modbus.NewTCPClientHandler(fmt.Sprintf("%s:%s", ip, port))
		handler.SlaveID = 1
		handler.Timeout = timeout

		err = handler.Connect()
		if err != nil {
			return Client{}, err
		}
		client := modbus.NewClient(handler)
		return Client{client}, nil
	} else if protocol == "RTU_TCP" {
		handler := modbus.NewRTUOverTCPClientHandler(fmt.Sprintf("%s:%s", ip, port))
		handler.SlaveID = 1
		handler.Timeout = timeout

		err = handler.Connect()
		if err != nil {
			return Client{}, err
		}
		client := modbus.NewClient(handler)
		return Client{client}, nil
	} else {
		err = errors.New("invalid protocol")
		return Client{}, err
	}
}

// Load stuff
func (c Client) GetLoadStatus() (int, error) {
	status, err := c.ReadCoils(0x2, 1)
	if err != nil {
		return 0, err
	}
	return int(status[0]), nil
}

func (c Client) SetLoadOff() error {
	_, err := c.WriteSingleCoil(0x2, 0x0000)
	return err
}

func (c Client) SetLoadOn() error {
	_, err := c.WriteSingleCoil(0x2, 0xFF00)
	return err
}

// Solar stuff
func (c Client) GetPvVoltage() (float32, error) {
	r, err := c.readInputRegister(0x3100, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

func (c Client) GetPvCurrent() (float32, error) {
	r, err := c.readInputRegister(0x3101, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

func (c Client) GetPvPower() (float32, error) {
	r, err := c.readInputRegister(0x3102, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

// Battery stuff

func (c Client) GetBatteryVoltage() (float32, error) {
	r, err := c.readInputRegister(0x3104, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

func (c Client) GetBatteryCurrent() (float32, error) {
	r, err := c.readInputRegister(0x3105, 1)
	if err != nil {
		return 0, err
	}
	bc := unpack(r) / 100
	if bc > 32768 {
		bc = bc - 65536
	}
	return bc / 100, nil
}

func (c Client) GetBatteryMaxVoltage() (float32, error) {
	r, err := c.readInputRegister(0x3302, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

func (c Client) GetBatteryMin() (float32, error) {
	r, err := c.readInputRegister(0x3303, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

func (c Client) GetTemp() (float32, error) {
	r, err := c.readInputRegister(0x3110, 1)
	if err != nil {
		return 0, err
	}
	t := unpack(r)
	if t > 32768 {
		t = t - 65536
	}
	return t / 100, nil
}

func (c Client) GetBatterySOC() (float32, error) {
	r, err := c.readInputRegister(0x311A, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

// Load stuff

func (c Client) GetLoadVoltage() (float32, error) {
	r, err := c.readInputRegister(0x310C, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

func (c Client) GetLoadCurrent() (float32, error) {
	r, err := c.readInputRegister(0x310D, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

func (c Client) GetLoadPower() (float32, error) {
	r, err := c.readInputRegister(0x310E, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

// Stats stuff

func (c Client) GeneratedDay() (float32, error) {
	r, err := c.readInputRegister(0x330C, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

func (c Client) GeneratedMonth() (float32, error) {
	r, err := c.readInputRegister(0x330E, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

func (c Client) GeneratedAnnual() (float32, error) {
	r, err := c.readInputRegister(0x3310, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

func (c Client) GeneratedTotal() (float32, error) {
	r, err := c.readInputRegister(0x3312, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

func (c Client) ConsumedDay() (float32, error) {
	r, err := c.readInputRegister(0x3304, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

func (c Client) ConsumedMonth() (float32, error) {
	r, err := c.readInputRegister(0x3306, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

func (c Client) ConsumedAnnual() (float32, error) {
	r, err := c.readInputRegister(0x3308, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}

func (c Client) ConsumedTotal() (float32, error) {
	r, err := c.readInputRegister(0x330A, 1)
	if err != nil {
		return 0, err
	}
	return unpack(r) / 100, nil
}
