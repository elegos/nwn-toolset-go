package gff

import (
	"aurora/tools/fileReader"
	"os"
	"strings"
)

func extractHeaderFromBytes(bytes []byte) Header {
	var result = Header{}

	result.FileType = strings.Trim(string(bytes[0:4]), "\x00")
	result.FileVersion = strings.Trim(string(bytes[4:8]), "\x00")
	result.StructOffset = fileReader.BytesToUint32LE(bytes[8:12])
	result.StructCount = fileReader.BytesToUint32LE(bytes[12:16])
	result.FieldOffset = fileReader.BytesToUint32LE(bytes[16:20])
	result.FieldCount = fileReader.BytesToUint32LE(bytes[20:24])
	result.LabelOffset = fileReader.BytesToUint32LE(bytes[24:28])
	result.LabelCount = fileReader.BytesToUint32LE(bytes[28:32])
	result.FieldDataOffset = fileReader.BytesToUint32LE(bytes[32:36])
	result.FieldDataCount = fileReader.BytesToUint32LE(bytes[36:40])
	result.FieldIndicesOffset = fileReader.BytesToUint32LE(bytes[40:44])
	result.FieldIndicesCount = fileReader.BytesToUint32LE(bytes[44:48])
	result.ListIndicesOffset = fileReader.BytesToUint32LE(bytes[48:52])
	result.ListIndicesCount = fileReader.BytesToUint32LE(bytes[52:56])

	return result
}

func extractHeaderFromFile(file *os.File) Header {
	var bytes = fileReader.ReadAndCheck(file, 56)

	return extractHeaderFromBytes(bytes)
}
