package gobrickpi

import (
	"log"
)

const speedHz = 500000
const mode = 0
const bitsPerWord = 8
const spiDev = "/dev/spidev0.1"

const (
	bp_MSG_NONE = iota //0

	bp_MSG_GET_MANUFACTURER     = iota //1
	bp_MSG_GET_NAME             = iota
	bp_MSG_GET_HARDWARE_VERSION = iota
	bp_MSG_GET_FIRMWARE_VERSION = iota
	bp_MSG_GET_ID               = iota
	bp_MSG_SET_LED              = iota
	bp_MSG_GET_VOLTAGE_3V3      = iota
	bp_MSG_GET_VOLTAGE_5V       = iota
	bp_MSG_GET_VOLTAGE_9V       = iota
	bp_MSG_GET_VOLTAGE_VCC      = iota
	bp_MSG_SET_ADDRESS          = iota //11

	bp_MSG_SET_SENSOR_TYPE = iota //12
	bp_MSG_GET_SENSOR_1    = iota
	bp_MSG_GET_SENSOR_2    = iota
	bp_MSG_GET_SENSOR_3    = iota
	bp_MSG_GET_SENSOR_4    = iota //16

	bp_MSG_12C_TRANSACT_1 = iota //17
	bp_MSG_12C_TRANSACT_2 = iota
	bp_MSG_12C_TRANSACT_3 = iota
	bp_MSG_12C_TRANSACT_4 = iota //20

	bp_MSG_SET_MOTOR_POWER       = iota //21
	bp_MSG_SET_MOTOR_POSITION    = iota
	bp_MSG_SET_MOTOR_POSITION_KP = iota
	bp_MSG_SET_MOTOR_POSITION_KD = iota
	bp_MSG_SET_MOTOR_DPS         = iota
	bp_MSG_SET_MOTOR_DPS_KP      = iota
	bp_MSG_SET_MOTOR_DPS_KD      = iota
	bp_MSG_SET_MOTOR_LIMITS      = iota //28

	bp_MSG_OFFSET_MOTOR_ENCODER = iota //29
	bp_MSG_GET_MOTOR_A_ENCODER  = iota
	bp_MSG_GET_MOTOR_B_ENCODER  = iota
	bp_MSG_GET_MOTOR_C_ENCODER  = iota
	bp_MSG_GET_MOTOR_D_ENCODER  = iota //33

	bp_MSG_GET_MOTOR_A_STATUS = iota //34
	bp_MSG_GET_MOTOR_B_STATUS = iota
	bp_MSG_GET_MOTOR_C_STATUS = iota
	bp_MSG_GET_MOTOR_D_STATUS = iota //37

)

const (
	_                  = iota
	SENSOR_TYPE_NONE   = iota
	SENSOR_TYPE_I2C    = iota
	SENSOR_TYPE_CUSTOM = iota

	SENSOR_TYPE_TOUCH     = iota
	SENSOR_TYPE_TOUCH_NXT = iota
	SENSOR_TYPE_TOUCH_EV3 = iota

	SENSOR_TYPE_NXT_LIGHT_ON  = iota
	SENSOR_TYPE_NXT_LIGHT_OFF = iota

	SENSOR_TYPE_NXT_COLOR_RED   = iota
	SENSOR_TYPE_NXT_COLOR_GREEN = iota
	SENSOR_TYPE_NXT_COLOR_BLUE  = iota
	SENSOR_TYPE_NXT_COLOR_FULL  = iota
	SENSOR_TYPE_NXT_COLOR_OFF   = iota

	SENSOR_TYPE_NXT_ULTRASONIC = iota

	SENSOR_TYPE_EV3_GYRO_ABS = iota
	SENSOR_TYPE_EV3_GYRO_DPS = iota
	SENSOR_TYPE_EV3_ABS_DPS  = iota

	SENSOR_TYPE_EV3_COLOR_REFLECTED        = iota
	SENSOR_TYPE_EV3_COLOR_AMBIENT          = iota
	SENSOR_TYPE_EV3_COLOR_COLOR            = iota
	SENSOR_TYPE_EV3_COLOR_RAW_REFLECTED_   = iota
	SENSOR_TYPE_EV3_COLOR_COLOR_COMPONENTS = iota

	SENSOR_TYPE_EV3_ULTRASONIC_CM     = iota
	SENSOR_TYPE_EV3_ULTRASONIC_INCHES = iota
	SENSOR_TYPE_EV3_ULTRASONIC_LISTEN = iota

	SENSOR_TYPE_EV3_INFRARED_PROXIMITY = iota
	SENSOR_TYPE_EV3_INFRARED_SEEK      = iota
	SENSOR_TYPE_EV3_INFRARED_REMOTE    = iota
)

const (
	SENSOR_STATE_VALID_DATA     = iota
	SENSOR_STATE_NOT_CONFIGURED = iota
	SENSOR_STATE_CONFIGURING    = iota
	SENSOR_STATE_NO_DATA        = iota
	SENSOR_STATE_I2C_ERROR      = iota
)

const (
	PORT_1 = 1
	PORT_2 = 2
	PORT_3 = 4
	PORT_4 = 8
)

const (
	PORT_A = 1
	PORT_B = 2
	PORT_C = 4
	PORT_D = 8
)

type touchSensor struct {
	pressed bool
}

func spiTransfer(outData []byte) {

	err := BP.Transfer(outData)

	if err != nil {

		log.Fatal(err)
	}
}

func spiWrite8(msg byte, value byte) {
	outData := []byte{address, msg, value}
	spiTransfer(outData)
}

func spiWrite16(msg byte, value int16) {

	outData := []byte{address, msg, byte((value >> 8) & 0xFF), byte(value & 0xFF)}
	spiTransfer(outData)
}

func spiWrite24(msg byte, value int) {

	outData := []byte{address, msg, byte((value >> 16) & 0xFF), byte((value >> 8) & 0xFF), byte(value & 0xFF)}
	spiTransfer(outData)
}

func spiWrite32(msg byte, value int32) {

	outData := []byte{address, msg, byte((value >> 24) & 0xFF), byte((value >> 16) & 0xFF), byte((value >> 8) & 0xFF), byte(value & 0xFF)}
	spiTransfer(outData)
}

func spiRead16(msg byte) (reply int) {

	outData := []byte{address, msg, 0, 0, 0, 0}
	err := BP.Transfer(outData)

	if err != nil {
		log.Fatal(err)
	}

	if outData[3] != 0xA5 {
		log.Fatal(err)
	}

	return (int(outData[4]) << 8) | int(outData[5])
}

func spiRead32(msg byte) (reply int32) {

	outData := []byte{address, msg, 0, 0, 0, 0, 0, 0}
	err := BP.Transfer(outData)

	if err != nil {
		log.Fatal(err)
	}

	if outData[3] != 0xA5 {
		log.Fatal(err)
	}
	return int32((int(outData[4]) << 24) | (int(outData[5]) << 16) | (int(outData[6]) << 8) | int(outData[7]))
}
