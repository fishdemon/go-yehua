package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Age int8 `json:"age"`
}

func Stuct2json()  {
	user := User{"12345", "allen", 28}
	jsonByte, err :=json.Marshal(user)		// 二进制
	if err == nil {
		fmt.Println(string(jsonByte))		// 用 string() 强转
	}
}

func Slice2json()  {
	// string slice
	arr := []string{"allen","john","june"}
	arrByte, err :=json.Marshal(arr)		// 二进制
	if err == nil {
		fmt.Println(string(arrByte))		// 用 string() 强转
	}

	// struct slice
	users := []User{
		{"12345","allen",28},
		{"12346","john", 15},
		{"12347","june",37},
	}
	userByte, err :=json.Marshal(users)		// 二进制
	if err == nil {
		fmt.Println(string(userByte))		// 用 string() 强转
	}

}

func Map2json()  {
	// string, string map
	userMap := map[string]string{"id":"12345","name":"allen","age":"28"}
	userByte, err := json.Marshal(userMap) // 二进制
	if err == nil {
		fmt.Println(string(userByte)) // 用 string() 强转
	}
}

func json2struct()  {
	userStr := "{\"id\":\"12345\",\"name\":\"allen\",\"age\":28}"
	var user User
	err := json.Unmarshal([]byte(userStr), &user)
	if err == nil {
		fmt.Println(user)
	}
}

func json2slice()  {
	arrStr := "[\"allen\",\"john\",\"june\"]"
	var names []string
	err := json.Unmarshal([]byte(arrStr), &names)
	if err == nil {
		fmt.Println(names)
	}

	usersStr := `[{"id":"12345","name":"allen","age":28},{"id":"12346","name":"john","age":15},{"id":"12347","name":"june","age":37}]`
	var users []User
	err1 := json.Unmarshal([]byte(usersStr), &users)
	if err1 == nil {
		fmt.Println(users)
	}
}

func json2map()  {
	mapStr := "{\"age\":\"28\",\"id\":\"12345\",\"name\":\"allen\"}"
	var user map[string]string
	err := json.Unmarshal([]byte(mapStr), &user)
	if err == nil {
		fmt.Println(user)
	}
}


func main() {
	fmt.Println("starting...")
	Stuct2json()
	Slice2json()
	Map2json()

	fmt.Println("json 2 other")
	json2struct()
	json2slice()
	json2map()
}

