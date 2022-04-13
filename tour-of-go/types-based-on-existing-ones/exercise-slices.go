package main

import (
	"golang.org/x/tour/pic"
	"math"
)

func f(x, y int) uint8 {
	return uint8(math.Atan2(float64(x-130), float64(y-130))*122.6 + float64(x))
}

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for y := range picture {
		picture[y] = make([]uint8, dx)
		for x := range picture[y] {
			picture[y][x] = f(x, y)
		}
	}
	return picture
}
func main() {
	pic.Show(Pic)
}
