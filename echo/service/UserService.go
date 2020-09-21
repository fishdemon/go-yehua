package service

import "fmt"

type UserService struct {

}

func (userService *UserService) Print()  {
	fmt.Println("this is user service")
}
