package service

import "fmt"

type TestService struct {

}

func (testService *TestService) Print()  {
	fmt.Println("this is test service")
}