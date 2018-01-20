package gobrickpi

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/ecc1/spi"
)

var BP *spi.Device
var address byte = 0 // Address can be set to allow multiple brickpis be chained

func Start() {

	dev, err := spi.Open(spiDev, speedHz, mode)

	if err != nil {
		panic(err)
	}

	BP = dev

	//Create goroutine that will listen for Ctrl-C so that will clean up correctly

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	go func() {

		sig := <-sigs
		fmt.Println(sig)
		SetLed(255)
		BP.Close()
		os.Exit(1)
	}()

	BP.SetMode(mode)
	BP.SetBitsPerWord(bitsPerWord)
	BP.SetLSBFirst(false)
	BP.SetMaxSpeed(speedHz)
	BP.Transfer([]byte{0, 6, 40})
}

func SetLed(ledValue byte) {
	spiWrite8(bp_MSG_SET_LED, ledValue)
}

func ResetAll() {
	SetSensorType(PORT_1+PORT_2+PORT_3+PORT_4, SENSOR_TYPE_NONE)
}

func Close() {
	BP.Close()
}
