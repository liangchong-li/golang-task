package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

//题目2：实现类型安全映射
//假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
//要求 ：
//定义一个 Book 结构体，包含与 books 表对应的字段。
//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

type Book struct {
	Id     int
	Title  string
	Author string
	Price  int
}

func main() {
	db, err := sqlx.Open("mysql", "root:admin@tcp(192.168.2.2:3306)/gorm?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	//db.MustExec("insert into books(title, author, price) values(?,?,?)", "天龙八部", "金庸", 20)
	//db.MustExec("insert into books(title, author, price) values(?,?,?)", "丰乳肥臀", "莫言", 50)
	//db.MustExec("insert into books(title, author, price) values(?,?,?)", "呐喊", "鲁迅", 60)
	//db.MustExec("insert into books(title, author, price) values(?,?,?)", "诡秘之主", "爱潜水的乌贼", 30)
	//db.MustExec("insert into books(title, author, price) values(?,?,?)", "资本论", "马克思", 80)

	var books []Book
	err = db.Select(&books, "select * from books where price > ?", 50)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(books)
}
