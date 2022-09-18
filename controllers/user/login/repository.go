package loginAuth

import (
	"database/sql"

	model "github.com/Yuki-TU/todo-go/models"
	util "github.com/Yuki-TU/todo-go/utils"
)

type Repository interface {
	LoginRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *sql.DB
}

func NewRepositoryLogin(db *sql.DB) *repository {
	return &repository{db: db}
}

// ログイン確認処理
// @params input 入力した値(email, password)
//
// @return ユーザ情報, エラーコード
func (r *repository) LoginRepository(input *model.EntityUsers) (*model.EntityUsers, string) {
	errorCode := make(chan string, 1)

	var user model.EntityUsers

	row := r.db.QueryRow(`SELECT * FROM users where email = ? LIMIT 1`, input.Email)
	err := row.Scan(&user.ID, &user.Fullname, &user.Email, &user.Password, &user.Active, &user.UpdatedAt, &user.CreatedAt)

	// アカウントが見つからない
	if err != nil {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return input, <-errorCode
	}

	// パスワードが一致しない
	comparePasswordErr := util.ComparePassword(user.Password, input.Password)
	if comparePasswordErr != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return &user, <-errorCode
	}

	errorCode <- "nil"
	return &user, <-errorCode
}
