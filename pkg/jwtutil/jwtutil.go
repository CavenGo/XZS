package jwtutil

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个UserID字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserName string `json:"userName"`
	jwt.StandardClaims
}

var mySecret = []byte("9df462be0b4bbf4743a00cb7b2b134ed")

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

const TokenExpireDuration = time.Hour * 24 * 7

// GenToken 生成access token 和 refresh token
func GenToken(userName string) (aToken, rToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userName, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "xzs",                                      // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)

	// refresh token 不需要存任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(), // 过期时间
		Issuer:    "bluebell",                              // 签发人
	}).SignedString(mySecret)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	// 解析token
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid { // 校验token
		err = errors.New("invalid token")
	}
	return
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token无效直接返回
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}

	// 从旧access token中解析出claims数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	// 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserName)
	}
	return
}
/*
该代码实现了JWT（JSON Web Token）的生成、解析和刷新功能。它引用了以下包：

"github.com/dgrijalva/jwt-go"：一个Go语言的JWT库，提供了JWT的生成、解析和验证等功能。
该代码包含了以下三个函数：

GenToken
该函数用于生成AccessToken和RefreshToken，实现过程如下：

创建一个自定义的MyClaims结构体，用于存储JWT中的自定义字段（此处为userName）和标准字段（如过期时间和签发人）。

使用jwt.NewWithClaims创建一个新的JWT Token，并将自定义的MyClaims结构体作为声明参数传递给它。

使用jwt.SigningMethodHS256对JWT进行签名，并将签名后的字符串作为AccessToken返回。

创建一个新的RefreshToken，与AccessToken不同的是，RefreshToken不需要保存任何自定义数据。

ParseToken

该函数用于解析AccessToken，实现过程如下：

使用jwt.ParseWithClaims解析AccessToken，并将自定义的MyClaims结构体作为声明参数传递给它。

如果解析出错，则返回错误信息；否则，返回解析出的自定义声明。

RefreshToken

该函数用于刷新AccessToken，实现过程如下：

使用jwt.Parse解析RefreshToken，如果解析出错，则直接返回错误信息。
从旧的AccessToken中解析出自定义声明。
如果AccessToken已过期，则创建一个新的AccessToken并返回；否则，直接返回旧的AccessToken和RefreshToken。
*/