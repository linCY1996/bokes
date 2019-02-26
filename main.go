package main

import (
	"boke/control"
	"boke/model"
	"boke/util"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

const tokenKey = "this is a bdd"

//中间件
func Mid(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		token := r.Form.Get("token")
		if token == "" {
			w.Write(util.NewResult(301, "没有token,请重试"))
			return
		}
		d1 := &model.Jwt{}
		j, err := jwt.ParseWithClaims(token, d1, func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenKey), nil
		})
		// j, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// 	return []byte(tokenKey), nil
		// })

		if err != nil {
			w.Write(util.NewResult(302, "token 非法,请重试"))
			return
		}
		// 合法
		if j.Valid {
			// fmt.Println(d1)

			// var id = j.Claims.(jwt.MapClaims)["Id"]
			// buf, _ := json.Marshal(d1)
			var num = d1.Num
			// fmt.Println(id)
			buf, _ := json.Marshal(num)
			// str := string(buf)
			// fmt.Println(str)
			r.Form.Add("num", string(buf))
			next.ServeHTTP(w, r)
		} else { // 不合法
			w.Write(util.NewResult(303, "不合法,请重试"))
			return
		}

	})
}

func main() {
	http.Handle(`/static/`, http.StripPrefix(`/static/`, http.FileServer(http.Dir("static"))))
	http.HandleFunc(`/`, control.ViewIndex)
	http.HandleFunc(`/moremsg`, control.Showmoremsg)
	http.HandleFunc(`/recommand`, control.ShowRecommand)
	http.HandleFunc(`/boke`, control.ViewBoke)

	//后台管理页面
	http.HandleFunc(`/login`, control.ShowLogin)                             //后台登陆
	http.HandleFunc(`/api/login`, control.CheckLogin)                        //验证登陆
	http.Handle(`/show/backtab`, Mid(http.HandlerFunc(control.Viewbacktab))) // 后台
	http.HandleFunc(`/post/api/msgs`, control.PostTwenmsgs)                  //图文信息msgs

	//数据逻辑处理
	http.HandleFunc(`/api/showfirstmsg`, control.ShowSecond)      //第二部分第一条数据
	http.HandleFunc(`/api/showmsgnext`, control.ShowSecondNext)   //第二部分剩下所有数据
	http.HandleFunc(`/api/workmsgs`, control.ShowNews)            //程序员要问
	http.HandleFunc(`/api/showrecommand`, control.ShowRecommend)  // 特别推荐消息
	http.HandleFunc(`/api/showresnews`, control.ShowNewRecommend) //显示最新推荐信息
	http.HandleFunc(`/api/showbowen`, control.ShowBowen)          //博文
	http.HandleFunc(`/api/timecar`, control.ShowTimeCar)          //时间轴信息
	http.HandleFunc(`/api/myself`, control.ShowMyself)            //显示个人信息
	http.HandleFunc(`/api/sendmsg`, control.SendLY)               //发送留言信息
	// http.HandleFunc(`/api/usermsgs`, control.ShowUserMsgs)        //显示用户留言信息
	http.HandleFunc(`/api/article/page`, control.ArtiPage) //分页

	// http.ListenAndServe(`:36`, nil)
	// fmt.Println("123")
	http.ListenAndServeTLS(":36", "cert-1542427206238_www.linchongyang.cn.crt", "cert-1542427206238_www.linchongyang.cn.key", nil)
	// fmt.Println("456")
}
