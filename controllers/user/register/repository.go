package registerAuth

import (
	"database/sql"

	model "github.com/Yuki-TU/todo-go/models"
)

type Repository interface {
	RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	// DBハンドラーインスタンス
	db *sql.DB
}

func NewRepositoryRegister(db *sql.DB) *repository {
	return &repository{db: db}
}

// DBに対してでユーザ登録登録を行う
// @params input 登録情報
func (r *repository) RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string) {
	errorCode := make(chan string, 1)

	// 重複するメールがあれば、490エラー
	row := r.db.QueryRow(`SELECT COUNT(*) FROM users where email = ? LIMIT 1`, input.Email)
	count := 0
	row.Scan(&count)
	if count == 1 {
		errorCode <- "REGISTER_CONFLICT_409"
		return input, <-errorCode
	}

	// ユーザ登録
	_, err := r.db.Exec(`INSERT INTO users (fullname, email, password, active, updatedAt, createdAt) VALUES(?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`,
		input.Fullname, input.Email, input.Password, false)
	if err != nil {
		errorCode <- "REGISTER_FAILED_403"
		return input, <-errorCode
	}
	errorCode <- "nil"
	return input, <-errorCode
}
