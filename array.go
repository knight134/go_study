package main

import "fmt"

//切片是动态数组
func main(){
    var numbers=make([]int,3,5)
    var num1 []int
    printSlice(numbers)
    if(num1==nil){
        fmt.Printf("切片是空的\n")
    }
    //切片截取
    num2:=[]int{0,1,2,3,4,5,6,7,8}
    printSlice(num2)
    fmt.Println("num2==",num2)
    fmt.Println("num2[1:4]==",num2[1:4])
    fmt.Println("num2[:3]==",num2[:3])
    fmt.Println("num2[4:]==",num2[4:])
    
    num3:=make([]int,0,5)
    printSlice(num3)

    num4:=num2[:2]
    printSlice(num4)

    num5:=num2[2:5]
    printSlice(num5)
}

func printSlice(x []int){
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
