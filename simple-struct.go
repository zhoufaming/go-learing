package main

import "fmt"

type Vertex struct { //结构体的定义

	X int

	Y int

	z string

}

func print(vertex Vertex){

	fmt.Println(vertex.X,vertex.Y,vertex.z)

}

func main() {

	vertex := Vertex{X: 12, Y: 13, z: "12"} //结构体的声明
	vertex.z="jake" //结构体中变量的访问与赋值
	fmt.Println(vertex)

	var vertex1 Vertex //声明vertex1为Vertex类型
	vertex1.X = 123
	vertex1.Y =  456
	print(vertex1)

	vertex2 := &vertex1
	fmt.Println(*vertex2)

	var vertex3 *Vertex //指针变量可以存储结构体变量的地址
	vertex3 = &vertex
	fmt.Println(*vertex3)


}
