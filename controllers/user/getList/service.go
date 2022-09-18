package getlist

import (
	model "github.com/Yuki-TU/todo-go/models"
)

type Service interface {
	GetUserListService() (*[]model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceRegister(repository Repository) *service {
	return &service{repository: repository}
}

// ハンドラーとリポジトリをつなげる
// @params input POSTデータ
//
// @return resultRegister 登録結果
//
// @return errRegister 登録の際発生したエラー
func (s *service) GetUserListService() (*[]model.EntityUsers, string) {
	resultRegister, errRegister := s.repository.UsersRepository()
	return resultRegister, errRegister
}
