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

func TestExpectFloat32(t *testing.T) {
	var tStruct = testing.T{}
	var field = "field"
	var expected = float32(0.987)
	var incorrectValue = float32(9.546)

	test.ExpectFloat32(&tStruct, field, expected, expected)

	if tStruct.Failed() {
		t.Error("Expected no fail, failed")
	}

	test.ExpectFloat32(&tStruct, field, expected, incorrectValue)

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

func TestExpectInt32(t *testing.T) {
	var tStruct = testing.T{}
	var field = "field"
	var expected = int32(15)
	var incorrectValue = int32(27)

	test.ExpectInt32(&tStruct, field, expected, expected)

	if tStruct.Failed() {
		t.Error("Expected no fail, failed")
	}

	test.ExpectInt32(&tStruct, field, expected, incorrectValue)

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

func TestExpectUint8(t *testing.T) {
	var tStruct = testing.T{}
	var field = "field"
	var expected = uint8(66)
	var incorrectValue = uint8(109)

	test.ExpectUint8(&tStruct, field, expected, expected)

	if tStruct.Failed() {
		t.Error("Expected no fail, failed")
	}

	test.ExpectUint8(&tStruct, field, expected, incorrectValue)

	if !tStruct.Failed() {
		t.Error("Expected fail, not failed")
	}
}

func TestExpectInt8(t *testing.T) {
	var tStruct = testing.T{}
	var field = "field"
	var expected = int8(15)
	var incorrectValue = int8(27)

	test.ExpectInt8(&tStruct, field, expected, expected)

	if tStruct.Failed() {
		t.Error("Expected no fail, failed")
	}

	test.ExpectInt8(&tStruct, field, expected, incorrectValue)

	if !tStruct.Failed() {
		t.Error("Expected fail, not failed")
	}
}
