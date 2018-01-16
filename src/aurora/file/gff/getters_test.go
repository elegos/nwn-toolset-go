package gff_test

import (
	"aurora/file"
	"aurora/file/gff"
	"aurora/tools"
	"aurora/tools/test"
	"testing"
)

var cExoStringField = gff.Field{
	Label: "String field",
	Type:  gff.FieldTypeCExoString,
	Data: gff.FieldDataStruct{
		CExoStringValue: "Test value",
	},
}

var cResRefField = gff.Field{
	Label: "CResRef field",
	Type:  gff.FieldTypeCResRef,
	Data: gff.FieldDataStruct{
		CResRefValue: "ResRef",
	},
}

var byteField = gff.Field{
	Label: "Byte field",
	Type:  gff.FieldTypeByte,
	Data: gff.FieldDataStruct{
		ByteValue: int8(56),
	},
}

var uint16Field = gff.Field{
	Label: "Uint16 (WORD) field",
	Type:  gff.FieldTypeWord,
	Data: gff.FieldDataStruct{
		WordValue: uint16(14),
	},
}

var int32Field = gff.Field{
	Label: "Int32 (INT) field",
	Type:  gff.FieldTypeInt,
	Data: gff.FieldDataStruct{
		ByteValue: int8(56),
	},
}

var uint32Field = gff.Field{
	Label: "Uint32 (DWORD) field",
	Type:  gff.FieldTypeDword,
	Data: gff.FieldDataStruct{
		DwordValue: uint32(768),
	},
}

var float32Field = gff.Field{
	Label: "Float32 (FLOAT) field",
	Type:  gff.FieldTypeFloat,
	Data: gff.FieldDataStruct{
		FloatValue: float32(0.59),
	},
}

var listField = gff.Field{
	Label: "List (List) field",
	Type:  gff.FieldTypeList,
	Data: gff.FieldDataStruct{
		ListOffsetValue: uint32(559),
	},
}

var cExoLocStringField = gff.Field{
	Label: "cExoLocString (cExoLocString) field",
	Type:  gff.FieldTypeCExoLocString,
	Data: gff.FieldDataStruct{
		CExoLocStringValue: gff.CExoLocString{
			StringRef: 0xffffffff,
			SubStrings: []gff.CExoLocStringSubString{
				gff.CExoLocStringSubString{
					LanguageID: file.LangEnglish,
					Masculine:  true,
					String:     "Localized test string",
				},
			},
		},
	},
}

func getFields() []gff.Field {
	var fields []gff.Field

	fields = append(fields, byteField)
	fields = append(fields, cExoStringField)
	fields = append(fields, cResRefField)
	fields = append(fields, float32Field)
	fields = append(fields, int32Field)
	fields = append(fields, uint16Field)
	fields = append(fields, uint32Field)
	fields = append(fields, listField)
	fields = append(fields, cExoLocStringField)

	return fields[:]
}

func TestGetField(t *testing.T) {
	t.Parallel()

	var fields = getFields()

	fieldResult, err := gff.GetField(byteField.Label, fields[:])
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	test.ExpectString(t, "Search for the byte field", byteField.Label, fieldResult.Label)

	_, err = gff.GetField("unknown field", fields[:])
	if err == nil {
		t.Error("Unexpected missing error")
	}
}

func TestGetFieldInt32Value(t *testing.T) {
	t.Parallel()

	var fields = getFields()
	var errorBag = tools.ErrorBag{}

	result := gff.GetFieldInt32Value(int32Field.Label, fields, &errorBag)
	if errorBag.Error != nil {
		t.Errorf("Unexpected error: %v", errorBag.Error)
	}

	test.ExpectInt32(t, "Int32 value", int32Field.Data.IntValue, result)

	_ = gff.GetFieldInt32Value(cExoStringField.Label, fields, &errorBag)
	if errorBag.Error == nil {
		t.Error("Expected error, found none")
	}

	result = gff.GetFieldInt32Value(int32Field.Label, fields, &errorBag)
	test.ExpectInt32(t, "Int32 default value (error)", 0, result)
}

func TestGetFieldUint32Value(t *testing.T) {
	t.Parallel()

	var fields = getFields()
	var errorBag = tools.ErrorBag{}

	result := gff.GetFieldUint32Value(uint32Field.Label, fields, &errorBag)
	if errorBag.Error != nil {
		t.Errorf("Unexpected error: %v", errorBag.Error)
	}

	test.ExpectUint32(t, "Uint32 value", uint32Field.Data.DwordValue, result)

	_ = gff.GetFieldUint32Value(cExoStringField.Label, fields, &errorBag)
	if errorBag.Error == nil {
		t.Error("Expected error, found none")
	}

	result = gff.GetFieldUint32Value(uint32Field.Label, fields, &errorBag)
	test.ExpectUint32(t, "Uint32 default value (error)", 0, result)
}

func TestGetFieldUint16Value(t *testing.T) {
	t.Parallel()

	var fields = getFields()
	var errorBag = tools.ErrorBag{}

	result := gff.GetFieldUint16Value(uint16Field.Label, fields, &errorBag)
	if errorBag.Error != nil {
		t.Errorf("Unexpected error: %v", errorBag.Error)
	}

	test.ExpectUint16(t, "Uint16 value", uint16Field.Data.WordValue, result)

	_ = gff.GetFieldUint16Value(cExoStringField.Label, fields, &errorBag)
	if errorBag.Error == nil {
		t.Error("Expected error, found none")
	}

	result = gff.GetFieldUint16Value(uint32Field.Label, fields, &errorBag)
	test.ExpectUint16(t, "Uint16 default value (error)", 0, result)
}

