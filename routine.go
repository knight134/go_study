package main
import "fmt"

import "time"

func say(s string){
    for i:=0;i<15;i++{
        time.Sleep(100*time.Millisecond)
        fmt.Println(s)
    }
}
func main(){
    go say("这是协程")
    say("hello")
}
