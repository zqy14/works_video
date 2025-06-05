package server

import (
	"context"
	"fmt"
	"video1/video-rpc/global"
	"video1/video-rpc/handler/model"
	__ "video1/video-rpc/proto"
)

type Server struct {
	__.UnimplementedUserServer
}

func (s *Server) Login(_ context.Context, in *__.LoginRequest) (*__.LoginRepine, error) {
	var user model.User

	err := global.DB.Where("mobile = ?", in.Mobile).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("用户存在")
	}
	return &__.LoginRepine{
		Id: user.Id,
	}, nil
}
