# gobrickpi
Go library to control brick pi robot

Using Go SPI and Raspberry Pi

TODO: There is a lot to do :)

The following functionailty has been implmented.

1. Read/Write from the BrickPi using SPI
2. Read BrickPi voltages
3. Set BrickPi LED
4. Set Sensor type (basic version)
5. Get Sensor value (basic version)
6. Read NXT Touch Sensor


# Example Usage

Below is a very basic example that shows how to use a Lego Touch Sensor in your program

```go
/*Simple example to show how to use Lego Touch Sensor connected to port 1
 */
package main

import (
	"fmt"
	"github.com/r3d-dy3-n0-2/gobrickpi"
	"time"
)

func main() {
	//Create connection to BrickPi
	gobrickpi.Start()
	//Set Sensor Port 1 to be a lego touch sensor
	gobrickpi.SetSensorType(gobrickpi.PORT_1, gobrickpi.SENSOR_TYPE_TOUCH)
	//Wait for 50 seconds to ensure sensor is setup correctly
	time.Sleep(50 * time.Millisecond)

	//Loop forever reading the value of the touch sensor on port 1.
	//If pressed will return 1 from sensor if not pressed the sensor will read 0
	for {
		touch_state := gobrickpi.GetSensorValue(gobrickpi.PORT_1)
		if touch_state == 1 {

			fmt.Println("Touch Sensor: HIT")
		} else {
			fmt.Println("Touch Sensor: NOT HIT!")
		}
		time.Sleep(40 * time.Millisecond) //Sleep to stop swamping the CPU
	}

}
```
