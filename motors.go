package gobrickpi

import "log"

func GetMotorStatus(port byte) (state uint8, power int8, position int32, dps int16) {
	/*Read motor status
	  Works on on port at a time

	  Returns a list

	  flags -> 8-buts of bit flags that indicate motor status:
	  	bit 0 = LOW_VOLTAGE_FLOAT -> The motors are automatically disabled when the battery voltage is too low
	  	bit 1 = OVERLOAED -> The motors aren't close to the target (apploes to position control and dps speed control
	  	power = The raw PWM power in percent (-100 to 100)
	  	encoder = The encoder posotion
	  	dps = The current speed in degress per second
	*/
	var msgType byte
	switch port {

	case PORT_A:
		msgType = bp_MSG_GET_MOTOR_A_STATUS

	case PORT_B:
		msgType = bp_MSG_GET_MOTOR_B_STATUS

	case PORT_C:
		msgType = bp_MSG_GET_MOTOR_C_STATUS

	case PORT_D:
		msgType = bp_MSG_GET_MOTOR_D_STATUS

	default:
		log.Fatal("Can only get status of one motor port at a time!")
	}

	dataIn := []byte{address, msgType, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	spiTransfer(dataIn)

	if dataIn[3] == 0xA5 {

		state = uint8(dataIn[4])

		power = int8(dataIn[5])

		position = int32((int(dataIn[6]) << 24) | (int(dataIn[7]) << 16) | (int(dataIn[8]) << 8) | int(dataIn[9]))

		dps = int16((int(dataIn[10]) << 8) | int(dataIn[11]))

	}

	return state, power, position, dps
}
func SetMotorPower(port byte, power int8) {
	/* Set the motor power in percent
	port -- The Motor port(s). PORT_A, PORT_B, PORT_C, and/or PORT_D.
	power -- The power from -100 to 100, or -128 for float*/
	dataIn := []byte{address, bp_MSG_SET_MOTOR_POWER, port, byte(power)}
	spiTransfer(dataIn)
}

func SetMotorLimits(port byte, power byte, dps int) {
	/* Set the motor speed limit

	power = Power limit in percent (0 to 100) with 0 being no limit
	dps = Speed limit in degress per second with 0 being no limit*/
	dataIn := []byte{address, bp_MSG_SET_MOTOR_LIMITS, port, power, byte((dps >> 8) & 0xFF), byte(dps & 0xFF)}
	spiTransfer(dataIn)
}

func SetMotorPosition(port byte, position int32) {
	/* Set the motor target position in degress

	 */
	dataIn := []byte{address, bp_MSG_SET_MOTOR_POSITION, port, byte((position >> 24) & 0xFF), byte((position >> 16) & 0xFF), byte((position >> 8) & 0xFF), byte(position & 0xFF)}
	spiTransfer(dataIn)

}

func SetMotorPositionKP(port byte, kp byte) {
	/* Set the motor target position KP constant

	If you set kp higher, the motor will be more responsive to errors in position, at the cost of perhaps overshooting and oscillating.
	kd slows down the motor as it approaches the target, and helps to prevent overshoot.
	In general, if you increase kp, you should also increase kd to keep the motor from overshooting and oscillating.

	Default kp to 25*/
	if kp == 0 {
		kp = 25
	}
	dataIn := []byte{address, bp_MSG_SET_MOTOR_POSITION_KP, port, kp}
	spiTransfer(dataIn)
}

func SetMotorPositionKD(port byte, kd byte) {
	/* Set the motor targer position KD constant

	If you set kp higher, the motor will be more responsive to errors in position, at the cost of perhaps overshooting and oscillating.
	kd slows down the motor as it approaches the target, and helps to prevent overshoot.
	In general, if you increase kp, you should also increase kd to keep the motor from overshooting and oscillating.

	Default KD to 70*/
	if kd == 0 {
		kd = 70
	}

	dataIn := []byte{address, bp_MSG_SET_MOTOR_POSITION_KD, port, kd}
	spiTransfer(dataIn)
}

func setMotorDPS(port byte, dps int) {
	/* Set the motor target speed in degress per second*/
	dataIn := []byte{address, bp_MSG_SET_MOTOR_DPS, port, byte((dps >> 8) & 0xFF), byte(dps & 0xFF)}
	spiTransfer(dataIn)
}
