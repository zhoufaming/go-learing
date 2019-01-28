package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var j []int = primes[1:4]
	fmt.Println(len(j))
	fmt.Println(j)
	
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)
	// 拓展其长度
	s = s[:4]
	printSlice(s)
	// 舍弃前两个值 s[0]、s[1] 舍去 从从头数 不关注尾部
	s = s[2:]
	printSlice(s)
	
	a := make([]int, 5)
	printSlice1("a", a)

	b := make([]int, 0, 5)
	printSlice1("b", b)

	c := b[:2]
	printSlice1("c", c)

	d := c[2:5]
	printSlice1("d", d)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)//cap求容量
	
}
