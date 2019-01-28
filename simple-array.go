package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
   for index, num := range a{
     fmt.Println(num,a[index])
}
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
