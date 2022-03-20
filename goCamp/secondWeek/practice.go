package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

//mysql init
const name = "root:123456@tcp(localhost:3306)/mysql"

func main() {
	user, err := GetUserById(110)
	//handle err
	if err != nil {
		fmt.Printf("original error : %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace : \n %+v\n", err)
		return
	}

	//business handle, 如果不处理，则会导致给当前人员分配任务，可能当前人员已被删除
	fmt.Println("hi ", user.name)
	AssignedTasks(user)
}

func AssignedTasks(user *User) {
	user.do("design db")
}

type User struct {
	name string
}

func (u User) do(work string) {
	fmt.Println(work, u.name)
}

// GetUserById 通过id获取人员
func GetUserById(id int) (*User, error) {
	user := User{}
	open, err := sql.Open("mysql", name)
	defer open.Close()
	if err != nil {
		return nil, errors.Wrapf(err, "open")
	}
	row := open.QueryRow(fmt.Sprintf("select name from userInfo where name = %d", id))
	err = row.Scan(&user.name)
	if err != nil {
		return nil, errors.Wrapf(err, "not exist name of moonus")
	}
	//if not handle err, that will return zero value
	return &user, nil
}
