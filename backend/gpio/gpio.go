// gpio/relay.go
package gpio

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var initialized = false

// Init GPIO (WAJIB DIPANGGIL SEKALI)
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

// Close GPIO (optional, saat shutdown)
func Close() {
	if initialized {
		rpio.Close()
	}
}

// Trigger relay (ACTIVE LOW)
func TriggerRelay(pinNumber int, duration time.Duration) error {
	if !initialized {
		return fmt.Errorf("gpio not initialized")
	}

	pin := rpio.Pin(pinNumber)
	pin.Output()

	// relay ON (active low)
	pin.Low()
	time.Sleep(duration)

	// relay OFF
	pin.High()

	return nil
}

// func TriggerRelay(pinNumber int) error {
// 	fmt.Println("Mock trigger relay (Windows), pin:", pinNumber)
// 	return nil
// }
