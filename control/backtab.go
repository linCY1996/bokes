package control

import (
	"boke/model"
	"boke/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const tokenKey = "this is a bdd"

//登陆页面
func ShowLogin(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile(`view/login.html`)
	w.Write(buf)
}

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var num = r.Form.Get(`num`)
	nums, _ := strconv.Atoi(num)
	var pass = r.Form.Get(`pass`)
	mod, err := model.Logins(nums)
	fmt.Println(err)
	if err != nil {
		w.Write([]byte(`信息错误`))
		return
	}
	if pass != mod.Pass {
		w.Write([]byte(`信息错误`))
		return
	}
	data := model.Jwt{
		Num:  mod.Num,
		Pass: mod.Pass,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}
	jwts := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	token, err := jwts.SignedString([]byte(tokenKey))
	// fmt.Println(err)
	if err != nil {
		w.Write(util.NewResult(util.Fail, "token生成失败，请重试"))
		return
	}
	w.Write(util.NewResult(util.Success, `登录成功`, token))
}

func Viewbacktab(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile(`view/backtab.html`)
	w.Write(buf)
}

//图文信息的msgs
func PostTwenmsgs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var html = r.FormValue(`html`)
	fmt.Println(html)
}
