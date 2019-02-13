package main

import "fmt"

func main(){

    type skills struct {
        Name string
    }
    var people = skills{Name:"use python"}
    var iftrue bool = true
    var num int = 32
    var inputnum  int
    var inputstr string
    var scanstr string
    fmt.Printf("%v\n",people)
    fmt.Printf("%+v\n",people)
    fmt.Printf("%#v\n",people)
    fmt.Printf("%T\n",people)
    fmt.Printf("%%\n")
    fmt.Printf("%t\n",iftrue)
    fmt.Printf("%d\n",num)
    fmt.Printf("%o\n",num)
    fmt.Printf("%b\n",num)
    fmt.Printf("%x\n",num)
    fmt.Printf("%X\n",num)
    //Print 将参数列表 a 中的各个参数转换为字符串并写入到标准输出中,非字符串参数之间会添加空格
    fmt.Print("123",num,456)
    fmt.Println()
    //// Println 功能类似 Print，只不过最后会添加一个换行符 所有参数之间会添加空格
    fmt.Println("123",num,456)
    //Scanln 和 Scan 类似，只不过遇到换行符就停止扫描
    fmt.Scanln(&inputnum,&inputstr)
    fmt.Println(inputnum,inputstr)
        //将转换结果以字符串形式返回
    a := fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
    fmt.Println(a)
    fmt.Sscan("scaninput",&scanstr)
    fmt.Println(scanstr)
    }
