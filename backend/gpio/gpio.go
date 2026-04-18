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

	fmt.Println("gpio opened")

	initialized = true
	return nil
}

func Close() {
	if initialized {
		rpio.Close()
	}
}

func TriggerRelay(pinNumber int, durationSeconds int) error {
	if err := rpio.Open(); err != nil {
		fmt.Println("Failed to open GPIO:", err)
		return nil
	}
	defer rpio.Close()

	pin := rpio.Pin(pinNumber)
	pin.Output()

	pin.High()
	time.Sleep(time.Duration(durationSeconds) * time.Second)

	// 🔥 OFF
	pin.Low()
	time.Sleep(time.Duration(durationSeconds) * time.Second)

	fmt.Println("Selesai")

	return nil
}
