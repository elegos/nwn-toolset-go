package gff_test

import (
	"aurora/file/gff"
	"aurora/tools/test"
	"testing"
)

func TestExtractListStructArrayIndexes_willFailOutOfRange(t *testing.T) {
	t.Parallel()

	gffFile, _ := gff.FromFile("../are/test/barrowsinterior8.are")
	_, err := gff.ExtractListStructArrayIndexes(1024, gffFile)

	if err == nil {
		t.Error("Expected error, none found")
	}
}

func TestExtractListStructArrayIndexes(t *testing.T) {
	t.Parallel()

	gffFile, _ := gff.FromFile("../are/test/barrowsinterior8.are")
	// in ARE files, these are the tiles
	indexes, err := gff.ExtractListStructArrayIndexes(4, gffFile)

	if err != nil {
		t.Errorf("Expected no error, found: %v", err)
	}

	test.ExpectUint32(t, "Extracted list struct array length", uint32(64), uint32(len(indexes)))
	// Tiles are an ordered list
	var i = uint32(0)
	for ; i < uint32(len(indexes)); i++ {
		test.ExpectUint32(t, "Extractes list indexes", i+1, indexes[i])
	}
}

func TestExtractStructFields_willFailStructSimpleValue(t *testing.T) {
	t.Parallel()

	gffFile, _ := gff.FromFile("../are/test/barrowsinterior8.are")
	var element = gff.Struct{
		DataOrDataOffset: uint32(123),
		FieldCount:       uint32(1),
		Type:             uint32(0),
	}

	_, err := gff.ExtractStructFields(element, gffFile)

	if err == nil {
		t.Error("Expected error, found none")
	}
}

func TestExtractStructFields(t *testing.T) {
	t.Parallel()

	gffFile, _ := gff.FromFile("../are/test/barrowsinterior8.are")
	// Example tile field
	var element = gff.Struct{
		DataOrDataOffset: uint32(2688),
		FieldCount:       uint32(10),
		Type:             uint32(1),
	}

	fields, err := gff.ExtractStructFields(element, gffFile)

	if err != nil {
		t.Errorf("Expected no error, found: %v", err)
	}

	test.ExpectUint32(t, "Extracted fields length", element.FieldCount, uint32(len(fields)))

	var expected = map[string]gff.FieldType{
		"Tile_ID":          gff.FieldTypeInt,
		"Tile_Orientation": gff.FieldTypeInt,
		"Tile_Height":      gff.FieldTypeInt,
		"Tile_MainLight1":  gff.FieldTypeByte,
		"Tile_MainLight2":  gff.FieldTypeByte,
		"Tile_SrcLight1":   gff.FieldTypeByte,
		"Tile_SrcLight2":   gff.FieldTypeByte,
		"Tile_AnimLoop1":   gff.FieldTypeByte,
		"Tile_AnimLoop2":   gff.FieldTypeByte,
		"Tile_AnimLoop3":   gff.FieldTypeByte,
	}

	for i := 0; i < len(fields); i++ {
		test.ExpectUint32(t, fields[i].Label, uint32(expected[fields[i].Label]), uint32(fields[i].Type))
	}
}
