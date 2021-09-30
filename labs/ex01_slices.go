package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	bild := make([][]uint8, dy)
	
	for i := range bild {
		bild[i] = make([]uint8, dx)
		for j := range bild[i] {
			bild[i][j] = (uint8(i) ^ uint8(j)) + 32
		}
	}
	return bild
}

func main() {
	pic.Show(Pic)
}
