package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	dySlice := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		dxSlice := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			dxSlice[j] = uint8(i * j)
		}
        dySlice[i] = dxSlice
	}
	return dySlice
}

func _main() {
	pic.Show(Pic)
}
