package registerAuth

import (
	model "github.com/Yuki-TU/todo-go/models"
)

type Service interface {
	RegisterService(input *InputRegister) (*model.EntityUsers, string)
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
func (s *service) RegisterService(input *InputRegister) (*model.EntityUsers, string) {
	// 入力データをきれいにして登録処理を行う
	users := model.EntityUsers{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
	}

	users.HandleBeforeCreateUser()

	resultRegister, errRegister := s.repository.RegisterRepository(&users)
	return resultRegister, errRegister
}
