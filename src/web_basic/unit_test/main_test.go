package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

/**
* define required variable
*/
var (
	kubus Kubus = Kubus{4}
	volume float64 = 64
	luas float64 = 96
	keliling float64 = 48
)

/**
* method for testing
*/

// without testify package
// func TestHitungVolume(t *testing.T) {
//     t.Logf("Volume : %.2f", kubus.Volume())

//     if kubus.Volume() != volume {
//         t.Errorf("SALAH! harusnya %.2f", volume)
//     }
// }

// func TestHitungLuas(t *testing.T) {
//     t.Logf("Luas : %.2f", kubus.Luas())

//     if kubus.Luas() != luas {
//         t.Errorf("SALAH! harusnya %.2f", luas)
//     }
// }

// func TestHitungKeliling(t *testing.T) {
//     t.Logf("Keliling : %.2f", kubus.Keliling())

//     if kubus.Keliling() != keliling {
//         t.Errorf("SALAH! harusnya %.2f", keliling)
//     }
// }

// using testify package
func TestHitungVolume(t *testing.T) {
    assert.Equal(t, kubus.Volume(), volume, "perhitungan volume salah")
}

func TestHitungLuas(t *testing.T) {
    assert.Equal(t, kubus.Luas(), luas, "perhitungan luas salah")
}

func TestHitungKeliling(t *testing.T) {
    assert.Equal(t, kubus.Keliling(), keliling, "perhitungan keliling salah")
}

/**
* method for benchmark
*/
func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Luas()
	}
}

/**
* run test: go test main.go main_test.go -v
* run test with benchmark: go test main.go main_test.go -bench=.
*/
