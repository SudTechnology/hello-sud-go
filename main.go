package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", Login)
	r.POST("/get_sstoken", GetSsToken)
	r.POST("/update_sstoken", UpdateSSToken)
	r.POST("/get_user_info", GetUserInfo)
	r.POST("/report_game_info", ReportGameInfo)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
