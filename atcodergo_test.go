package atcodergo

import "testing"

func TestIsEven(t *testing.T) {
	evenNum := 8
	evenActual := isEven(evenNum)
	evenExpected := true
	if evenActual != evenExpected {
		t.Errorf("got: %v, want: %v", evenActual, evenExpected)
	}

	oddNum := 1
	oddActual := isEven(oddNum)
	oddExpected := false
	if oddActual != oddExpected {
		t.Errorf("got: %v, want: %v", oddActual, oddExpected)
	}
}
