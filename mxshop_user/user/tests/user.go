package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mxshop/user/proto"
	"strconv"
)

var userClient proto.UserClient
var coon *grpc.ClientConn

func Init() {
	var err error
	coon, err = grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(coon)
}

func TestCreatUserlist() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreatUser(context.Background(), &proto.CreatUserInfo{
			Password: "admin",
			Nickname: "bobby" + strconv.Itoa(i),
			Mobile:   "1786729698" + strconv.Itoa(i),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}

}

func TestGetUserMobile() {
	rsp, err := userClient.GetUserMobile(context.Background(), &proto.MobileRequest{
		Mobile: "17867296981"})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Id, rsp.Mobile, rsp.Nickname, rsp.Gender)

}

func TestUpdateUser() {
	req, err := userClient.UpdateUser(context.Background(), &proto.UpdateUserInfo{
		Id: 1,
	})
	if err != nil {
		panic(err)
	}
	req.String()

}

func main() {
	Init()
	TestCreatUserlist()
	_ = coon.Close()

}
