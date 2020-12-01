package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func main() {
	GetStudentByID("123")
}

//service
func GetStudentByID(id string) (*Student, error) {
	stu, err := getStudent(id)
	//根据业务对err进行区分处理
	if err == sql.ErrNoRows {
		return nil, err
	}
	return stu, nil
}

//dao
type Student struct {
	ID   string
	Name string
}

func getStudent(id string) (*Student, error) {
	db, err := sql.Open("mysql", "root:123456@/test")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	user := new(Student)

	err = db.QueryRow("SELECT id,name FROM students WHERE id = ?", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil
}
