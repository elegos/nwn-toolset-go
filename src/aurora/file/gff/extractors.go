package gff

import (
	"aurora/tools/fileReader"
	"errors"
	"fmt"
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
func extractStructArrayFromBytes(bytes []byte, header Header) []Struct {
	var result = []Struct{}

	var i = uint32(0)
	var index = header.StructOffset
	for ; i < header.StructCount; i++ {
		result = append(result, Struct{
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
			Type:             FieldType(fileReader.BytesToUint32LE(bytes[index : index+4])),
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
		result = append(result, strings.Trim(string(bytes[index:index+16]), "\x00"))

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

func extractListIndicesArrayFromBytes(bytes []byte, header Header) [][]uint32 {
	var result = [][]uint32{}
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

		result = append(result, element)
		bytesRead += (size * 4) + 4 // size + the length of the size bytes
	}

	return result
}

// ExtractListStructArrayIndexes extract the indexes list from the ListIndicesArray
func ExtractListStructArrayIndexes(byteOffset uint32, file GFF) ([]uint32, error) {
	var offset = uint32(0)
	for i := 0; i < len(file.ListIndicesArray); i++ {
		var listIndicesArray = file.ListIndicesArray[i]

		if byteOffset == offset {
			return file.ListIndicesArray[i], nil
		}

		// list size (4 bytes) + list indexes (* 4 bytes)
		offset += 4 + uint32(len(listIndicesArray))*4
	}

	return []uint32{}, fmt.Errorf("Byte offset %d not found", byteOffset)
}

// ExtractStructFields extract the fields of the given struct array element
func ExtractStructFields(element Struct, file GFF) ([]Field, error) {
	var result = []Field{}

	if element.FieldCount == 1 {
		return result, errors.New("The struct has only one field, read it directly")
	}

	// The offset is in bytes, but we're accessing via slice index.
	// Being all the field indices uint32 (4 bytes long), we have to divide by 4.
	var startOffset = file.FieldIndicesArray[element.DataOrDataOffset/4]
	var endOffset = startOffset + element.FieldCount

	for index := startOffset; index < endOffset; index++ {
		fieldArrayElement := file.FieldArray[index]

		result = append(result, Field{
			Type:  fieldArrayElement.Type,
			Label: file.LabelArray[fieldArrayElement.LabelIndex],
			Data:  extractFieldData(fieldArrayElement.Type, fieldArrayElement.DataOrDataOffset, file),
		})
	}

	return result, nil
}
