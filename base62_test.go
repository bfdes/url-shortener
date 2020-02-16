package main

import (
	"math/rand"
	"testing"
)

type testInput struct {
	decoded int
	encoded string
}

var testInputs = []testInput{
	{0, ""},
	{1, "1"},
	{62, "01"},
	{1504, "go"},
}

func TestEncode(t *testing.T) {
	for _, testInput := range testInputs {
		actual, err := Encode(testInput.decoded)
		if err != nil {
			t.Fatal(err)
		}
		if actual != testInput.encoded {
			t.Errorf("%d encoded to %s, not %s", testInput.decoded, actual, testInput.encoded)
		}
	}
}

func TestEncodeNegative(t *testing.T) {
	arg := -1
	encoded, err := Encode(arg)
	if err == nil {
		t.Errorf("negative argument %d encoded to %s", arg, encoded)
	}
}

func TestDecode(t *testing.T) {
	for _, testInput := range testInputs {
		actual, err := Decode(testInput.encoded)
		if err != nil {
			t.Fatal(err)
		}
		if actual != testInput.decoded {
			t.Errorf("%s decoded to %d, not %d", testInput.encoded, actual, testInput.decoded)
		}
	}
}

func TestDecodeIllegalCharacter(t *testing.T) {
	arg := "!llegal"
	decoded, err := Decode(arg)
	if err == nil {
		t.Errorf("malformed slug %s decoded to %d", arg, decoded)
	}
}

func TestEncodeDecode(t *testing.T) {
	testInputs := []int{0, 1, 62}
	for i := 0; i < 20; i++ {
		testInputs = append(testInputs, rand.Int())
	}
	for _, testInput := range testInputs {
		encoded, err := Encode(testInput)
		if err != nil {
			t.Fatal(err)
		}
		actual, err := Decode(encoded)
		if err != nil {
			t.Fatal(err)
		}
		if actual != testInput {
			t.Errorf("%d does not roundtrip", testInput)
		}
	}
}
