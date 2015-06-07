package main

import (
	"fmt"
)

func pointers() {
	i, j := 21, 427

	p := &i
	*p = 67
	fmt.Println(i)

	p = &j
	*p = *p / 47
	fmt.Println(j)
}

type Car struct {
	Brand string
	Model string
	Year  int
}

func printSlice(desc string, s []int) {
	fmt.Printf("%s len=%d cap=%d content=%v\n", desc, len(s), cap(s), s)
}

func __main() {
	pointers()
	c := Car{"Volkswagen", "Golf", 2012}
	d := &Car{Model: "Panda"}
	e := Car{}
	p := &c
	p.Model = "Polo"
	fmt.Println(c, d, e)
	// ----------------------------
	var a [10]int
	a[0] = 2
	a[5] = 8
	fmt.Println(a)
	s := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d]=%d\n", i, s[i])
	}
	fmt.Println(s[1:4])
	fmt.Println(s[4:])
	fmt.Println(s[:7])

	s2 := make([]int, 5)
	s3 := make([]int, 0, 5)
	printSlice("s2", s2)
	printSlice("s3", s3)
	printSlice("s4", s2[:3])
	printSlice("s5", s3[1:5])
	var z []int
	printSlice("nil", z)
	z = append(z, 1, 2, 3, 4, 5, 77)
	printSlice("z", z)
	for _, w := range z {
		fmt.Printf("%d, ", w)
	}
}
