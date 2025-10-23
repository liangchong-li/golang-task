package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Student struct {
	Id     int
	Name   string
	Age    int
	Gender string
}

//题目1：基本CRUD操作
//假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
//要求 ：
//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

func main() {
	db, err := sql.Open("mysql", "root:admin@tcp(192.168.2.2:3306)/gorm?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	//err = db.Ping()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	//result, err := db.Exec("insert into students (name, age, grade) values (?, ?, ?)", "张三", 20, "三年级")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//id, err := result.LastInsertId()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//affected, err := result.RowsAffected()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("result of id:", id)
	//fmt.Println("result of affected:", affected)

	// 查询 students 表中所有年龄大于 18 岁的学生信息。
	//rows, err := db.Query("select * from students where age > ?", 18)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var student Student
	//for rows.Next() {
	//	//rows.Scan(student)
	//	err := rows.Scan(&student.Id, &student.Name, &student.Age, &student.Gender)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(student)
	//}

	// 将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	//upd, err := db.Exec("update students set grade = ? where name = ?", "四年级", "张三")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//id, err := upd.LastInsertId()
	//affected, err := upd.RowsAffected()
	//fmt.Println(id, affected)

	// 删除 students 表中年龄小于 15 岁的学生记录。
	del, err := db.Exec("delete from students where age < ?", 15)
	if err != nil {
		log.Fatal(err)
	}
	id, err := del.LastInsertId()
	affected, err := del.RowsAffected()
	fmt.Println(id, affected)

	defer db.Close()
}
