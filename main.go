// college_journal project main.go
package main

import (
	"college_journal/model"
	"fmt"
	"net/http"

	"github.com/drone/routes"
)

func main() {
	mux := routes.New()
	//	注册 同时更新权限表
	mux.Post("/user/", model.PostUser)
	//	登录
	mux.Post("/login/", model.GetLogin)
	//改签名
	mux.Put("/user/:id", model.Putuser)
	//	更新用户资料
	mux.Put("/user/", model.PutUser)
	//改密码
	mux.Put("/pwd/:param/:id", model.PutPwd)
	//	发布信息
	mux.Post("/info/", model.PostInfo)
	//	获取信息
	mux.Get("/info/", model.GetInfo)
	//	获取用户信息
	mux.Get("/user/", model.GetUser)
	//	反馈建议
	mux.Post("/suggest/", model.PostSuggest)
	//点赞
	mux.Post("/praise/", model.Postpraise)

	http.Handle("/", mux)
	http.ListenAndServe(":8081", nil)
	fmt.Println("END OF SERVERS")
	model.CloseDb()
}
