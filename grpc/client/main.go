package main

import (
	"github.com/fishdemon/go-yehua/grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
)

const (
	address = "localhost:12345"
)


func addUser(client proto.UserServiceClient, customer *proto.User) {
	resp, err := client.AddUser(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Code == 0 {
		log.Printf("A new Customer has been added")
	}
}

func getUsers(client proto.UserServiceClient, filter *proto.UserFilter) {
	stream, err := client.GetUsers(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers: %v", err)
	}
	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}
		log.Printf("Customer: %v", customer)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("error")
	}
	defer conn.Close()

	client := proto.NewUserServiceClient(conn)

	user := &proto.User{
		Id:    10,
		Name:  "allen",
		Email: "allen@qq.com",
		Phone: "18111222211",
		Parents: []*proto.User_Parent{
			&proto.User_Parent{
				Name:     "zhangsan",
				Email:    "zhangsan@qq.com",
				Phone:    "18111222211",
				Relation: "father",
			},
			&proto.User_Parent{
				Name:     "lisi",
				Email:    "lisi@qq.com",
				Phone:    "18111222211",
				Relation: "mother",
			},
		},
	}

	addUser(client, user)

	user = &proto.User{
		Id:    11,
		Name:  "Jack",
		Email: "jack@qq.com",
		Phone: "13111222211",
		Parents: []*proto.User_Parent{
			&proto.User_Parent{
				Name:     "wangwu",
				Email:    "wangwu@qq.com",
				Phone:    "18111222211",
				Relation: "father",
			},
		},
	}

	addUser(client, user)


	filter := &proto.UserFilter{Id: 11}
	getUsers(client, filter)

}