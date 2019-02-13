package main

import (
	"fmt"
	"time"
)



func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total  // send total to c
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}
func fibonacci2(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <-quit:
			fmt.Println("quit")
			return
		case <- time.After(5 * time.Second):
			println("timeout")
			break
		}
	}
}

func set(a string,c chan string){
    c <- a

}
func get(c chan string)string{
    return <-c
}

func main(){
	ch := make(chan string)
    go set("hello",ch)
	fmt.Println(get(ch))
	//select {
	//case <-ch:fmt.Println("ok")}
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y:= <-c, <-c  // receive from c
	fmt.Println(x, y,  x + y)
    ci := make(chan int,10)
	go fibonacci(cap(ci), ci)
	for i := range ci {
		fmt.Println(i)
	}
	cj := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-cj)
		}
		quit <- 0
	}()
	fibonacci2(cj, quit)


}
