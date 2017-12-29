package fileReader

import (
	"encoding/binary"
	"fmt"
	"os"
)

// ReadAndCheck reads the data from the file and check whether it has been read
func ReadAndCheck(file *os.File, toRead uint32) ([]byte, error) {
	buffer := make([]byte, toRead)
	read, _ := file.Read(buffer)

	if toRead != uint32(read) {
		return buffer, fmt.Errorf(
			"Expected %d bytes to be read, %d read instead",
			toRead,
			read,
		)
	}

	return buffer, nil
}

// BytesToUint32LE converts an array byte in an uint32
func BytesToUint32LE(slice []byte) uint32 {
	return binary.LittleEndian.Uint32(slice)
}

// BytesToUint16LE converts an array byte in an uint16
func BytesToUint16LE(slice []byte) uint16 {
	return binary.LittleEndian.Uint16(slice)
}
