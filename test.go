package main

import "fmt"

func main() {
	f0()
	fmt.Println("======================")
	f1()
	fmt.Println("======================")
	f2()
}
func f0() {
	e := []int32{1, 2, 3}
	fmt.Println("cap of e before:", cap(e))
	e = append(e, 4, 5, 6, 7)
	fmt.Println("cap of e after:", cap(e))
}
func f1() {
	//3->6->12
	e := []int32{1, 2, 3}
	fmt.Println("cap of e before:", cap(e))
	e = append(e, 4, 5)
	fmt.Println("cap of e after:", cap(e))
	e = append(e, 6, 7)
	fmt.Println("cap of e after:", cap(e))
}
func f2() {
	e := []int32{1, 2, 3}
	fmt.Println("cap of e before:", cap(e))
	e = append(e, 4)
	fmt.Println("cap of e after:", cap(e))
	e = append(e, 5)
	fmt.Println("cap of e after:", cap(e))
	e = append(e, 6)
	fmt.Println("cap of e after:", cap(e))
	e = append(e, 7)
	fmt.Println("cap of e after:", cap(e))
}
