package main

import (
	"boke/control"
	"net/http"
)

func main() {
	http.Handle(`/static/`, http.StripPrefix(`/static/`, http.FileServer(http.Dir("static"))))
	http.HandleFunc(`/`, control.ViewIndex)

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

	http.ListenAndServe(`:36`, nil)
	// fmt.Println("123")
	// http.ListenAndServeTLS(":36", "cert-1542427206238_www.linchongyang.cn.crt", "cert-1542427206238_www.linchongyang.cn.key", nil)
	// fmt.Println("456")
}
