package model

import jwt "github.com/dgrijalva/jwt-go"

type Login struct {
	ID   int    `json:"id" xml:"id"`
	Num  int    `json:"num" xml:"num"`
	Pass string `json:"pass" xml:"pass"`
}

// Jwt json web token
type Jwt struct {
	Num  int
	Pass string
	jwt.StandardClaims
}

func Logins(num int) (Login, error) {
	mod := Login{}
	err := DB.Get(&mod, `select * from login where num=?`, num)
	return mod, err
}
