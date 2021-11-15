package main
import (
    "fmt"
)
//定义数据写入器
type DataWriter interface{
    WriteData(data interface{}) error
}

type file struct{
}

func (d file) WriteData(data interface{}) error{
    //模拟写入数据
    fmt.Println("WriteData:",data)
    return nil
}

type Socket struct{
}
func (s *Socket) Write(p []byte)(n int,err error){
    return 0,nil
}
func (s *Socket) Close() error{
    return nil
}
type Writer interface{
    Write(p []byte)(n int,err error)
}
//第二个接口
type Closer interface{
    Close() error
}
// 使用定义好的接口
func usingCloser(closer Closer){
    closer.Close()
}
func usingWriter(writer Writer){
    writer.Write(nil)
}

type Service interface{
    Start()
    Log(string)
}
type Logger struct{
}
func (g *Logger)Log(l string){
    fmt.Println(l)
}
type GameService struct{
    Logger
}
func (g *GameService) Start(){
}

func main(){
    f:=new(file)
    var writer DataWriter
    //将接口赋值f 也就是*file类型
    writer=f

    writer.WriteData("data")
    var s Service=new(GameService)
    s.Start()
    s.Log("这是日志管理器服务hello")

}

