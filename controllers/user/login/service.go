package loginAuth

import (
	model "github.com/Yuki-TU/todo-go/models"
)

type Service interface {
	LoginService(input *InputLogin) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

// 入力要素を整形し、リポジトリにつなげr
func (s *service) LoginService(input *InputLogin) (*model.EntityUsers, string) {
	user := model.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.LoginRepository(&user)

	return resultLogin, errLogin
}
