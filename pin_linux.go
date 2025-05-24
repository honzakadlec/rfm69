package rfm69

import "github.com/davecheney/gpio"

const (
	irqPin = gpio.GPIO25
	base   = 571
)

func getPin() (gpio.Pin, error) {
	return gpio.OpenPin(base+irqPin, gpio.ModeInput)
}
