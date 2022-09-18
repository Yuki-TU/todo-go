package utils

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

type MetaToken struct {
	ID            string
	Email         string
	ExpiredAt     time.Time
	Authorization bool
}

type AccessToken struct {
	Claims MetaToken
}

// アクセストークンお発行する
//
// @params Clames Claimsで利用するデータ
//
// @params SecreublicKeyEnvName 署名作成に医療するシークレットキー
//
// @pramas expiredAt 有効期間(分)
//
// @returns accessToken, err アクセストークン, err
func Sign(Claims map[string]interface{}, SecretPublicKeyEnvName string, ExpiredAt time.Duration) (string, error) {
	jwtSecretKey := GodotEnv(SecretPublicKeyEnvName)

	// clamsの作成
	claims := jwt.MapClaims{}
	expiredAt := time.Now().Add(time.Duration(time.Minute) * ExpiredAt).Unix()
	claims["exp"] = expiredAt
	claims["authorization"] = true
	for index, claim := range Claims {
		claims[index] = claim
	}

	// ヘッダーとペイロードの作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		logrus.Error(err.Error())
		return accessToken, err
	}

	return accessToken, nil
}

// リクエストヘッダーよりトークンを検証、解析する
//
// @paramas ctx ginコンテキスト
//
// @params SecretPublicKeyEnvName envに指定されたシークレットキー
//
// @return 解析されたトークン情報
func VerifyTokenHeader(ctx *gin.Context, SecretPublicKeyEnvName string) (*jwt.Token, error) {
	tokenHeader := ctx.GetHeader("Authorization")
	accessToken := strings.SplitAfter(tokenHeader, "Bearer")[1]
	jwtSecretKey := GodotEnv(SecretPublicKeyEnvName)

	token, err := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}
