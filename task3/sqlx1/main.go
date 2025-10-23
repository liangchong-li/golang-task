package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

//题目1：使用SQL扩展库进行查询
//假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
//要求 ：
//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

type Employees struct {
	Id         int
	Name       string
	Department string
	Salary     int
}

func main() {
	db, err := sqlx.Open("mysql", "root:admin@tcp(192.168.2.2:3306)/gorm?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	//err = db.Ping()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//db.MustExec("insert into employees(name, department, salary) values(?,?,?)", "张三", "技术部", 15000)
	//db.MustExec("insert into employees(name, department, salary) values(?,?,?)", "李四", "技术部", 16000)
	//db.MustExec("insert into employees(name, department, salary) values(?,?,?)", "王五", "技术部", 17000)
	//db.MustExec("insert into employees(name, department, salary) values(?,?,?)", "赵六", "产品部", 8000)
	//db.MustExec("insert into employees(name, department, salary) values(?,?,?)", "孙七", "人事部", 7000)

	//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	var employees []Employees
	err = db.Select(&employees, "select * from employees where department = ?", "技术部")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(employees)

	//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	var employeeTopOfSalary Employees
	err = db.Get(&employeeTopOfSalary, "select * from employees order by salary desc limit 1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(employeeTopOfSalary)
}
