package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostInfo(w http.ResponseWriter, r *http.Request) {
	var info PostInfo_
	result, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(result))
	json.Unmarshal(result, &info)
	if len(info.Title) > 0 && len(info.School) > 0 && len(info.Usernickname) != 0 && len(info.Usertelephone) == 11 && len(info.Content) >= 5 && info.Praise == "0" {
		stmt, _ := db.Prepare("INSERT INTO information VALUES(?,?,?,now(),?,?,?,?,?)")
		if _, err := stmt.Exec(0, info.Title, info.School, info.Userid, info.Usernickname, info.Usertelephone, info.Content, info.Praise); err != nil {
			panic(err)
		}
		defer stmt.Close()
		w.WriteHeader(403);
		fmt.Fprintf(w, "发布校刊成功！")

	} else {
		w.WriteHeader(403)
		fmt.Fprintf(w, "信息不完整！")
	}
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
	var info Getinfo_
	var count int
	db.QueryRow("select count(*) from information").Scan(&count)
	info.Count = count
	result, _ := db.Query("SELECT * FROM information order by id desc")
	defer result.Close()
	for result.Next() {
		var in PostInfo_
		result.Scan(&in.Id, &in.Title, &in.School, &in.Createtime, &in.Userid, &in.Usernickname, &in.Usertelephone, &in.Content, &in.Praise)
		info.Data = append(info.Data, in)
	}
	response, _ := json.Marshal(&info)
	fmt.Fprintf(w, string(response))


}

func Postpraise(w http.ResponseWriter, r *http.Request){
	var data Postpraise_
	resulte, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(resulte))
	h := r.Header.Get("table")
	json.Unmarshal(resulte, &data)
	if h == "information" {
		stmt, _ := db.Prepare("UPDATE information SET praise = ? WHERE id = ?")
		stmt.Exec(data.Praise, data.Id)
		fmt.Println("s")
	}else if h == "user" {
		stmt, _ :=db.Prepare("UPDATE user SET praise = ? WHERE  id = ?")
		stmt.Exec(data.Praise, data.Id)
		fmt.Println("b")
	}else {
		w.WriteHeader(403)
		fmt.Fprintf(w, "对不起你没有权限！")
	}
	w.WriteHeader(403)
	fmt.Fprintf(w, "赞")

}

