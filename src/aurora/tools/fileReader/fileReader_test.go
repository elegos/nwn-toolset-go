package fileReader_test

import (
	"aurora/tools/fileReader"
	"aurora/tools/test"
	"fmt"
	"os"
	"testing"
)

func TestBytesToUint32LE(t *testing.T) {
	// Little endian pow disposition:
	//     16^1 16^0 |     16^3 16^2 |     16^5    16^4 |     16^7 16^6
	// 0 x 0    f    | 0 x 0    1    | 0 x 1       0    | 0 x 0    4
	//     0    15   |     0    256  |     1048576 0    |     0    67108864
	// --------------------------------------------------------------------
	//                                                             68157711
	var bytes = []byte{0x0f, 0x01, 0x10, 0x04}

	test.ExpectUint32(t, fmt.Sprintf("%v", bytes), 68157711, fileReader.BytesToUint32LE(bytes))
}

func TestBytesToUint16LE(t *testing.T) {
	var bytes = []byte{0x04, 0x01} // 260

	test.ExpectUint16(t, fmt.Sprintf("%v", bytes), 260, fileReader.BytesToUint16LE(bytes))
}

func TestReadAndCheck(t *testing.T) {
	var file, _ = os.Open("test/testRead.txt")
	defer file.Close()

	bytes, _ := fileReader.ReadAndCheck(file, 32)

	test.ExpectString(t, "Test reading", "This file is made of 256 bytes.\n", string(bytes))

	bytes, err := fileReader.ReadAndCheck(file, 26)
	if err == nil {
		t.Errorf("Expected reading fail, got '%s'", string(bytes))
	}
}
