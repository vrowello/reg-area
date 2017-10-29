package apothem

import (
	"math"
)

var (
	angle float64
)

func Apothem(n float64, l float64, a chan float64) {
	angle = ((180 - (360 / n)) / 2) * 0.01745329252 //finds the angle in degrees and converts to radians
	a <- math.Tan(angle) * (l / 2)
}
