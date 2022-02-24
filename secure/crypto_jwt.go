// Copyright (c) 2019-2029 DY All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package secure

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/wengoldx/wing/invar"
	"strings"
	"time"
)

// Claims jwt claims data
type Claims struct {
	Keyword string `json:"keyword"`
	jwt.StandardClaims
}

// Generate a jwt token with keyword and salt string,
// the token will expired after the given duration
func GenJwtToken(keyword, salt string, dur time.Duration) (string, error) {
	expireAt := time.Now().Add(dur).Unix()
	claims := Claims{
		keyword,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    keyword,
		},
	}

	// create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signs the token with a salt.
	signedToken, err := token.SignedString([]byte(salt))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// Verify the encoded jwt token witch salt string
func ViaJwtToken(signedToken, salt string) (string, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(salt), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Keyword, nil
	}
	return "", err
}

// Encode account uuid and optianl datas as claims content of jwt token
func EncClaims(uuid string, params ...string) string {
	sets := []string{uuid}
	if len(params) > 0 {
		sets = append(sets, params...)
	}
	orikey := strings.Join(sets, ";")
	return EncodeBase64(orikey)
}

// Decode claims of jwt token and return datas as string array
func DecClaims(keyword string, count ...int) ([]string, error) {
	orikeys, err := DecodeBase64(keyword)
	if err != nil {
		return nil, err
	}

	sets := strings.Split(orikeys, ";")

	// check claims content fields if give the verify count param
	if cl := len(count); cl > 0 && count[0] > 0 && count[0] != len(sets) {
		return nil, invar.ErrInvalidNum
	}
	return sets, nil
}

// Encode account uuid, password and subject string
//
// `DEPRECATED`:
//
// Please use EncClaims instead
func EncJwtKeyword(uuid, pwd string, subject string) string {
	sets := []string{uuid, pwd, subject}
	orikey := strings.Join(sets, ";")
	return EncodeBase64(orikey)
}

// Decode account uuid, password and subject from jwt keyword string
//
// `DEPRECATED`:
//
// Please use DecClaims instead
func DecJwtKeyword(keyword string) (string, string, string) {
	orikeys, err := DecodeBase64(keyword)
	if err != nil {
		return "", "", ""
	}

	sets := strings.Split(orikeys, ";")
	for i := len(sets); i < 3; i++ {
		sets = append(sets, "")
	}
	return sets[0], sets[1], sets[2]
}
