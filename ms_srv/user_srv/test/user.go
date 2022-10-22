package main

import (
	"context"
	"fmt"
	"time"

	"example.com/ms_srv/user_srv/proto"
	"google.golang.org/grpc"
)
var (
	userClient proto.UserClient
	connect *grpc.ClientConn
	err error
)

func init() {
	connect, err = grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	userClient = proto.NewUserClient(connect)
}

func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		PageNum: 2,
		PageSize: 5,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user)
		checkRsp, err := userClient.CheckPasswd(context.Background(), &proto.PasswdCheckInfo{
			Password: "admin123",
			EncryptedPassword: user.Password,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkRsp.Success)
	}
}

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("name%d", i),
			Mobile:   fmt.Sprintf("1375384451%d", i),
			Password: "admin123",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}

func TestGetUserByID() {
	rsp, err := userClient.GetUserByID(context.Background(), &proto.IDRequest{
		Id: 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func TestGetUserByMobile() {
	rsp, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: "13753844526",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func TestUpdateUser() {
	_, err := userClient.UpdateUser(context.Background(), &proto.UpdateUserInfo{
		Id: 1,
		NickName: "cloaks",
		Gender: "famale",
		Birthday: uint64(time.Now().Unix()),
	})
	if err != nil {
		panic(err)
	}
}



func main() {
	defer connect.Close()
	// TestGetUserList()
	// TestCreateUser()
	// TestGetUserByID()
	// TestUpdateUser()
	TestGetUserByMobile()
}