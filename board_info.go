package gobrickpi

import "fmt"
import "log"

func GetFirmwareVersion() (firmware string) {

	version := spiRead32(bp_MSG_GET_FIRMWARE_VERSION)
	firmware = fmt.Sprintf("%d.%d.%d", (version / 1000000), ((version / 1000) % 1000), (version % 1000))
	return firmware
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

func GetManufacturer() (manufactor string) {
	outData := []byte{address, bp_MSG_GET_MANUFACTURER, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	err := BP.Transfer(outData)

	if err != nil {
		log.Fatal(err)
	}

	if outData[3] == 0XA5 {
		for i := 4; i < 24; i++ {
			manufactor += string(rune(outData[i]))
		}
	}

	return manufactor
}

func GetBoard() (boardName string) {
	outData := []byte{address, bp_MSG_GET_NAME, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	err := BP.Transfer(outData)

	if err != nil {
		log.Fatal(err)
	}

	if outData[3] == 0XA5 {
		for i := 4; i < 24; i++ {
			boardName += string(rune(outData[i]))
		}
	}
	return boardName
}

func GetHardwareVersion() (hardwareVersion string) {
	version := spiRead32(bp_MSG_GET_HARDWARE_VERSION)
	hardwareVersion = fmt.Sprintf("%d.%d.%d", (version / 1000000), ((version / 1000) % 1000), (version % 1000))
	return hardwareVersion
}

func GetID() (id string) {
	outData := []byte{address, bp_MSG_GET_ID, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	err := BP.Transfer(outData)

	if err != nil {
		log.Fatal(err)
	}

	if outData[3] == 0XA5 {
		id = fmt.Sprintf("%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X", outData[4], outData[5], outData[6], outData[7], outData[8], outData[9], outData[10], outData[11], outData[12], outData[13], outData[14], outData[15], outData[16], outData[17], outData[18], outData[19])
	}
	return id
}
