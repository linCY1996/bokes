package control

import (
	"boke/model"
	"boke/util"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func Lunimg(w http.ResponseWriter, r *http.Request) {
	mod, err := model.Lun()
	if err != nil {
		w.Write([]byte("查询信息错误"))
		return
	}
	w.Header().Set(`Content-Type`, `application/json`)
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//博客第二部分的第一条数据显示
func ShowSecond(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var kid = r.FormValue(`id`)
	id, _ := strconv.Atoi(kid)
	mod, err := model.WriteSecond(id)
	if err != nil {
		w.Write([]byte("查询信息失败"))
		return
	}
	w.Header().Set(`Content-Type`, `application/json`)
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//显示第二部分剩下所有信息
func ShowSecondNext(w http.ResponseWriter, r *http.Request) {
	mod, err := model.WriteSecondnext()
	if err != nil {
		w.Write([]byte("查询信息失败"))
		return
	}
	w.Header().Set(`Content-Type`, `application/json`)
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//最新推荐中的新闻推荐
func ShowNews(w http.ResponseWriter, r *http.Request) {
	mod, err := model.ShowNews()
	if err != nil {
		w.Write([]byte("查询信息失败"))
		return
	}
	w.Header().Set(`Content-Type`, `application/json`)
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//显示特别推荐信息
func ShowRecommend(w http.ResponseWriter, r *http.Request) {
	mod, err := model.ShowRec()
	if err != nil {
		w.Write([]byte("查询信息失败"))
		return
	}
	w.Header().Set(`Content-Type`, `application/json`)
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//显示最新推荐的新闻
func ShowNewRecommend(w http.ResponseWriter, r *http.Request) {
	mod, err := model.ShowNewRecommend()
	if err != nil {
		w.Write([]byte("查询信息失败"))
		return
	}
	w.Header().Set(`Content-Type`, `application/json`)
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//显示精彩博文信息
func ShowBowen(w http.ResponseWriter, r *http.Request) {
	mod, err := model.ShowBowen()
	if err != nil {
		w.Write([]byte("查询信息失败"))
		return
	}
	w.Header().Set(`Content-Type`, `application/json`)
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//显示时间轴上的信息
func ShowTimeCar(w http.ResponseWriter, r *http.Request) {
	mod, err := model.ShowTimecar()
	if err != nil {
		w.Write([]byte("查询信息失败"))
		return
	}
	w.Header().Set(`Content-Type`, `application/json`)
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//显示个人信息
func ShowMyself(w http.ResponseWriter, r *http.Request) {
	mod, err := model.ShowMyself()
	if err != nil {
		w.Write([]byte("查询信息失败"))
		return
	}
	w.Header().Set(`Content-Type`, `application/json`)
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//发送留言信息
func SendLY(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var name = r.FormValue(`name`)
	var email = r.FormValue(`email`)
	var msgs = r.FormValue(`msgs`)
	var time = time.Now().Format(`2006-01-02`)
	if name == `` || email == `` || msgs == `` {
		w.Write([]byte(`发送信息不能存在空`))
		return
	}
	err := model.SendLY(name, email, msgs, time)

	w.Header().Set(`Content-Type`, `application/json`)
	if err != nil {
		w.Write([]byte(`信息发送失败`))
		return
	}
	w.Write([]byte(`发送成功`))
}

//显示用户留言信息
// func ShowUserMsgs(w http.ResponseWriter, r *http.Request) {
// 	mod, err := model.ShowUsermsgs()
// 	if err != nil {
// 		w.Write([]byte("查询信息失败"))
// 		return
// 	}
// 	w.Header().Set(`Content-Type`, `application/json`)
// 	md, _ := json.Marshal(mod)
// 	w.Write(md)
// }

//显示分页的逻辑操作
func ArtiPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	count, _ := model.ArticleCount()
	var pi = r.FormValue(`pi`)
	pii, _ := strconv.Atoi(pi)

	var ps = r.FormValue(`ps`)
	pss, _ := strconv.Atoi(ps)
	mod, err := model.ArticlPage(pii, pss)
	if err != nil {
		w.Write(util.NewResult(300, `信息错误`))
		return
	}
	// md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(util.NewPage(200, `正常`, mod, count))
}