func TestGetFieldByteValue(t *testing.T) {
	t.Parallel()

	var fields = getFields()
	var errorBag = tools.ErrorBag{}

	result := gff.GetFieldByteValue(byteField.Label, fields, &errorBag)
	if errorBag.Error != nil {
		t.Errorf("Unexpected error: %v", errorBag.Error)
	}

	test.ExpectInt8(t, "Byte value (Uint8)", byteField.Data.ByteValue, result)

	_ = gff.GetFieldByteValue(cExoStringField.Label, fields, &errorBag)
	if errorBag.Error == nil {
		t.Error("Expected error, found none")
	}

	result = gff.GetFieldByteValue(byteField.Label, fields, &errorBag)
	test.ExpectInt8(t, "Byte default value (error)", 0, result)
}

func TestGetFieldFloatValue(t *testing.T) {
	t.Parallel()

	var fields = getFields()
	var errorBag = tools.ErrorBag{}

	result := gff.GetFieldFloatValue(float32Field.Label, fields, &errorBag)
	if errorBag.Error != nil {
		t.Errorf("Unexpected error: %v", errorBag.Error)
	}

	test.ExpectFloat32(t, "Float value (float32)", float32Field.Data.FloatValue, result)

	_ = gff.GetFieldFloatValue(cExoStringField.Label, fields, &errorBag)
	if errorBag.Error == nil {
		t.Error("Expected error, found none")
	}

	result = gff.GetFieldFloatValue(float32Field.Label, fields, &errorBag)
	test.ExpectFloat32(t, "Float32 default value (error)", 0, result)
}

func TestGetFieldCExoStringValue(t *testing.T) {
	t.Parallel()

	var fields = getFields()
	var errorBag = tools.ErrorBag{}

	result := gff.GetFieldCExoStringValue(cExoStringField.Label, fields, &errorBag)
	if errorBag.Error != nil {
		t.Errorf("Unexpected error: %v", errorBag.Error)
	}

	test.ExpectString(t, "CExoString value (string)", cExoStringField.Data.CExoStringValue, result)

	_ = gff.GetFieldCExoStringValue(float32Field.Label, fields, &errorBag)
	if errorBag.Error == nil {
		t.Error("Expected error, found none")
	}

	result = gff.GetFieldCExoStringValue(cExoStringField.Label, fields, &errorBag)
	test.ExpectString(t, "CExoString default value (error)", "", result)
}

func TestGetFieldCResRefValue(t *testing.T) {
	t.Parallel()

	var fields = getFields()
	var errorBag = tools.ErrorBag{}

	result := gff.GetFieldCResRefValue(cResRefField.Label, fields, &errorBag)
	if errorBag.Error != nil {
		t.Errorf("Unexpected error: %v", errorBag.Error)
	}

	test.ExpectString(t, "CExoString value (string)", cResRefField.Data.CResRefValue, result)

	_ = gff.GetFieldCResRefValue(float32Field.Label, fields, &errorBag)
	if errorBag.Error == nil {
		t.Error("Expected error, found none")
	}

	result = gff.GetFieldCResRefValue(cResRefField.Label, fields, &errorBag)
	test.ExpectString(t, "CResRef default value (error)", "", result)
}

func TestGetFieldListByteOffsetValue(t *testing.T) {
	t.Parallel()

	var fields = getFields()
	var errorBag = tools.ErrorBag{}

	result := gff.GetFieldListByteOffsetValue(listField.Label, fields, &errorBag)
	if errorBag.Error != nil {
		t.Errorf("Unexpected error: %v", errorBag.Error)
	}

	test.ExpectUint32(t, "List value byte offset (uint32)", listField.Data.ListOffsetValue, result)

	_ = gff.GetFieldListByteOffsetValue(float32Field.Label, fields, &errorBag)
	if errorBag.Error == nil {
		t.Error("Expected error, found none")
	}

	result = gff.GetFieldListByteOffsetValue(listField.Label, fields, &errorBag)
	test.ExpectUint32(t, "List byte offset default value (error)", 0, result)
}

func TestGetFieldCExoLocStringValue(t *testing.T) {
	t.Parallel()

	var fields = getFields()
	var errorBag = tools.ErrorBag{}

	result := gff.GetFieldCExoLocStringValue(cExoLocStringField.Label, fields, &errorBag)
	if errorBag.Error != nil {
		t.Errorf("Unexpected error: %v", errorBag.Error)
	}

	test.ExpectUint32(
		t, "cExoLocString resref value (uint32)",
		cExoLocStringField.Data.CExoLocStringValue.StringRef,
		result.StringRef,
	)

	test.ExpectString(
		t, "cExoLocString substring value (string)",
		cExoLocStringField.Data.CExoLocStringValue.SubStrings[0].String,
		result.SubStrings[0].String,
	)

	_ = gff.GetFieldCExoLocStringValue(float32Field.Label, fields, &errorBag)
	if errorBag.Error == nil {
		t.Error("Expected error, found none")
	}

	result = gff.GetFieldCExoLocStringValue(listField.Label, fields, &errorBag)
	test.ExpectUint32(t, "cExoLocString StringRef default value (error)", 0, result.StringRef)
}
