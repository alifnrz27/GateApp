package gpio

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var initialized = false

func Init() error {
	if initialized {
		return nil
	}

	if err := rpio.Open(); err != nil {
		return fmt.Errorf("failed to open gpio: %v", err)
	}

	initialized = true
	return nil
}

func Close() {
	if initialized {
		rpio.Close()
	}
}

func TriggerRelay(pinNumber int, durationSeconds int) error {
	if !initialized {
		return fmt.Errorf("gpio not initialized")
	}

	pin := rpio.Pin(pinNumber)
	pin.Output()

	// 🔥 ACTIVE LOW (ubah kalau perlu)
	pin.Low()
	time.Sleep(time.Duration(durationSeconds) * time.Second)
	pin.High()

	return nil
}
