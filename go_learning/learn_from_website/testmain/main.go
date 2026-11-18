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

// func jia(x ...int) int {
// 	sum := 0
// 	for _, v := range x {
// 		sum += v
// 	}
// 	return sum
// }

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

	// var arr1 [3]int
	// var arr2 [4]int
	// arr1[0] = 1
	// arr1[1] = 2
	// arr1[2] = 3
	// arr2[0] = 4
	// arr2[1] = 5
	// arr2[2] = 6
	// arr2[3] = 7
	// fmt.Println(arr1)
	// fmt.Println(arr2)

	// var aar1 = [...]string{"你", "好", "GOlang"}
	// fmt.Println(aar1)

	// var arr1 = [3]int{1, 2, 3}
	// for i := 0; i < len(arr1); i++ {
	// 	fmt.Println(arr1[i])
	// }

	// var arr1 = []int{1, 2, 3}
	// arr2 := arr1
	// arr1[0] = 100
	// fmt.Println(arr1)
	// fmt.Println(arr2)

	// var arr1 = [3][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// 	{5, 6},
	// }
	// arr1[0][0] = 100
	// fmt.Println(arr1)

	// var arr1 []int
	// fmt.Printf("%v-%T-%v", arr1, arr1, len(arr1))
	// fmt.Println()

	// a := [5]int{1, 2, 3, 4, 5}
	// b := a[:]
	// c := a[1:4]
	// fmt.Println(b)
	// fmt.Println(c)

	// var arr1 = make([]int, 5, 10)
	// fmt.Printf("%v-%T-%v-%v\n", arr1, arr1, len(arr1), cap(arr1))

	// var arr1 [1]int
	// arr1[0] = 1
	// fmt.Println(arr1)
	// var arr2 []int
	// arr2 = append(arr2, 1)
	// fmt.Println(arr2)

	// var arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// arr1 = append(arr1[:3], arr1[6:]...)
	// fmt.Println(arr1)

	// var userinfo = make(map[string]string)
	// userinfo["username"] = "潘徐锦鹏"
	// userinfo["age"] = "22"
	// userinfo["sex"] = "男"
	// userinfo["username"] = "张三"
	// v, k := userinfo["username"]
	// delete(userinfo, "username")
	// userinfo["username"] = "潘徐锦鹏"
	// fmt.Println(v, k)
	// fmt.Println(userinfo)
	// fmt.Println(userinfo["username"])
	// fmt.Println(userinfo["age"])
	// fmt.Println(userinfo["sex"])

	// var userinfo = make([]map[string]string, 2)
	// if userinfo[0] == nil {
	// 	userinfo[0] = make(map[string]string)
	// 	userinfo[0]["username"] = "潘徐锦鹏"
	// 	userinfo[0]["age"] = "22"
	// 	userinfo[0]["sex"] = "男"
	// }
	// if userinfo[1] == nil {
	// 	userinfo[1] = make(map[string]string)
	// 	userinfo[1]["username"] = "张三"
	// 	userinfo[1]["age"] = "20"
	// 	userinfo[1]["sex"] = "女"
	// }
	// fmt.Println(userinfo)

	// var userinfo = make(map[string]string)
	// userinfo["user1"] = "潘徐锦鹏" + "," + "22" + "," + "男"
	// userinfo["user2"] = "张三" + "," + "20" + "," + "女"
	// fmt.Println(userinfo)

	// var user1 = make([]string, 3)
	// user1[0] = "潘徐锦鹏"
	// user1[1] = "22"
	// user1[2] = "男"
	// var user2 = make([]string, 3)
	// user2[0] = "张三"
	// user2[1] = "20"
	// user2[2] = "女"
	// var userinfo2 = make(map[string][]string)
	// userinfo2["user1"] = user1
	// userinfo2["user2"] = user2
	// fmt.Println(userinfo2)

	// fmt.Println("请输入三个数字：")
	// var x, y, z int
	// fmt.Scanln(&x, &y, &z)
	// result := jia(x, y, z)
	// fmt.Println(result)
}
