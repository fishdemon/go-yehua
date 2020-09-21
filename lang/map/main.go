package main

import "fmt"

func main() {
	userMap := map[string]string{"name":"allen","age":"18","sex":"male"}

	fmt.Println(userMap["test"])

	// go并没有提供map中是否存在某个key的判断方法，我们可以用下面的方式来判断
	// 判断 key 是否存在
	if _,ok := userMap["airport-data"]; ok {
		fmt.Println(ok)
	}

	// 同时也没有提供 array 中存在某个item的判断方法，我们也可以将 array 转化为map的key来判断，这样可以提高效率

}
