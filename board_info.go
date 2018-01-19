package gobrickpi

import "fmt"

func GetFirmwareVersion() {

	version := spiRead32(bp_MSG_GET_FIRMWARE_VERSION)
	fmt.Printf("%d.%d.%d", (version / 1000000), ((version / 1000) % 1000), (version % 1000))
}

func ReadVoltage3V3() (voltage float32) {
	return float32(spiRead16(bp_MSG_GET_VOLTAGE_3V3)) / 1000.0
}

func ReadVoltage5V() (voltage float32) {
	return float32(spiRead16(bp_MSG_GET_VOLTAGE_5V)) / 1000.0
}

func ReadVoltage9V() (voltage float32) {
	return float32(spiRead16(bp_MSG_GET_VOLTAGE_9V)) / 1000.0
}

func ReadVoltageBattery() (voltage float32) {
	return float32(spiRead16(bp_MSG_GET_VOLTAGE_VCC)) / 1000.0
}

func GetManufacturer() {
}

func GetBoard() {
}

func GetHardwareVersion() {
}

func GetID() {
}
