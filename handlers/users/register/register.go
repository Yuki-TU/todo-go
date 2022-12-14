package users

import (
	"net/http"

	registerAuth "github.com/Yuki-TU/todo-go/controllers/user/register"
	util "github.com/Yuki-TU/todo-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service registerAuth.Service
}

// ファクトリー関数
func NewHandlerRegister(service registerAuth.Service) *handler {
	return &handler{service: service}
}

// ユーザ東麓処理を行う
// @params ctx ginコンテキスト
func (h *handler) RegisterHandler(ctx *gin.Context) {
	// POSTデータを変換JSONから構造体にマッピング
	var input registerAuth.InputRegister
	ctx.ShouldBindJSON(&input)

	// ユーザ登録
	user, errRegister := h.service.RegisterService(&input)

	switch errRegister {
	case "REGISTER_CONFLICT_409":
		util.APIResponse(ctx, "Email already exist", http.StatusConflict, http.MethodPost, nil)
		return
	case "REGISTER_FAILED_403":
		util.APIResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, nil)
		return
	default:
		accessTokenData := map[string]interface{}{"id": user.ID, "email": user.Email}
		accessToken, errToken := util.Sign(accessTokenData, util.GodotEnv("JWT_SECRET"), 60)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}
		util.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, map[string]string{"accessToken": accessToken})
	}
}
