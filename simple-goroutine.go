package main

import "fmt"



func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total  // send total to c
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
	go sum(a[len(a)/2:], c)
	go sum(a[len(a)/2:], c)
	x, y:= <-c, <-c  // receive from c

	fmt.Println(x, y,  x + y)


}
