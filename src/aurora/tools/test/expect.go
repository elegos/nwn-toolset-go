package test

import "testing"

// ExpectString simple expect / string comparison
func ExpectString(t *testing.T, field string, expected string, found string) {
	if expected != found {
		t.Errorf("%s [%s], found '%s'", field, expected, found)
	}
}

// ExpectFloat32 simple expect / float32 comparison
func ExpectFloat32(t *testing.T, field string, expected float32, found float32) {
	if expected != found {
		t.Errorf("%s [%f], found '%f'", field, expected, found)
	}
}

// ExpectUint32 simple expect / unsigned 32 bits comparison
func ExpectUint32(t *testing.T, field string, expected uint32, found uint32) {
	if expected != found {
		t.Errorf("%s [%d], found '%d'", field, expected, found)
	}
}

// ExpectInt32 simple expect / signed 32 bits comparison
func ExpectInt32(t *testing.T, field string, expected int32, found int32) {
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

// ExpectUint8 simple expect / unsigned 8 bits comparison
func ExpectUint8(t *testing.T, field string, expected uint8, found uint8) {
	if expected != found {
		t.Errorf("%s [%d], found '%d'", field, expected, found)
	}
}

// ExpectInt8 simple expect / signed 8 bits comparison
func ExpectInt8(t *testing.T, field string, expected int8, found int8) {
	if expected != found {
		t.Errorf("%s [%d], found '%d'", field, expected, found)
	}
}

// ExpectBool simple expect bool comparison
func ExpectBool(t *testing.T, field string, expected bool, found bool) {
	if expected != found {
		t.Errorf("%s [%v], found '%v'", field, expected, found)
	}
}
