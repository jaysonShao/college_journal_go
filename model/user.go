package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	var js Postuser
	result, _ := ioutil.ReadAll(r.Body)

	fmt.Println(string(result))
	json.Unmarshal(result, &js)
	var check = PostUser_mysql
	stmt, err := db.Query("select telephone from user where telephone=? ", js.Telephone)
	defer stmt.Close()
	if err != nil {
		panic(err)
	}
	for stmt.Next() {
		stmt.Scan(&check.telephone)
	}
	if js.Telephone == check.telephone {
		w.WriteHeader(403)
		fmt.Fprintf(w, "对不起，您的手机已经注册！")
	} else {
		db_user_p, _ := db.Prepare("INSERT INTO user VALUES(?,?,?,?,?,?,now(),?,?,?)")
		_, err := db_user_p.Exec(0, js.Telephone, js.School, js.Nickname, js.Sex, js.Signature, js.Password, "normal", js.Praise)
		defer db_user_p.Close()
		if err != nil {
			panic(err)
		}
		w.WriteHeader(403)
		fmt.Fprintf(w, "注册成功！")
	}
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
	var js = GetUser_
	var check Postuser
	result, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(result, &js)
	stmt, _ := db.Prepare("SELECT * FROM user where telephone=?")
	res, _ := stmt.Query(js.Telephone)
	defer res.Close()
	for res.Next() {

		res.Scan(&check.Id, &check.Telephone, &check.School, &check.Nickname, &check.Sex, &check.Signature, &check.Createtime, &check.Password, &check.Authority, &check.Praise)
	}
	if js.Telephone == check.Telephone {
		if js.Password == check.Password {
			w.WriteHeader(200)
			response, _ := json.Marshal(&check)
			fmt.Println(string(response))
			fmt.Fprintf(w, string(response))
		} else {
			w.WriteHeader(403)
			fmt.Fprintf(w, "密码错误，请重试！")
		}
	} else {
		w.WriteHeader(403)
		fmt.Fprintf(w, "系统没有您的信息，请先注册!")
	}
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	var js Postuser
	result, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(result))

	json.Unmarshal(result, &js)
	stmt, _ := db.Prepare("UPDATE user set telephone=?, school=?, nickname=?, sex=? where id=?")
	stmt.Exec(js.Telephone, js.School, js.Nickname, js.Sex, js.Id)
	defer stmt.Close()
	w.WriteHeader(403)
	fmt.Fprintf(w, "资料已经更新")
}

func Putuser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get(":id")
	fmt.Println(id)
	var idd int
	idd, _ = strconv.Atoi(id)
	var sign = Putsign
	result, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(result))
	json.Unmarshal(result, &sign)
	stmt, _ := db.Prepare("update user set signature=? where id=?")
	stmt.Exec(sign.Signature, idd)
	defer stmt.Close()
	w.WriteHeader(403)
	fmt.Fprintf(w, "签名更新成功！")
}

func PutPwd(w http.ResponseWriter, r *http.Request) {
	pwd := r.URL.Query().Get(":param")
	id := r.URL.Query().Get(":id")
	var idd int
	idd, _ = strconv.Atoi(id)
	fmt.Println(pwd, idd)
	stmt, _ := db.Prepare("update user set password=? where id=?")
	stmt.Exec(pwd, idd)
	defer stmt.Close()
	w.WriteHeader(403)
	fmt.Fprintf(w, "密码更新成功！")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	hander := r.Header.Get("authority")
	fmt.Println(hander)
	if hander == "normal" {
		var userinfo Getuserinfo
		result, err := db.Query("SELECT * FROM user")
		if err != nil {
			panic(err)
		}
		defer result.Close()
		for result.Next() {
			var info Getuserin
			result.Scan(&info.Id, &info.Telephone, &info.School, &info.Nickname, &info.Sex, &info.Signature, &info.Createtime, &info.Password, &info.Authority, &info.Praise)
			info.Createtime = "0"
			info.Password = "0"
			info.Authority = "0"
			userinfo.Data = append(userinfo.Data, info)
		}
		response, _ := json.Marshal(&userinfo)
		fmt.Fprintf(w, string(response))

	} else {
		w.WriteHeader(403)
		fmt.Fprintf(w, "要先注册了才可以看见校友哦！")
	}
}
