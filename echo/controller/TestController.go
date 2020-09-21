package controller

import "github.com/fishdemon/go-yehua/echo/service"

type TestController struct {
	flowService *service.FlowService
	testService *service.TestService
	userService *service.UserService
}

var flowService *service.FlowService = service.NewFlowService()

func (testController *TestController) Print()  {
	flowService.Print()
}
