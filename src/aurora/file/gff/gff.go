package gff

import (
	"aurora/tools"
	"io/ioutil"
)

// Header structure of the GFF file format's header
type Header struct {
	FileType           string // 4 chars, 4 bytes
	FileVersion        string // 4 chars, 4 bytes
	StructOffset       uint32 // Offset of Struct array as bytes from the beginning of the file
	StructCount        uint32 // Number of elements in Struct array
	FieldOffset        uint32 // Offset of Field array as bytes from the beginning of the file
	FieldCount         uint32 // Number of elements in Field array
	LabelOffset        uint32 // Offset of Label array as bytes from the beginning of the file
	LabelCount         uint32 // Number of elements in Label array
	FieldDataOffset    uint32 // Offset of Field Data as bytes from the beginning of the file
	FieldDataCount     uint32 // Number of bytes in Field Data block
	FieldIndicesOffset uint32 // Offset of Field Indices array as bytes from the beginning og the file
	FieldIndicesCount  uint32 // Number of bytes in Field Indices array
	ListIndicesOffset  uint32 // Offset of List Indices array as bytes from the beginning of the file
	ListIndicesCount   uint32 // Number of bytes in List Indices array
}

// StructArrayElement the element of the StructArray
type StructArrayElement struct {
	// Type programmer-defined integer ID
	Type uint32 // Programmer-defined integer ID
	// DataOrDataOffset if FieldCount = 1, this is an index into the Field Array.
	// If FieldCount > 1, this is a byte offset into the FieldIndices array, where there is
	// an array of DWORDs having a number of elements equal to FieldCount. Each one of these
	// DWORDs is an index into the Field Array
	DataOrDataOffset uint32
	FieldCount       uint32 // Number of fields in this Struct
}

// FieldArrayElement the element of FieldArray
type FieldArrayElement struct {
	Type             uint32
	LabelIndex       uint32 // Index into the Label Array
	DataOrDataOffset uint32 // If Type is a simple data type, then this is the value
	// actual of the field. If Type is a complex data type,
	// then this is a byte offset into the Field Data block.
}

// GFF Generic File Format
type GFF struct {
	Header      Header
	StructArray []StructArrayElement
	FieldArray  []FieldArrayElement
}

// FromBytes read the bytes and return a GFF struct
func FromBytes(bytes []byte) GFF {
	var result = GFF{}

	result.Header = extractHeaderFromBytes(bytes)
	result.StructArray = extractStructArrayFromBytes(bytes[result.Header.StructOffset:], result.Header)
	result.FieldArray = extractFieldArrayFromBytes(bytes[result.Header.FieldOffset:], result.Header)

	return result
}

// FromFile read the file and return a GFF struct
func FromFile(fileName string) GFF {
	bytes, err := ioutil.ReadFile(fileName)
	tools.EasyPanic(err)

	return FromBytes(bytes)
}
