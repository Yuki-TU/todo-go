package utils

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// パスワードをハッシュ化する
// @params password ハッシュ化したいパスワード
//
// @return ハッシュ化されたパスワード
func HashPassword(password string) string {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return string(hashPassword)
}

// パスワードが一意するか判断する
// @params hashPassword ハッシュ化されたパスワード
//
// @parasm password 比較対象のハッシュ化されていないパスワード
//
// @return 一致：nil, 一致しない: エラー
func ComparePassword(hashPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err
}
