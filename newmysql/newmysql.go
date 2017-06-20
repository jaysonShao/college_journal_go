// college_journal project newmysql.go
package newmysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func newmysql() {
	db, err := sql.Open("mysql", "name:password@tcp(127.0.0.1:3306)/?timeout=90s&collation=utf8mb4_unicode_ci")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println(db.Stats())
	fmt.Println(db.Ping())
	db.Exec("create database college_journal DEFAULT CHARACTER SET utf8")
	db.Exec("use college_journal")
	db.Exec("create table information (id int not null auto_increment,title varchar(100) not null,school varchar(20) not null,createtime timestamp default '2017-01-01 00:00:00',userid int not null,usernickname varchar(18) not null,usertelephone char(11) not null,content varchar(5000) not null ,praise varchar(6) not null, primary key (id))ENGINE=InnoDB  DEFAULT CHARSET=utf8")
	db.Exec("create index information_userid_usertelephone on information(userid,usertelephone)")
	db.Exec("create index information_title on information(title)")
	db.Exec("create table suggest (id int not null auto_increment,content varchar(5000) not null,school varchar(20) not null,createtime timestamp default '2017-01-01 00:00:00',userid int not null,primary key(id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8")
	db.Exec("create table authority_management(id int not null auto_increment primary key,authority varchar(10) not null  default 'base',userid int not null,createtime timestamp default '2017-01-01 00:00:00')ENGINE=InnoDB  DEFAULT CHARSET=utf8")
	db.Exec("create table user (id smallint not null auto_increment,telephone char(11)not null,school varchar(50) not null,nickname varchar(18) not null,sex char(1) not null,signature varchar(60), createtime timestamp default '2017-01-01 00:00:00',password varchar(16) not null, authority varchar(10) not null , praise varchar(6) not null default '0', primary key(id))ENGINE=InnoDB  DEFAULT CHARSET=utf8")
	db.Exec("create index user_telephone_nickname on user(telephone,nickname)")
	fmt.Println("end of newmysql.go")
}
