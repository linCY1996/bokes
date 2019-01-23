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

	http.ListenAndServe(`:80`, nil)
}
