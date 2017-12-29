package test

import "testing"

// ExpectString simple expect / string comparison
func ExpectString(t *testing.T, field string, expected string, found string) {
	if expected != found {
		t.Errorf("%s [%s], found '%s'", field, expected, found)
	}
}

// ExpectUint32 simple expect / unsigned 32 bits comparison
func ExpectUint32(t *testing.T, field string, expected uint32, found uint32) {
	if expected != found {
		t.Errorf("%s [%d], found '%d'", field, expected, found)
	}
}

// ExpectUint16 simple expect / unsigned 16 bits comparison
func ExpectUint16(t *testing.T, field string, expected uint16, found uint16) {
	if expected != found {
		t.Errorf("%s [%d], found '%d'", field, expected, found)
	}
}
