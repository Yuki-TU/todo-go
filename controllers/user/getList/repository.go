package getlist

import (
	"database/sql"
	"log"

	model "github.com/Yuki-TU/todo-go/models"
)

type Repository interface {
	UsersRepository() (*[]model.EntityUsers, string)
}

type repository struct {
	// DBハンドラーインスタンス
	db *sql.DB
}

func NewRepositoryUserList(db *sql.DB) *repository {
	return &repository{db: db}
}

// DBに対してでユーザ登録登録を行う
// @params input 登録情報
func (r *repository) UsersRepository() (*[]model.EntityUsers, string) {
	errorCode := make(chan string, 1)

	var userList []model.EntityUsers
	rows, err := r.db.Query(`SELECT id, fullname, email, active, createdAt, updatedAt FROM users`)
	for rows.Next() {
		user := model.EntityUsers{}
		if err := rows.Scan(&user.ID, &user.Fullname, &user.Email, &user.Active, &user.UpdatedAt, &user.CreatedAt); err != nil {
			log.Fatal(err)
		}
		userList = append(userList, user)
	}

	// ユーザ登録
	if err != nil {
		errorCode <- "USERLIST_NOT_FOUND_404"
		return &userList, <-errorCode
	}
	errorCode <- "nil"
	return &userList, <-errorCode
}
