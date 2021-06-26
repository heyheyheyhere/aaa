package main

import (
	"database/sql"
	"fmt"

	"log"
	"net/smtp"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jordan-wright/email"
)

const (
	USERNAME = "root"
	PASSWORD = ""
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "demo"
)

func CreateTable(db *sql.DB) error {
	sql := `CREATE TABLE IF NOT EXISTS users(
	id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64)
	); `

	if _, err := db.Exec(sql); err != nil {
		fmt.Println("建立 Table 發生錯誤:", err)
		return err
	}
	fmt.Println("建立 Table 成功！")
	return nil
}
func InsertUser(DB *sql.DB, username, password string) error {
	_, err := DB.Exec("insert INTO users(username,password) values(?,?)", username, password)
	if err != nil {
		fmt.Printf("建立使用者失敗，原因是：%v", err)
		return err
	}
	fmt.Println("建立使用者成功！")
	return nil
}

type User struct {
	ID       string
	Username string
	Password string
}

func QueryUser(db *sql.DB, username string) {
	user := new(User)
	row := db.QueryRow("select * from users where username=?", username)
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		//InsertUser(db, "1310931031", "s1310931031@nutc.edu.tw")
		mail(db, "1310931031")
		fmt.Printf("註冊成功")
		return
	} else {
		fmt.Println("已註冊過了")
	}

}
func mail(db *sql.DB, username string) {
	//user := new(User)
	sql := fmt.Sprintf("select password from users where username=?", username)
	//row := db.("select * from users where username=?", mail)
	//if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
	//fmt.Printf("映射使用者失敗，原因為：%v\n", err)
	//return
	//}
	//fmt.Println("查詢使用者成功", *user)
	e := email.NewEmail()
	e.From = "<jim887576@gmail.com>"
	//mail := []string{"s1310931031@nutc.edu.tw"}
	e.To = []string{sql}
	//e.Cc = []string{"test1@126.com", "test2@126.com"}
	//e.Bcc = []string{"secret@126.com"}
	e.Subject = "嘿嘿嘿在這兒"
	e.Text = []byte("黑在黑在我是柏睿")
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "jim887576@gmail.com", "bbjgrxgvsqwiiatg", "smtp.gmail.com"))
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("開啟 MySQL 連線發生錯誤，原因為：", err)
		return
	}
	if err := db.Ping(); err != nil {
		fmt.Println("資料庫連線錯誤，原因為：", err.Error())
		return
	}

	defer db.Close()
	CreateTable(db)
	//InsertUser(db, "1310931031", "s1310931031@nutc.edu.tw")
	QueryUser(db, "1310931031")
}
