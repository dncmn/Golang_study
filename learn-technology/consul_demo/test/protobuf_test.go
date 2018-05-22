package test

import (
	"fmt"
	proto "self/test/example/proto"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Println(proto.OrderStatus_ONSALE == 0)
}

func TestGetUserTitle(t *testing.T) {
	req := proto.UserTitleRequest{Uid: "123456"}

	title := &proto.GameUserTitle{
		TitleID:    5001,
		Progress:   -2,
		IsFinished: false,
		IsUsed:     true,
	}
	title2 := &proto.GameUserTitle{
		TitleID:    5002,
		Progress:   -3,
		IsFinished: true,
		IsUsed:     true,
	}

	res := proto.UserTitleReply{
		Code:   0,
		Status: true,
		Data:   []*proto.GameUserTitle{title, title2},
	}
	fmt.Println(req.Uid)
	for _, v := range res.Data {
		fmt.Printf("titleId=%v,progress=%v\n", v.TitleID, v.Progress)
	}

	fmt.Println(proto.OrderStatus_NOTSALE, proto.OrderStatus_SALEOUT)

}
