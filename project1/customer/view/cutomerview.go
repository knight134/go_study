package main
import "fmt"
import "project1/customer/service"
import "project1/customer/model"
type customerView struct {
    key string //用户输入内容
    loop bool  //是否循环显示菜单

    customerService *service.CustomerService
}

//显示所有的客户信息
func (this *customerView)list(){
    customers:=this.customerService.List()
    fmt.Println("---------------------客户列表-------------------")
    fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
    for i:=0;i<len(customers);i++{
        fmt.Println(customers[i].GetInfo())
    }
    fmt.Println("-----------------------客户列表完成----------------")
}
//添加客户 从用户输入
func (this *customerView)add(){
    fmt.Println("-----------------------添加客户--------------------")
    fmt.Println("姓名：")
    name:=""
    fmt.Scanln(&name)
    fmt.Println("性别：")
    gender:=""
    fmt.Scanln(&gender)
    fmt.Println("年龄：")
    age:=0
    fmt.Scanln(&age)
    fmt.Println("电话：")
    phone:=""
    fmt.Scanln(&phone)
    fmt.Println("电邮：")
    email:=""
    fmt.Scanln(&email)

    //id是唯一的 系统分配的
    customer := model.NewCustomer2(name,gender,age,phone,email)
    if this.customerService.Add(customer){
        fmt.Println("------------------添加成功------------------")
    }else{
        fmt.Println("------------------添加失败------------------")
    }
}
//删除客户
func (this *customerView)delete(){
    fmt.Println("--------------------删除客户------------------")
    fmt.Println("请选择待删除客户编号（-1退出）：")
    id:=-1
    fmt.Scanln(&id)
    if id==-1{
        return
    }
    fmt.Println("确认是否删除（Y/N）：")
    choice:=""
    fmt.Scanln(&choice)

    if choice=="Y"||choice=="y"{
        if this.customerService.Delete(id){
            fmt.Println("-----------------删除成功------------------")
        }else{
            fmt.Println("-----------------输入的ID不存在-----------------")
        }
    }
}

func (this *customerView)exit(){
    fmt.Println("确认是否退出（Y/N）：")
    for{
        fmt.Scanln(&this.key)
        if this.key=="Y"||this.key=="y"||this.key=="N"||this.key=="n"{
            break
        }
        fmt.Println("输入有误请重新输入")
    }
    if this.key=="Y"||this.key=="y"{
        this.loop=false
    }
}

//显示主菜单
func (this *customerView) mainView(){
    for{
        fmt.Println("--------------------客户信息管理软件--------------------")
        fmt.Println("                    1 添加客户")
        fmt.Println("                    2 修改客户")
        fmt.Println("                    3 删除客户")
        fmt.Println("                    4 客户列表")
        fmt.Println("                    5 退    出")
        fmt.Println("请选择1-5：")
        fmt.Scanln(&this.key)
        switch this.key{
            case "1":
                fmt.Println("添加客户")
                this.add()
            case "2":
                fmt.Println("修改客户")
            case "3":
                fmt.Println("删除客户")
                this.delete()
            case "4":
                fmt.Println("客户列表")
                this.list()
            case "5":
                this.loop=false
                this.exit()
            default:
                fmt.Println("你的输入有误，请重新输入")
            }
        if !this.loop{
            break
        }
    }
    fmt.Println("已经退出")
}

func main(){
    customerView:=customerView{
        key:"",
        loop:true,
    }
    //这里完成对customerView结构体的customerService字段的初始化
    customerView.customerService=service.NewCustomerService()
    //显示菜单1
    customerView.mainView()
}
