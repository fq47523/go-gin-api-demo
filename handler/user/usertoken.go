package user

import (
	myjwt "apidemo/router/middleware"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

func generateToken() (string ,error) {
	j := &myjwt.JWT{
		[]byte("woshifuqing"),
	}
	claims := myjwt.CustomClaims{
		"1",
		"admin",
		"13476155550",
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "woshifuqing",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	return token,err
}
