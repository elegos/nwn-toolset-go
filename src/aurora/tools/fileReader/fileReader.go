package fileReader

import (
	"encoding/binary"
	"fmt"
	"os"
)

// ByteReaderBag container used to gracefully manage file reading errors
type ByteReaderBag struct {
	File *os.File
	Err  error
}

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

// ReadBytes read toRead bytes and store any error in the bag
func ReadBytes(readerBag *ByteReaderBag, toRead uint32) []byte {
	if readerBag.Err != nil {
		return nil
	}

	bytes, err := ReadAndCheck(readerBag.File, toRead)

	if err != nil {
		readerBag.Err = err

		return nil
	}

	return bytes
}

// ReadStringFromBytes helper to transform ReadBytes into a string
func ReadStringFromBytes(readerBag *ByteReaderBag, toRead uint32) string {
	var bytes = ReadBytes(readerBag, toRead)

	if readerBag.Err != nil {
		return ""
	}

	return string(bytes)
}

// ReadUint32FromBytes read uint32 (little endian) from bytes
func ReadUint32FromBytes(readerBag *ByteReaderBag) uint32 {
	var bytes = ReadBytes(readerBag, 4)

	if readerBag.Err != nil {
		return 0
	}

	return BytesToUint32LE(bytes)
}

// ReadUint16FromBytes read uint16 (little endian) from bytes
func ReadUint16FromBytes(readerBag *ByteReaderBag) uint16 {
	var bytes = ReadBytes(readerBag, 2)

	if readerBag.Err != nil {
		return 0
	}

	return BytesToUint16LE(bytes)
}
