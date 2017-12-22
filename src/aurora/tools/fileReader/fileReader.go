package fileReader

import (
	"aurora/tools"
	"encoding/binary"
	"fmt"
	"os"
)

// ReadAndCheck reads the data from the file and check whether it has been read
func ReadAndCheck(file *os.File, toRead uint32) []byte {
	buffer := make([]byte, toRead)
	read, err := file.Read(buffer)
	tools.EasyPanic(err)

	if toRead != uint32(read) {
		panic(
			fmt.Sprintf(
				"Expected %d bytes to be read, %d read instead",
				toRead,
				read,
			),
		)
	}

	return buffer
}

// BytesToUint32LE converts an array byte in an uint32
func BytesToUint32LE(slice []byte) uint32 {
	return binary.LittleEndian.Uint32(slice)
}

// BytesToUint16LE converts an array byte in an uint16
func BytesToUint16LE(slice []byte) uint16 {
	return binary.LittleEndian.Uint16(slice)
}
