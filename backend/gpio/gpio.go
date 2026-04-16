// gpio/relay.go
package gpio

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func TriggerRelay(pinNumber int) error {
	err := rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close()

	pin := rpio.Pin(pinNumber)
	pin.Output()

	pin.High()
	time.Sleep(2 * time.Second)
	pin.Low()

	return nil
}

// func TriggerRelay(pinNumber int) error {
// 	fmt.Println("Mock trigger relay (Windows), pin:", pinNumber)
// 	return nil
// }
