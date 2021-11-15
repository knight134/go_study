package service
import "project1/customer/model"
type CustomerService struct{
    customers []model.Customer
    //当前客户含有多少个客户
    coustmerNum int
}

// 返回 *CustomerService

func NewCustomerService() *CustomerService{
    customerservice:=&CustomerService{}
    customerservice.coustmerNum=1
    customer:=model.NewCustomer(1,"张三","男",20,"121","zhangshang@xinlaing.com")
    customerservice.customers=append(customerservice.customers,customer)

    return customerservice
}

//返回客户切片
func (this *CustomerService)List() []model.Customer{
    return this.customers
}
//添加客户到customers数组中
func (this *CustomerService)Add(customer model.Customer)bool{
    this.coustmerNum++
    customer.Id=this.coustmerNum
    this.customers=append(this.customers,customer)
    return true
}

//根据id删除客户
func (this *CustomerService)Delete(id int)bool{
    index:=this.FindById(id)
    if index==-1{
        return false
    }
    this.customers=append(this.customers[:index],this.customers[index+1:]...)
    return true
}

//根据id返回客户
func (this *CustomerService)FindById(id int) int{
    index:=-1
    for i:=0;i<len(this.customers);i++{
        if this.customers[i].Id==id{
            index=i
        }
    }
    return index
}
