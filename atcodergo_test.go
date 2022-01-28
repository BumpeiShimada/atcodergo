package atcodergo

import (
	"encoding/json"
	"testing"
)

func TestIsEven(t *testing.T) {
	evenNum := 8
	evenActual := IsEven(evenNum)
	evenExpected := true
	if evenActual != evenExpected {
		t.Errorf("got: %v, want: %v", evenActual, evenExpected)
	}

	oddNum := 1
	oddActual := IsEven(oddNum)
	oddExpected := false
	if oddActual != oddExpected {
		t.Errorf("got: %v, want: %v", oddActual, oddExpected)
	}
}

func TestReverseString(t *testing.T) {
	word := "dream wonderland"
	actual := ReverseString(word)
	expected := "dnalrednow maerd"
	if actual != expected {
		t.Errorf("got: %v, want: %v", actual, expected)
	}
}

func TestUnique(t *testing.T) {
	slice := []string{
		"alice",
		"alice",
		"bob",
		"dave",
	}
	actualSlice := Unique(slice)
	actualJSON, _ := json.Marshal(actualSlice)
	actual := string(actualJSON)

	expectedSlice := []string{
		"alice",
		"bob",
		"dave",
	}
	expectedJSON, _ := json.Marshal(expectedSlice)
	expected := string(expectedJSON)

	if string(actual) != expected {
		t.Errorf("got: %v, want: %v", actual, expected)
	}
}
