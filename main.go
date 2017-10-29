package main

import (
	"fmt"
	"github.com/vrowello/reg-area/apothem"
	"github.com/vrowello/reg-area/area"
	"github.com/vrowello/reg-area/perimeter"
)

var (
	side_n float64
	side_l float64
	angle  float64
)

func main() {
	a := make(chan float64)
	p := make(chan float64)

	fmt.Printf("Number of Sides: ")
	fmt.Scanln(&side_n)
	fmt.Printf("Length of Side: ")
	fmt.Scanln(&side_l)

	go apothem.Apothem(side_n, side_l, a)
	go perimeter.Perimeter(side_n, side_l, p)

	fmt.Println(area.Area(<-a, <-p))
}
