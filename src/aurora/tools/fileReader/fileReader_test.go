package fileReader_test

import (
	"aurora/tools/fileReader"
	"aurora/tools/test"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestBytesToUint16LE(t *testing.T) {
	var bytes = []byte{0x04, 0x01} // 260

	test.ExpectUint16(t, fmt.Sprintf("%v", bytes), 260, fileReader.BytesToUint16LE(bytes))
}

// Internally tests ReadAndCheck too
func TestReadStringFromBytes(t *testing.T) {
	var file, _ = os.Open("test/testRead.txt")
	defer file.Close()

	var readerBag = fileReader.ByteReaderBag{File: file}
	var str = fileReader.ReadStringFromBytes(&readerBag, 32)

	test.ExpectString(t, "Test reading", "This file is made of 256 bytes.\n", str)
	if readerBag.Err != nil {
		t.Errorf("Expected no errors, got '%v'", readerBag.Err)
	}

	// ReadAndCheck: generate an error as there are no bytes left
	str = fileReader.ReadStringFromBytes(&readerBag, 1)
	test.ExpectString(t, "Test reading with error generated", "", str)
	if readerBag.Err == nil {
		t.Error("Expected error, got nil")
	}

	// ReadAndCheck: even if the file has been rewinded, the previous error
	// will abort subsequent reads
	readerBag.File.Seek(0, os.SEEK_SET)
	str = fileReader.ReadStringFromBytes(&readerBag, 1)
	test.ExpectString(t, "Test reading with previous error", "", str)
}

// Internally tests BytesToUint32LE too
func TestReadUint32FromBytes(t *testing.T) {
	var file, _ = ioutil.TempFile(os.TempDir(), "")
	var filePath = file.Name()
	defer file.Close()
	defer os.Remove(filePath)

	var writer = bufio.NewWriter(file)
	// Little endian pow disposition:
	//     16^1 16^0 |     16^3 16^2 |     16^5    16^4 |     16^7 16^6
	// 0 x 0    f    | 0 x 0    1    | 0 x 1       0    | 0 x 0    4
	//     0    15   |     0    256  |     1048576 0    |     0    67108864
	// --------------------------------------------------------------------
	//                                                     TOTAL = 68157711
	writer.Write([]byte{0x0f, 0x01, 0x10, 0x04})
	writer.Flush()

	file.Seek(0, os.SEEK_SET)

	var readerBag = fileReader.ByteReaderBag{File: file}
	var number = fileReader.ReadUint32FromBytes(&readerBag)
	if readerBag.Err != nil {
		t.Errorf("Expected no errors, got '%v'", readerBag.Err)
	}

	test.ExpectUint32(t, "Number test", 68157711, number)

	fileReader.ReadUint32FromBytes(&readerBag)
	if readerBag.Err == nil {
		t.Errorf("Expected no bytes to read, but got no error")
	}
}

// Internally tests BytesToUint32LE too
func TestReadUint16FromBytes(t *testing.T) {
	var file, _ = ioutil.TempFile(os.TempDir(), "")
	var filePath = file.Name()
	defer file.Close()
	defer os.Remove(filePath)

	var writer = bufio.NewWriter(file)
	// Little endian pow disposition:
	//     16^1 16^0 |     16^3 16^2
	// 0 x 0    4    | 0 x 0    1
	//     0    4    |     0    256
	// -----------------------------
	//              TOTAL = 260
	writer.Write([]byte{0x04, 0x01})
	writer.Flush()

	file.Seek(0, os.SEEK_SET)

	var readerBag = fileReader.ByteReaderBag{File: file}
	var number = fileReader.ReadUint16FromBytes(&readerBag)
	if readerBag.Err != nil {
		t.Errorf("Expected no errors, got '%v'", readerBag.Err)
	}

	test.ExpectUint16(t, "Number test", 260, number)

	fileReader.ReadUint16FromBytes(&readerBag)
	if readerBag.Err == nil {
		t.Errorf("Expected no bytes to read, but got no error")
	}
}
