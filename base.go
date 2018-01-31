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
		ResetAll()
		//time.Sleep(200 * time.Millisecond)
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
	/* Reset all sensor types to NONE and reset all motor values, also return LED function back to brickpi firmware*/

	SetSensorType(PORT_1+PORT_2+PORT_3+PORT_4, SENSOR_TYPE_NONE) // Reset all sensor types
	SetMotorPower(PORT_A+PORT_B+PORT_C+PORT_D, MOTOR_FLOAT)      // Turn off all motors
	SetMotorLimits(PORT_A+PORT_B+PORT_C+PORT_D, 0, 0)
	SetMotorPositionKP(PORT_A+PORT_B+PORT_C+PORT_D, 0)
	SetMotorPositionKD(PORT_A+PORT_B+PORT_C+PORT_D, 0)
	SetLed(255)

}

func Close() {
	BP.Close()
}
