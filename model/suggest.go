package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostSuggest(w http.ResponseWriter, r *http.Request) {
	var suggest Suggest
	result, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(result))
	json.Unmarshal(result, &suggest)
	stmt, _ := db.Prepare("INSERT INTO suggest VALUES(?,?,?,now(),?)")
	stmt.Exec("0", suggest.Content, suggest.school,suggest.Userid)
	w.WriteHeader(403)
	fmt.Fprintf(w, "感谢您的意见，我们会更加努力！")
}
