package models

import (
	"time"

	util "github.com/Yuki-TU/todo-go/utils"
)

// データベースで扱っているユーザテーブル
type EntityUsers struct {
	ID        int    `gorm:"primaryKey;"`
	Fullname  string `gorm:"type:varchar(255);unique;not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Active    bool   `gorm:"type:bool;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// アカウント作成前にする処理
// * パスワードハッシュ化
func (entity *EntityUsers) HandleBeforeCreateUser() error {
	entity.Password = util.HashPassword(entity.Password)
	return nil
}
