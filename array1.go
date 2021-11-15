package main
import "fmt"

func main(){
    var numbers []int
    printSlice(numbers)

    numbers=append(numbers,0)
    printSlice(numbers)

    numbers=append(numbers,1)
    printSlice(numbers)
    var tem []int 
    fmt.Println("同时向切片中添加多个元素")
    tem=append(numbers,2,3,4)
    printSlice(numbers)
    printSlice(tem)
    numbers=tem

    numbers1:=make([]int,len(numbers),(cap(numbers))*2)

    copy(numbers1,numbers)
    printSlice(numbers1)
}

func printSlice(x []int){
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

