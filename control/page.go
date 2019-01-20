package control

import (
	"io/ioutil"
	"net/http"
)

//主页面
func ViewIndex(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile(`view/index.html`)
	w.Write(buf)
}
