package model

type Postuser struct {
	Id         int    `json:"id"`
	Telephone  string `json:"telephone"`
	School    string `json:"school"`
	Nickname   string `json:"nickname"`
	Sex        string `json:"sex"`
	Signature  string `json:"signature"`
	Createtime string `json:"createtime"`
	Password   string `json:"password"`
	Authority  string `json:"authority"`
	Praise     string `json:"praise"`
}

var PostUser_mysql struct {
	id         int
	telephone  string
	school   string
	nickname   string
	sex        string
	signature  string
	createtime string
	password   string
	authority  string
	praise     string
}

var GetUser_ struct {
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

var Putsign struct {
	Signature string `json:"signature"`
}

type Getinfo_ struct {
	Count int         `json:"count"`
	Data  []PostInfo_ `json:"data"`
}

type PostInfo_ struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	School        string `json:"school"`
	Createtime    string `json:"createtime"`
	Userid        int    `json:"userid"`
	Usernickname  string `json:"usernickname"`
	Usertelephone string `json:"usertelephone"`
	Content       string `json:"content"`
	Praise        string `json:"praise"`
}

type Suggest struct {
	Id         int    `json:"id"`
	Content    string `json:"content"`
	Createtime string `json:"createtime"`
	Userid     int    `json:"userid"`
	school string `json:"school"`
}

type Getuserinfo struct {
	Count int         `json:"count"`
	Data  []Getuserin `json:"data"`
}

type Getuserin struct {
	Id         int    `json:"id"`
	Telephone  string `json:"telephone"`
	School     string `json:"school"`
	Nickname   string `json:"nickname"`
	Sex        string `json:"sex"`
	Signature  string `json:"signature"`
	Createtime string `json:"createtime"`
	Password   string `json:"password"`
	Praise     string `json:"praise"`
	Authority  string `json:"authority"`
}

type Postpraise_ struct {
	Id int `json:"id"`
	Praise string `json:"praise"`
}
