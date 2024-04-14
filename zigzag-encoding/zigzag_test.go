package main

import (
	"testing"
)

func TestZiagZag(t *testing.T) {
	tests := []struct {
		val             int32
		expectedEncoded int32
	}{
		{-20, 39},
		{-19, 37},
		{-18, 35},
		{-17, 33},
		{4, 8},
		{5, 10},
	}

	for _, test := range tests {
		encodedVal := zigzagEncode(test.val)
		if encodedVal != test.expectedEncoded {
			t.Fatalf("encode not correct value expected %v, actual %v", test.expectedEncoded, encodedVal)
		}
		decodedVal := zigzadDecode(encodedVal)
		if decodedVal != test.val {
			t.Fatalf("decode not correct value expected %v, actual %v", test, decodedVal)
		}

		t.Logf("val: %v, encoded: %v, decoded: %v\n", test.val, encodedVal, decodedVal)
	}
}
