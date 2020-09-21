package main

import (
	"fmt"
	"github.com/fishdemon/go-yehua/lang/reflect/test"
	"reflect"
	"strconv"
)

// 定义一个接口
type PersonI interface {
	Info() string
	Eat()
	Sleep()
	Walk()
}

// go 语言虽然没有继承的概念，但如果一个 struct 来实现了接口中的所有方法，我们称之为 继承
// get /api/vi 处理方法 name
type Person struct {
	Name string 	`airport-data:"allen, max=10"`
	Age int64  		`airport-data:"max=50"`
	Sex string		`airport-data:"male,female"`
	Married bool	`airport-data:"true,false"`
	It interface{}  `airport-data:""`
	H HandlerFunc	`api:"m=get,url='',des='',h=''"`
}

func (p *Person) Info() string {
	fmt.Println(p.Name + " " + strconv.Itoa(int(p.Age)) + " " + p.Sex)
	return p.Name + " " + strconv.Itoa(int(p.Age)) + " " + p.Sex
}

func (p *Person) Eat() {
	fmt.Println(p.Name + " is eating")
}

func (p *Person) Sleep() {
	fmt.Println(p.Name + " is sleeping")
}

func (p *Person) Walk()  {
	fmt.Println(p.Name + " is walking")
}

type HandlerFunc func() error

func main() {
	var p *test.Person = new(test.Person)
	pt := reflect.TypeOf(p)
	fmt.Println(pt.PkgPath())
	// interface{} 是空接口，可以适配任何类型

	// 操作接口
	var personI PersonI
	personI = new(Person)
	fmt.Println(personI)
	// Type
	t := reflect.TypeOf(personI)
	fmt.Println("type: ", t)
	// Value
	fmt.Println("value: ", reflect.ValueOf(personI))


	// 操作struct
	p1 := Person{
		Name:    "allen",
		Age:     0,
		Sex:     "19",
		Married: false,
		It:       nil,
	}
	// 反射定律1：将接口类型变量转换为反射类型对象（Type/Value）
	t = reflect.TypeOf(p1)
	fmt.Println(t.NumField())
	fmt.Println(t.NumMethod())
	fmt.Println(t.PkgPath())
	//fmt.Println(t.NumIn())
	//fmt.Println(t.NumOut())
	//fmt.Println(t.In(2))
	//fmt.Println(t.Out(1))

	// 操作属性/标签
	f, _ := t.FieldByName("Name")
	fmt.Println(fmt.Sprintf("index:%d name:%s type:%s tag:%s pkg:%s", f.Index, f.Name, f.Type, f.Tag, f.PkgPath))

	v := reflect.ValueOf(&p1)
	v.Index(0).SetString("Jack")
	fmt.Println(p1)

}
