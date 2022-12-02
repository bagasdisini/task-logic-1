package main

import (
	"fmt"
	"sort"
)

func main() {
	angka := []int{5, 7, 7, 9, 10, 4, 5, 10, 6, 5}
	sort.Ints(angka)
	var jumlah = 0
	y := 1
	for i := 0; i < len(angka); i++ {
		if y != len(angka) {
			if angka[i] == angka[y] {
				jumlah += 1
				i += 1
				y += 2
			} else {
				y += 1
			}
		}
	}
	fmt.Println(jumlah)
}
