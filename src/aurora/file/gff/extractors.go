package gff

import (
	"aurora/tools/fileReader"
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

// data.Header MUST be initialized
// bytes are the bytes of the entire file
func extractStructArrayFromBytes(bytes []byte, header Header) []StructArrayElement {
	var result = []StructArrayElement{}

	var i = uint32(0)
	var index = header.StructOffset
	for ; i < header.StructCount; i++ {
		result = append(result, StructArrayElement{
			Type:             fileReader.BytesToUint32LE(bytes[index : index+4]),
			DataOrDataOffset: fileReader.BytesToUint32LE(bytes[index+4 : index+8]),
			FieldCount:       fileReader.BytesToUint32LE(bytes[index+8 : index+12]),
		})

		index += 12
	}

	return result
}

// bytes are the bytes of the entire file
func extractFieldArrayFromBytes(bytes []byte, header Header) []FieldArrayElement {
	var result = []FieldArrayElement{}

	var i = uint32(0)
	var index = header.FieldOffset
	for ; i < header.FieldCount; i++ {
		result = append(result, FieldArrayElement{
			Type:             fileReader.BytesToUint32LE(bytes[index : index+4]),
			LabelIndex:       fileReader.BytesToUint32LE(bytes[index+4 : index+8]),
			DataOrDataOffset: fileReader.BytesToUint32LE(bytes[index+8 : index+12]),
		})

		index += 12
	}

	return result
}

func extractLabelArrayFromBytes(bytes []byte, header Header) []string {
	var result = []string{}

	var i = uint32(0)
	var index = header.LabelOffset
	for ; i < header.LabelCount; i++ {
		result = append(result, strings.Trim(string(bytes[index:index+16]), "x00"))

		index += 16
	}

	return result
}

func extractFieldDataBlockFromBytes(bytes []byte, header Header) []byte {
	return bytes[header.FieldDataOffset : header.FieldDataOffset+header.FieldDataCount]
}

func extractFieldIndicesArrayFromBytes(bytes []byte, header Header) []uint32 {
	var result = []uint32{}
	var i = uint32(0)
	var index = header.FieldIndicesOffset
	for ; i*4 < header.FieldIndicesCount; i++ {
		result = append(result, fileReader.BytesToUint32LE(bytes[index:index+4]))

		index += 4
	}

	return result
}

func extractListIndicesArrayFromBytes(bytes []byte, header Header) []ListIndicesElement {
	var result = []ListIndicesElement{}
	var bytesRead = uint32(0)

	var index = header.ListIndicesOffset
	for bytesRead < header.ListIndicesCount {
		var element = []uint32{}
		var size = fileReader.BytesToUint32LE(bytes[index : index+4])
		index += 4
		var i = uint32(0)
		for ; i < size; i++ {
			element = append(element, fileReader.BytesToUint32LE(bytes[index:index+4]))
			index += 4
		}

		result = append(result, ListIndicesElement(element))
		bytesRead += (size * 4) + 4 // size + the length of the size bytes
	}

	return result
}
