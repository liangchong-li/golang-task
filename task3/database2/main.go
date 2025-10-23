package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Account struct {
	ID      int
	Balance int
}
type Transaction struct {
	ID            int
	FromAccountId int
	ToAccountId   int
	Amount        int
}

//题目2：事务语句
//假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和
//transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
//要求 ：
//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
//在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
//如果余额不足，则回滚事务。

func main() {
	db, err := sql.Open("mysql", "root:admin@tcp(192.168.2.2:3306)/gorm?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	// 初始化两个余额为0的账户
	//initAccount(db)

	// 向账户A中写入余额：50
	writeBalance(db, 1, 50)

	// 转账前
	accountA := Account{
		ID: 1,
	}
	accountB := Account{
		ID: 2,
	}
	selectAccount(db, &accountA)
	selectAccount(db, &accountB)

	fmt.Println("转账前，accountA: ", accountA)
	fmt.Println("转账前，accountB: ", accountB)

	amount := 100
	transactions(db, 1, 2, amount)

	// 转账后
	selectAccount(db, &accountA)
	selectAccount(db, &accountB)
	fmt.Println("转账后，accountA: ", accountA)
	fmt.Println("转账后，accountB: ", accountB)

	// 向账户A中写入余额：150
	writeBalance(db, 1, 150)

	transactions(db, 1, 2, amount)

	// 转账后
	selectAccount(db, &accountA)
	selectAccount(db, &accountB)
	fmt.Println("转账后，accountA: ", accountA)
	fmt.Println("转账后，accountB: ", accountB)

	defer db.Close()
}

func writeBalance(db *sql.DB, accountId int, balance int) {
	db.Exec("update accounts set balance = ? where id = ?", balance, accountId)
}

func selectAccount(db *sql.DB, account *Account) {
	db.QueryRow("select balance from accounts where id = ?", account.ID).Scan(&account.Balance)
	//fmt.Println("方法中:", account)
}

func initAccount(db *sql.DB) {
	db.Exec("insert into accounts(balance) values (?)", 0)
	db.Exec("insert into accounts(balance) values (?)", 0)
}

func transactions(db *sql.DB, accountA int, accountB int, amount int) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	var account Account
	err = tx.QueryRow("select balance from accounts where id = ?", accountA).Scan(&account.Balance)
	if err != nil {
		log.Fatal(err)
	}
	if account.Balance < 100 {
		fmt.Println("账户A余额为：", account.Balance, "。余额不足！")
		tx.Rollback()
		return
	}
	_, err = tx.Exec("update accounts set balance = balance + ? where id = ?", amount, accountB)
	if err != nil {
		tx.Rollback()
	}
	_, err = tx.Exec("update accounts set balance = balance - ? where id = ?", amount, accountA)
	if err != nil {
		tx.Rollback()
	}
	_, err = tx.Exec("insert into transactions(from_account_id, to_account_id, amount) values (?, ?, ?)", accountA, accountB, amount)
	if err != nil {
		tx.Rollback()
	}
	fmt.Println("账户A余额为：", account.Balance, "。转账成功")
	tx.Commit()
}
