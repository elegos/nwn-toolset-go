package gff

import (
	"aurora/tools"
	"fmt"
)

// GetField get the required field
func GetField(name string, fields []Field) (Field, error) {
	for i := 0; i < len(fields); i++ {
		if fields[i].Label == name {
			return fields[i], nil
		}
	}

	return Field{}, fmt.Errorf("No field with label '%s' was found", name)
}

// GetFieldInt32Value return the int32 value
func GetFieldInt32Value(name string, fields []Field, errorBag *tools.ErrorBag) int32 {
	if errorBag.Error != nil {
		return 0
	}

	field, err := GetField(name, fields)

	if err != nil {
		errorBag.Error = err

		return 0
	}

	if field.Type != FieldTypeInt {
		errorBag.Error = fmt.Errorf("Expected INT, found %s", FieldTypeLookup[field.Type])

		return 0
	}

	return field.Data.IntValue
}

// GetFieldUint32Value return the uint32 value
func GetFieldUint32Value(name string, fields []Field, errorBag *tools.ErrorBag) uint32 {
	if errorBag.Error != nil {
		return 0
	}

	field, err := GetField(name, fields)

	if err != nil {
		errorBag.Error = err

		return 0
	}

	if field.Type != FieldTypeDword {
		errorBag.Error = fmt.Errorf("Expected DWORD, found %s", FieldTypeLookup[field.Type])

		return 0
	}

	return field.Data.DwordValue
}

// GetFieldListByteOffsetValue return the byte offset from the beginning of the List Indices Array
func GetFieldListByteOffsetValue(name string, fields []Field, errorBag *tools.ErrorBag) uint32 {
	if errorBag.Error != nil {
		return 0
	}

	field, err := GetField(name, fields)

	if err != nil {
		errorBag.Error = err

		return 0
	}

	if field.Type != FieldTypeList {
		errorBag.Error = fmt.Errorf("Expected List, found %s", FieldTypeLookup[field.Type])

		return 0
	}

	return field.Data.ListOffsetValue
}

// GetFieldUint16Value return the uint16 value
func GetFieldUint16Value(name string, fields []Field, errorBag *tools.ErrorBag) uint16 {
	if errorBag.Error != nil {
		return 0
	}

	field, err := GetField(name, fields)

	if err != nil {
		errorBag.Error = err

		return 0
	}

	if field.Type != FieldTypeWord {
		errorBag.Error = fmt.Errorf("Expected WORD, found %s", FieldTypeLookup[field.Type])

		return 0
	}

	return field.Data.WordValue
}

// GetFieldByteValue return the int8 value
func GetFieldByteValue(name string, fields []Field, errorBag *tools.ErrorBag) int8 {
	if errorBag.Error != nil {
		return 0
	}

	field, err := GetField(name, fields)

	if err != nil {
		errorBag.Error = err

		return 0
	}

	if field.Type != FieldTypeByte {
		errorBag.Error = fmt.Errorf("Expected BYTE, found %s", FieldTypeLookup[field.Type])

		return 0
	}

	return field.Data.ByteValue
}

// GetFieldFloatValue get the field's float value
func GetFieldFloatValue(name string, fields []Field, errorBag *tools.ErrorBag) float32 {
	if errorBag.Error != nil {
		return 0
	}

	field, err := GetField(name, fields)

	if err != nil {
		errorBag.Error = err

		return 0
	}

	if field.Type != FieldTypeFloat {
		errorBag.Error = fmt.Errorf("Expected FLOAT, found %s", FieldTypeLookup[field.Type])

		return 0
	}

	return field.Data.FloatValue
}

// GetFieldCExoStringValue get the field's CExoString value
func GetFieldCExoStringValue(name string, fields []Field, errorBag *tools.ErrorBag) string {
	if errorBag.Error != nil {
		return ""
	}

	field, err := GetField(name, fields)

	if err != nil {
		errorBag.Error = err

		return ""
	}

	if field.Type != FieldTypeCExoString {
		errorBag.Error = fmt.Errorf("Expected CExoString, found %s", FieldTypeLookup[field.Type])

		return ""
	}

	return field.Data.CExoStringValue
}

// GetFieldCResRefValue get the fields's CResRef value
func GetFieldCResRefValue(name string, fields []Field, errorBag *tools.ErrorBag) string {
	if errorBag.Error != nil {
		return ""
	}

	field, err := GetField(name, fields)

	if err != nil {
		errorBag.Error = err

		return ""
	}

	if field.Type != FieldTypeCResRef {
		errorBag.Error = fmt.Errorf("Expected ResRef, found %s", FieldTypeLookup[field.Type])

		return ""
	}

	return field.Data.CResRefValue
}

// GetFieldCExoLocStringValue get the field's CExoLocString value(s)
func GetFieldCExoLocStringValue(name string, fields []Field, errorBag *tools.ErrorBag) CExoLocString {
	if errorBag.Error != nil {
		return CExoLocString{}
	}

	field, err := GetField(name, fields)

	if err != nil {
		errorBag.Error = err

		return CExoLocString{}
	}

	if field.Type != FieldTypeCExoLocString {
		errorBag.Error = fmt.Errorf("Expected ResRef, found %s", FieldTypeLookup[field.Type])

		return CExoLocString{}
	}

	return field.Data.CExoLocStringValue
}
