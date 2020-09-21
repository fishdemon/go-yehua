package main

import "fmt"

// 普通结构体，相当于 java bean
// 结构体成员的访问通过 变量名.成员名来访问
type Person struct {
	ID int32
	Name string
	Age int8
}

// 带 tag 属性的结构体
// 这个 tag 只能反射获取
type Person1 struct {
	ID int32    "id"
	Name string "name"
	Age int8    "age"
	_ string			// _为空标识符，表示这个字段被丢弃，不会被使用
}

// 匿名字段结构体
// 匿名的字段与类型同名
type Person2 struct {
	ID int32
	Name string
	Age int8
	int			// 相当于 int  int
	bool		// 相当于 bool  bool
}

// 具名嵌套结构体
type Person3 struct {
	ID int32
	Name string
	Age int8
	Person Person	// 访问这里的字段，通过 Person3.Person.ID 访问
}

// 匿名嵌套结构体, 相当于继承
// 若父子结构体含有相同的字段，则子结构覆盖父结构体
type Person4 struct {
	ID int32		// 覆盖 Person2
	Name string		// 覆盖 Person2
	Age int8		// 覆盖 Person2
	Person2			// Person4 会继承 Person2 的所有成员，可以通过 Person4.int 直接访问它的属性
}

// 嵌套自身结构体
// 可以组成一个二叉树
type Person5 struct {
	ID int32
	Name string
	Age int8
	left *Person5	// 嵌套自身必须加*，传的是指针
	right *Person5
}

// 带方法的结构体，相当于类的方法
// 结构体内只能定义成员，而方法是单独进行绑定定义，相当于做一个方法签名
// 这里给 Person struct 加一个 eat() 方法
func (p *Person) Eat() bool {	// 为 Person 类型绑定 Eat() 的方法，*Person 为指针引用，以为这可以修改自身的指
	isEat := p.Age > 25
	if isEat {
		fmt.Println(p.Name + " has eaten")
		return isEat
	}

	fmt.Println(p.Name + " has not eaten")
	return isEat			// 这个方法是属于这个类型的，不属于具体的对象，
}


// 结构体默认是值传递,	相当于传了一个拷贝，并不是原实例
// 这与 java 有很大的区别，各位 java 大大要注意
func changeName(p Person) {
	p.Name = "change"
}

// 引用传递必须使用指针
func changeName1(p *Person) {
	p.Name = "change"
}

/**
 * 结构体相当于 Java 中的 对象，可以自定义成员属性
 * 还有可以绑定方法
 */
func main() {
	p := Person{1,"allen", 28}
	p1 := Person{2,"json", 12}

	p.Eat()
	p1.Eat()

	changeName(p)
	fmt.Println("changeName " + p.Name)			// changeName allen
	changeName1(&p)
	fmt.Println("changeName1 " + p.Name)		// changeName1 change
}