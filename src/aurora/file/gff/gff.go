package gff

import (
	"aurora/file"
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

// Struct the element of the StructArray
type Struct struct {
	// Type programmer-defined integer ID
	Type uint32 // Programmer-defined integer ID
	// DataOrDataOffset if FieldCount = 1, this is an index into the Field Array.
	// If FieldCount > 1, this is a byte offset into the FieldIndices array, where there is
	// an array of DWORDs having a number of elements equal to FieldCount. Each one of these
	// DWORDs is an index into the Field Array
	DataOrDataOffset uint32
	FieldCount       uint32 // Number of fields in this Struct
}

// ListIndex the list of indexes of a list
type ListIndex struct {
	Indexes []uint32
}

// FieldArrayElement the element of FieldArray
type FieldArrayElement struct {
	Type             FieldType
	LabelIndex       uint32 // Index into the Label Array
	DataOrDataOffset uint32 // If Type is a simple data type, then this is the value
	// actual of the field. If Type is a complex data type,
	// then this is a byte offset into the Field Data block.
}

// CExoLocStringSubString sub-element of CExoLocString
type CExoLocStringSubString struct {
	LanguageID file.Language // The language's id
	Masculine  bool          // If false, feminine
	String     string
}

// CExoLocString localised string
type CExoLocString struct {
	StringRef  uint32 // Index into the user's dialog.tlk file. If 0xffffffff, it doesn't reference the tlk file.
	SubStrings []CExoLocStringSubString
}

// FieldDataStruct the field's data structure to keep any possible value. Only one is valorised.
type FieldDataStruct struct {
	ByteValue          int8
	CharValue          string
	WordValue          uint16
	ShortValue         int16
	DwordValue         uint32
	IntValue           int32
	Dword64Value       uint64
	Int64Value         int64
	FloatValue         float32
	DoubleValue        float64
	CExoStringValue    string
	CResRefValue       string
	CExoLocStringValue CExoLocString
	VoidValue          []byte
	StructIndexValue   uint32
	ListOffsetValue    uint32 // Index in the Struct Array Data
}

// Field a struct's field
type Field struct {
	Type  FieldType
	Label string
	Data  FieldDataStruct
}

// ListIndicesElement the element of ListIndicesArray
type ListIndicesElement []uint32

// GFF Generic File Format
type GFF struct {
	Header            Header
	StructArray       []Struct
	FieldArray        []FieldArrayElement
	LabelArray        []string
	FieldDataBlock    []byte
	FieldIndicesArray []uint32
	ListIndicesArray  [][]uint32
}

// FromBytes read the bytes and return a GFF struct
func FromBytes(bytes []byte) GFF {
	var result = GFF{}

	result.Header = extractHeaderFromBytes(bytes)
	result.StructArray = extractStructArrayFromBytes(bytes, result.Header)
	result.FieldArray = extractFieldArrayFromBytes(bytes, result.Header)
	result.LabelArray = extractLabelArrayFromBytes(bytes, result.Header)
	result.FieldDataBlock = extractFieldDataBlockFromBytes(bytes, result.Header)
	result.FieldIndicesArray = extractFieldIndicesArrayFromBytes(bytes, result.Header)
	result.ListIndicesArray = extractListIndicesArrayFromBytes(bytes, result.Header)

	return result
}

// FromFile read the file and return a GFF struct
func FromFile(fileName string) (GFF, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return GFF{}, err
	}

	return FromBytes(bytes), nil
}
