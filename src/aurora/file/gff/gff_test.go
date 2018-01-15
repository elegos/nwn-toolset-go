package gff_test

import (
	"aurora/file/gff"
	"aurora/tools/test"
	"testing"
)

func TestFromFile_willFailErrorReadingUnexistingFile(t *testing.T) {
	t.Parallel()

	var _, err = gff.FromFile("e9ed7a78-18f6-4a00-9eba-b92c1152b4c6.ext")

	if err == nil {
		t.Error("Expecting file reading error")
	}
}

func TestFromFile_willExtractAnAreFile(t *testing.T) {
	t.Parallel()

	var gffFile, err = gff.FromFile("../are/test/barrowsinterior8.are")

	if err != nil {
		t.Errorf("Expected no error, found: %v", err)
	}

	// Header checks
	test.ExpectString(t, "Header.FileType", "ARE ", gffFile.Header.FileType)
	test.ExpectString(t, "Header.FileVersion", "V3.2", gffFile.Header.FileVersion)
	test.ExpectUint32(t, "Header.StructOffset", 56, gffFile.Header.StructOffset)
	test.ExpectUint32(t, "Header.StructCount", 65, gffFile.Header.StructCount)
	test.ExpectUint32(t, "Header.FieldOffset", 836, gffFile.Header.FieldOffset)
	test.ExpectUint32(t, "Header.FieldCount", 682, gffFile.Header.FieldCount)
	test.ExpectUint32(t, "Header.LabelOffset", 9020, gffFile.Header.LabelOffset)
	test.ExpectUint32(t, "Header.LabelCount", 52, gffFile.Header.LabelCount)
	test.ExpectUint32(t, "Header.FieldDataOffset", 9852, gffFile.Header.FieldDataOffset)
	test.ExpectUint32(t, "Header.FieldDataCount", 115, gffFile.Header.FieldDataCount)
	test.ExpectUint32(t, "Header.FieldIndicesOffset", 9967, gffFile.Header.FieldIndicesOffset)
	test.ExpectUint32(t, "Header.FieldIndicesCount", 2728, gffFile.Header.FieldIndicesCount)
	test.ExpectUint32(t, "Header.ListIndicesOffset", 12695, gffFile.Header.ListIndicesOffset)
	test.ExpectUint32(t, "Header.ListIndicesCount", 264, gffFile.Header.ListIndicesCount)

	// Data checks
	test.ExpectUint32(t, "StructArray (len)", gffFile.Header.StructCount, uint32(len(gffFile.StructArray)))
	test.ExpectUint32(t, "FieldArray (len)", gffFile.Header.FieldCount, uint32(len(gffFile.FieldArray)))
	test.ExpectUint32(t, "LabelArray (len)", gffFile.Header.LabelCount, uint32(len(gffFile.LabelArray)))
	test.ExpectUint32(t, "FieldDataBlock (len)", gffFile.Header.FieldDataCount, uint32(len(gffFile.FieldDataBlock)))
	test.ExpectUint32(t, "FieldIndicesArray (len)", gffFile.Header.FieldIndicesCount/4, uint32(len(gffFile.FieldIndicesArray)))
	test.ExpectUint32(t, "ListIndicesArray (len)", gffFile.Header.ListIndicesCount/132, uint32(len(gffFile.ListIndicesArray)))
}
