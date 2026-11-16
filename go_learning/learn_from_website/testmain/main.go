package main

import (
	"fmt"
)

const (
// n0 = iota
// n1
// n2
// n3
)

func getUserinfo() (string, int, string) {
	return "潘徐锦鹏", 22, "男"
}

var (
	username, age, _     = getUserinfo()
	phone            int = 18952536510
	qq               int = 3178412159
	// a1 int8    = 8
	// a2 int16   = 16
	// pi float64 = 3.1415926
	// T          = true
	// Pi string  = "3.1415926"
)

func main() {
	fmt.Printf("username=%v,age=%v,qq=%v,phone=%v\n", username, age, qq, phone)
	// a, b, c := 10, 20, 30
	// fmt.Printf("a=%v b=%v c=%v\n", a, b, c)
	// fmt.Printf("n0=%v,n1=%v,n2=%v,n3=%v\n", n0, n1, n2, n3)
	// fmt.Println(unsafe.Sizeof(qq))
	// fmt.Println(int16(a1) + a2)
	// fmt.Printf("pi(f)=%f,pi(v)=%v,pi(.2f)=%.2f\n", pi, pi, pi)
	// fmt.Printf("T=%v\n", T)
	// if T {
	// 	fmt.Println("T is true")
	// }
	// testMaingo := "~\\projects\\go_learning\\learn_from_website\\testmain\\main.go"
	// fmt.Println(testMaingo)
	// fmt.Println(len(testMaingo))

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
	// Pi = fmt.Sprintf("%f", pi)
	// fmt.Println(Pi)

	// str1 := strconv.FormatInt(int64(phone), 10)
	// fmt.Println("str1 =", str1)
	// str2 := strconv.FormatFloat(float64(phone), 'f', -1, 64)
	// fmt.Println("str2 =", str2)
	// str3 := strconv.FormatBool(true)
	// fmt.Println("str3 =", str3)
	// str4, _ := strconv.ParseInt(str1, 10, 64)
	// fmt.Println("str4 =", str4)

	// 	a := 10
	// 	b := 20
	// 	fmt.Println(a == b)
	// 	fmt.Println(a != b)
	// 	fmt.Println(a > b)
	// 	fmt.Println(a < b)
	// 	fmt.Println(a >= b)
	// 	fmt.Println(a <= b)
	// 	c := a == b
	// 	fmt.Println(c)
	// 	if c {
	// 		fmt.Println("a==b")
	// 	} else {
	// 		fmt.Println("a!=b")
	// 	}
	// 	d := c || !c
	// 	fmt.Println(d)

	// 	i := 0
	// 	r := 0
	// xxxx:
	// 	for ; i <= 10; i++ {
	// 		for ; i <= 10 && i%2 == 0; i++ {
	// 			fmt.Print(i)
	// 		}
	// 		if i >= 10 {
	// 			i = 0
	// 			r += 1
	// 			fmt.Println()
	// 		} else if r == 5 {
	// 			break xxxx
	// 		}
	// 	}

	// for _, v := range username {
	// 	fmt.Printf("%c", v)
	// }

	// switch username {
	// case "潘徐锦鹏":
	// 	fmt.Println("潘徐锦鹏")
	// case "张三":
	// 	fmt.Println("张三")
	// default:
	// 	fmt.Println("false")
	// }

}
