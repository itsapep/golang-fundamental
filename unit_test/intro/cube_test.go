package intro

import (
	"testing"
)

var (
	sideSuccess    Cube    = Cube{Side: 4}
	sideFailed     Cube    = Cube{Side: -4}
	expectedVolume float64 = 64
	expectedArea   float64 = 96
)

func TestVolumeSuccess(t *testing.T) {

	actual, err := sideSuccess.Volume()
	if actual != expectedVolume {
		t.Errorf("Volume() = %v, want %v", actual, expectedVolume)

	}

	if err != nil {
		t.Errorf("Volume() = %v, want %v", err, nil)
	}
}

func TestVolumeFailed(t *testing.T) {

	actual, err := sideFailed.Volume()
	if actual < 0 {
		t.Errorf("Number must be greater than 0")
	}

	if err == nil {
		t.Errorf("Number must be greater than 0")
	}
}

func TestAreaSuccess(t *testing.T) {
	actual, err := sideSuccess.Area()
	if actual != expectedArea {
		t.Errorf("Area() = %v, want %v", actual, expectedArea)
		t.Errorf("Error: %v", err)
	}
}

func TestAreaFailed(t *testing.T) {
	actual, err := sideFailed.Area()
	if actual <= 0 {
		t.Errorf("Number must be greater than 0")
	}

	if err == nil {
		t.Errorf("Number must be greater than 0")
	}
}

func BenchmarkVolumeTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sideSuccess.Volume()
		sideFailed.Volume()
		sideSuccess.Area()
		sideFailed.Area()
	}
}
