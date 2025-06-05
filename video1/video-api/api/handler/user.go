package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"video1/video-api/api/request"
	__ "video1/video-api/basic/proto"
	"video1/video-api/sdk"
)

func Login(c *gin.Context) {
	var req request.LoginUser
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "参数验证",
			"data": err.Error(),
		})
		return
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	C := __.NewUserClient(conn)

	login, err := C.Login(c, &__.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "登录失败",
			"data": nil,
		})
		return
	}

	token, _ := sdk.NewJWT("2211a").CreateToken(sdk.CustomClaims{
		ID:             uint(login.Id),
		StandardClaims: jwt.StandardClaims{},
	})

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "登录成功",
		"data": token,
	})
}
