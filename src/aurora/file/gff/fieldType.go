package gff

// FieldType the gff field's type
type FieldType uint32

const (
	// FieldTypeByte type BYTE (int8)
	FieldTypeByte FieldType = 0
	// FieldTypeChar type CHAR (1 byte string)
	FieldTypeChar FieldType = 1
	// FieldTypeWord type WORD (uint16)
	FieldTypeWord FieldType = 2
	// FieldTypeShort type SHORT (int16)
	FieldTypeShort FieldType = 3
	// FieldTypeDword type DWORD (uint32)
	FieldTypeDword FieldType = 4
	// FieldTypeInt type INT (int32)
	FieldTypeInt FieldType = 5
	// FieldTypeDword64 type DWORD64 (uint64)
	FieldTypeDword64 FieldType = 6
	// FieldTypeInt64 type INT64 (int64)
	FieldTypeInt64 FieldType = 7
	// FieldTypeFloat type FLOAT (float32)
	FieldTypeFloat FieldType = 8
	// FieldTypeDouble type DOUBLE (float64)
	FieldTypeDouble FieldType = 9
	// FieldTypeCExoString type CExpString (string)
	FieldTypeCExoString FieldType = 10
	// FieldTypeCResRef type CResRef (string)
	FieldTypeCResRef FieldType = 11
	// FieldTypeCExoLocString type CExoLocString (struct)
	FieldTypeCExoLocString FieldType = 12
	// FieldTypeVoid type VOID (bytes)
	FieldTypeVoid FieldType = 13
	// FieldTypeStruct type Struct (struct)
	FieldTypeStruct FieldType = 14
	// FieldTypeList type List (array of homogeneous data)
	FieldTypeList FieldType = 15
)

// FieldTypeLookup from constant to name
var FieldTypeLookup = map[FieldType]string{
	FieldTypeByte:          "BYTE",
	FieldTypeChar:          "CHAR",
	FieldTypeWord:          "WORD",
	FieldTypeShort:         "SHORT",
	FieldTypeDword:         "DWORD",
	FieldTypeInt:           "INT",
	FieldTypeDword64:       "DWORD64",
	FieldTypeInt64:         "INT64",
	FieldTypeFloat:         "FLOAT",
	FieldTypeDouble:        "DOUBLE",
	FieldTypeCExoString:    "CExoString",
	FieldTypeCResRef:       "CResRef",
	FieldTypeCExoLocString: "CExoLocString",
	FieldTypeVoid:          "VOID",
	FieldTypeStruct:        "Struct",
	FieldTypeList:          "List",
}
