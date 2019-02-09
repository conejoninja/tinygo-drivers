package main

import (
	"machine"
	"time"

	"github.com/tinygo-org/drivers/easystepper"
)

func main() {
	//motor := easystepper.New(8, 9, 10, 11)
	motor := easystepper.New(machine.P13, machine.P15, machine.P14, machine.P16, 4)

	for {
		println("CLOCKWISE")
		motor.Step(2050)
		time.Sleep(time.Millisecond * 1000)

		println("COUNTERCLOCKWISE")
		motor.Step(-2050)
		time.Sleep(time.Millisecond * 1000)
	}
}
