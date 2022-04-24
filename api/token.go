package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/pkg/util"
)

func CreateToken(ctx *gin.Context) {
	token, err := util.GenerateToken("123", "zhoulei")

	fmt.Printf("%v %v", token, err)

	if err != nil {
	}
}

func CheckToken(ctx *gin.Context) {

	if userId, _ := ctx.Get("user_id"); userId != nil {
		fmt.Println(userId)
	}

	if userName, _ := ctx.Get("username"); userName != nil {
		fmt.Println(userName)
	}
}

func UpdateToken(ctx *gin.Context) {
	fmt.Println(util.RefreshToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzIiwidXNlcm5hbWUiOiJ6aG91bGVpIiwiZXhwIjoxNjUwMzM1NjYxLCJpYXQiOjE2NTAyNDkyNjEsImlzcyI6IkxvbmdUdUdhbWUifQ.iNTfzmVpCTWh31Bt14t_OwLmoe6K69lo20kCZune49k"))
}
