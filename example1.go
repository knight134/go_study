package main
import "fmt"
import "math"


func main(){
    for i:=1;i<9;i++ {
        for j:=1;j<=i;j++ {
            fmt.Printf("%d * %d = %d\t",j,i,j*i)
        }
        fmt.Println()
    }

    for num:=100;num<=999;num++{
        var i=num/100
        var j=num/10%10
        var k=num%10
        if math.Pow(float64(i),3)+math.Pow(float64(j),3)+math.Pow(float64(k),3)==float64(num){
            fmt.Println(num)
        }
    }

    var line int=10
    for i:=0;i<line;i++{
        for j:=0;j<line-i-1;j++{
            fmt.Print(" ")
        }
        for k:=0;k<i*2+1;k++{
            fmt.Print("*")
        }
        fmt.Println()
    }


    var i int=60
    fmt.Printf("%d 的阶乘是 %d\n",i,Factorial(uint64(i)))

    var str string
    fmt.Scan(&str)
    slice:=[]byte(str)
    m:=make(map[byte]int)
    for i:=0;i<len(slice);i++{
        m[slice[i]]++
    }
    for k,v:=range m{
        fmt.Printf("%c:%d\n",k,v)
    }
    
}

func Factorial(n uint64)(result uint64){
    if(n>0){
        result=n*Factorial(n-1)
        return result
    }
    return 1
}
