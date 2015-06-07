package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
	"time"

	"github.com/eric-fouillet/gotour/stringutil"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, java, python = true, false, "yes!"

func sum() int {
	sum := 1
	for sum < 10000 {
		sum += sum
	}
	return sum
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g > %g", v, lim)
	}
	return lim
}

var newtonIterations int = 20

func SqrtNewton(x float64) float64 {
	v := 1.0
	for i := 0; i < newtonIterations; i++ {
		v = v - (math.Pow(v, 2)-x)/(2*v)
	}
	return v
}

func checkOS() {
	fmt.Print("Go runs on: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}
}

func checkDay() {
	today := time.Now().Weekday()
	switch time.Sunday {
	case today:
		fmt.Println("today")
	case today + 1:
		fmt.Println("In 1 day")
	default:
		fmt.Printf("In %d day", time.Saturday-today)
	}
}

func checkTime() {
	now := time.Now().Hour()
	switch {
	case now < 12:
		fmt.Println("Good morning !")
	case now < 18:
		fmt.Println("Good afternoon !")
	default:
		fmt.Println("Good evening !")
	}
}

func oldmain() {
	defer fmt.Println("Finished !!!!!!!! ^_^")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println(stringutil.Reverse("Hello, World !"))
	fmt.Println(split(17))
	i := 1
	fmt.Println(i, c, java, python)
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	var t, j int32 = 3, 4
	var l float64 = math.Sqrt(float64(t*t + j*j))
	var k uint32 = uint32(l)
	fmt.Println(i, j, k)
	fmt.Println(sum())
	fmt.Println(sqrt(18), sqrt(-19))
	fmt.Println(pow(4, 56, 345), pow(2, 3, 345))
	fmt.Println("============ NEWTOWN SQRT ==============")
	value := 2.0
	vn := SqrtNewton(value)
	vs := math.Sqrt(value)
	fmt.Println("Newton method:", vn)
	fmt.Println("Standard library:", vs)
	fmt.Println("Diff:", math.Abs(vs-vn))
	fmt.Println("============ NEWTOWN SQRT ==============")
	checkOS()
	checkDay()
	checkTime()
}
