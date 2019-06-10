package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-start/dao"
	"go-start/dl"
	"os"
)

func init1()  {
	container := dl.NewContainer()
	db, err := sql.Open("mysql", "root:root@tcp(localhost)/sampledb")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
	container.SetSingleton("db", db)
	container.SetPrototype("b", func() (interface{}, error) {
		return dao.NewB(), nil
	})

	a := dao.NewA()
	if err := container.Ensure(a); err != nil {
		fmt.Println(err)
		return
	}
	// 打印指针，确保单例和实例的指针地址
	fmt.Printf("db: %p\ndb1: %p\nb: %p\nb1: %p\n", a.Db, a.Db1, &a.B, &a.B1)
}

func init2()  {
	container := dl.NewContainer()
	db, err := sql.Open("mysql", "root:root@tcp(localhost)/sampledb")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
	container.SetSingleton("db", db)
	container.SetPrototype("b", func() (interface{}, error) {
		return dao.NewB(), nil
	})

	a := dao.NewA()
	if err := container.Ensure(a); err != nil {
		fmt.Println(err)
		return
	}
	// 打印指针，确保单例和实例的指针地址
	fmt.Printf("db: %p\ndb1: %p\nb: %p\nb1: %p\n", a.Db, a.Db1, &a.B, &a.B1)
}

func main() {
	init1()
	init2()
}
