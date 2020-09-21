package service

import "fmt"

var flowService *FlowService = new(FlowService)

func NewFlowService() *FlowService {
	return flowService
}

type FlowService struct {

}

func (flowService *FlowService) Print()  {
	fmt.Println("this is flow service")
}