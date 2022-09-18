package userList

import (
	"net/http"

	userList "github.com/Yuki-TU/todo-go/controllers/user/getList"
	util "github.com/Yuki-TU/todo-go/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service userList.Service
}

func NewHandlerUserList(service userList.Service) *handler {
	return &handler{service: service}
}

// ユーザ情報一覧取得
// @params ctx コンテキスト
func (h *handler) GetUserListHandler(ctx *gin.Context) {
	// Authorizationヘッダーの確認
	if ctx.GetHeader("Authorization") == "" {
		util.APIResponse(ctx, "Authorization is required for this endpoint", http.StatusUnauthorized, http.MethodGet, nil)
		return
	}

	// JWT認証確認
	_, err := util.VerifyTokenHeader(ctx, "JWT_SECRET")
	if err != nil {
		util.APIResponse(ctx, "Verified activation token failed", http.StatusUnauthorized, http.MethodGet, nil)
		return
	}

	// ユーザ一覧取得
	userList, errUserList := h.service.GetUserListService()
	switch errUserList {
	case "USERLIST_NOT_FOUND_404":
		util.APIResponse(ctx, "user list is not found", http.StatusNotFound, http.MethodGet, nil)
		return
	default:
		util.APIResponse(ctx, "successfully", http.StatusOK, http.MethodGet, userList)
	}
}
