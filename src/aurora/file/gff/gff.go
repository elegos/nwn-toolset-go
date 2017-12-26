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

// GFF Generic File Format
type GFF struct {
	Header Header
}

// FromBytes read the bytes and return a GFF struct
func FromBytes(bytes []byte) GFF {
	var result = GFF{}

	result.Header = extractHeaderFromBytes(bytes)

	return result
}

// FromFile read the file and return a GFF struct
func FromFile(fileName string) GFF {
	bytes, err := ioutil.ReadFile(fileName)
	tools.EasyPanic(err)

	return FromBytes(bytes)
}
