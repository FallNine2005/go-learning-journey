package main

import (
	"fmt"
	"unsafe"
)

func getUserinfo() (string, int, string) {
	return "潘徐锦鹏", 22, "男"
}

const qq = 3178412159

const (
	n0 = iota
	n1
	n2
	n3
)

var username, age, _ = getUserinfo()

var (
	a1 int8    = 8
	a2 int16   = 16
	pi float64 = 3.1415926
	T          = true
	Pi string  = "3.1415926"
)

func main() {
	a, b, c := 10, 20, 30
	fmt.Printf("a=%v b=%v c=%v\n", a, b, c)
	fmt.Printf("username=%v,age=%v,qq=%v\n", username, age, qq)
	fmt.Printf("n0=%v,n1=%v,n2=%v,n3=%v\n", n0, n1, n2, n3)
	fmt.Println(unsafe.Sizeof(qq))
	fmt.Println(int16(a1) + a2)
	fmt.Printf("pi(f)=%f,pi(v)=%v,pi(.2f)=%.2f\n", pi, pi, pi)
	fmt.Printf("T=%v\n", T)
	if T {
		fmt.Println("T is true")
	}
	testMaingo := "~\\projects\\go_learning\\learn_from_website\\testmain\\main.go"
	fmt.Println(testMaingo)
	fmt.Println(len(testMaingo))

	// P1 := strings.Split(Pi, ".")
	// P2 := strings.Join(P1, ".")
	// P3 := []string{"3", "1415926"}
	// P4 := strings.Join(P3, ".")
	// P5 := strings.Contains(Pi, P2)
	// P6 := strings.Index(Pi, P2)
	// P7 := strings.LastIndex(Pi, P2)
	// fmt.Println(P1)
	// fmt.Println(P2)
	// fmt.Println(P3)
	// fmt.Println(P4)
	// fmt.Println(P5)
	// fmt.Println(P6)
	// fmt.Println(P7)
	// for i := 0; i < len(Pi); i++ {
	// 	fmt.Printf("%c\n", Pi[i])
	// }
	// for _, v := range Pi {
	// 	fmt.Printf("%v(%c)", v, v)
	// }
}
