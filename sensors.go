package gobrickpi

import (
	"log"
)

var sensorPortConfig [4]byte

func SetSensorType(port byte, sensorType byte) {

	var i uint
	for i = 0; i <= 3; i++ {

		if uint(port)&(1<<i) != 0 {
			sensorPortConfig[i] = sensorType
		}
	}

	outData := []byte{address, bp_MSG_SET_SENSOR_TYPE, port, sensorType}
	spiTransfer(outData)
}

func GetSensorValue(port uint) (sensorValue int) {

	var msgType byte
	var outData []byte
	portIndex := 0

	switch port {
	case PORT_1:
		msgType = bp_MSG_GET_SENSOR_1
		portIndex = 0
	case PORT_2:
		msgType = bp_MSG_GET_SENSOR_2
		portIndex = 1
	case PORT_3:
		msgType = bp_MSG_GET_SENSOR_3
		portIndex = 2
	case PORT_4:
		msgType = bp_MSG_GET_SENSOR_4
		portIndex = 3
	default:
		log.Fatal("Can only set one sensor port at a time")

	}

	switch sensorPortConfig[portIndex] {

	case SENSOR_TYPE_TOUCH,
		SENSOR_TYPE_TOUCH_NXT,
		SENSOR_TYPE_TOUCH_EV3,
		SENSOR_TYPE_NXT_ULTRASONIC,
		SENSOR_TYPE_EV3_COLOR_REFLECTED,
		SENSOR_TYPE_EV3_COLOR_AMBIENT,
		SENSOR_TYPE_EV3_ULTRASONIC_LISTEN,
		SENSOR_TYPE_EV3_INFRARED_PROXIMITY:

		outData = []byte{address, msgType, 0, 0, 0, 0, 0}
		spiTransfer(outData)
		sensorValue = int(outData[6])

	default:
		// If sensor was not setup on the port abort the program
		log.Fatalf("No Sensor setup on port %v", portIndex+1)
	}

	if outData[5] != SENSOR_STATE_VALID_DATA {
		log.Fatalf("Invalid data from sensor code returned %v", outData)
	}

	if !(outData[4] == sensorPortConfig[portIndex] || (sensorPortConfig[portIndex] == SENSOR_TYPE_TOUCH && (outData[4] == SENSOR_TYPE_TOUCH_NXT || outData[4] == SENSOR_TYPE_TOUCH_EV3))) {
		log.Fatalf("THe sensor configured on port %v does not match connected sensor: %v", outData[4], sensorPortConfig[portIndex])
	}
	return sensorValue
}
