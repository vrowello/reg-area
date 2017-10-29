package perimeter

func Perimeter(n float64, l float64, p chan float64) {
	p <- n * l
}
