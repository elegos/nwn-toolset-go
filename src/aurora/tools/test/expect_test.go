package test_test

import (
	"aurora/tools/test"
	"testing"
)

func TestExpectString(t *testing.T) {
	var tStruct = testing.T{}
	var field = "field"
	var expected = "test string"
	var incorrectValue = "incorrect text"

	test.ExpectString(&tStruct, field, expected, expected)

	if tStruct.Failed() {
		t.Error("Expected no fail, failed")
	}

	test.ExpectString(&tStruct, field, expected, incorrectValue)

	if !tStruct.Failed() {
		t.Error("Expected fail, not failed")
	}
}

func TestExpectUint32(t *testing.T) {
	var tStruct = testing.T{}
	var field = "field"
	var expected = uint32(15)
	var incorrectValue = uint32(27)

	test.ExpectUint32(&tStruct, field, expected, expected)

	if tStruct.Failed() {
		t.Error("Expected no fail, failed")
	}

	test.ExpectUint32(&tStruct, field, expected, incorrectValue)

	if !tStruct.Failed() {
		t.Error("Expected fail, not failed")
	}
}

func TestExpectUint16(t *testing.T) {
	var tStruct = testing.T{}
	var field = "field"
	var expected = uint16(15)
	var incorrectValue = uint16(27)

	test.ExpectUint16(&tStruct, field, expected, expected)

	if tStruct.Failed() {
		t.Error("Expected no fail, failed")
	}

	test.ExpectUint16(&tStruct, field, expected, incorrectValue)

	if !tStruct.Failed() {
		t.Error("Expected fail, not failed")
	}
}
