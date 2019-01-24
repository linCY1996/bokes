package model

//处理图片信息的go文件

type Lunimgs struct {
	ID   int    `json:"id" form:"id"`
	Imgs string `json:"imgs" form:"imgs"`
	Msgs string `json:"msgs" form:"msgs"`
	Time string `json:"time" form:"time"`
}

//轮播图
func Lun() ([]Lunimgs, error) {
	mod := make([]Lunimgs, 0)
	err := DB.Select(&mod, `select * from lunimgs`)
	return mod, err
}

//second
type Second struct {
	ID       int    `json:"id" xml:"id"`
	Imgs1    string `json:"imgs1" xml:"imgs1"`
	Msgs1    string `json:"msgs1" xml:"msgs1"`
	Imgs2    string `json:"imgs2" xml:"imgs2"`
	Msgs2    string `json:"msgs2" xml:"msgs2"`
	Moremsgs string `json:"moremsgs" xml:"moremsgs"`
	Time     string `json:"time" xml:"time"`
}

//显示第二部分第一条信息
func WriteSecond(id int) (Second, error) {
	mod := Second{}
	err := DB.Get(&mod, `select * from second where id=?`, id)
	return mod, err
}

//显示第二部分剩下所有信息
func WriteSecondnext() ([]Second, error) {
	mod := make([]Second, 0)
	err := DB.Select(&mod, `select * from second where id>1`)
	return mod, err
}

//程序员的新闻
type News struct {
	ID    int    `json:"id" xml:"id"`
	Msgs  string `json:"msgs" xml:"msgs"`
	Time  string `json:"time" xml:"time"`
	Hrefs string `json:"hrefs" xml:"hrefs"`
}

func ShowNews() ([]News, error) {
	mod := make([]News, 0)
	err := DB.Select(&mod, `select * from news`)
	return mod, err
}

//特别推荐板块
type Recommend struct {
	ID    int    `json:"id" xml:"id"`
	Imgs  string `json:"imgs" xml:"imgs"`
	Title string `json:"title" xml:"title"`
	Msgs  string `json:"msgs" xml:"msgs"`
	Time  string `json:"time" xml:"time"`
	Hrefs string `json:"hrefs" xml:"hrefs"`
}

func ShowRec() ([]Recommend, error) {
	mod := make([]Recommend, 0)
	err := DB.Select(&mod, `select * from recommend`)
	return mod, err
}

//最新推荐
type Newrecommend struct {
	ID    int    `json:"id" xml:"id"`
	Imgs  string `json:"imgs" xml:"imgs"`
	Msgs  string `json:"msgs" xml:"msgs"`
	Hrefs string `json:"hrefs" xml:"hrefs"`
	Time  string `json:"time" xml:"time"`
}

func ShowNewRecommend() ([]Newrecommend, error) {
	mod := make([]Newrecommend, 0)
	err := DB.Select(&mod, `select * from newrecommend where id<=6`)
	return mod, err
}

//精彩博文
type Bowen struct {
	ID    int    `json:"id" xml:"id"`
	Imgs  string `json:"imgs" xml:"imgs"`
	Title string `json:"title" xml:"title"`
	Msgs  string `json:"msgs" xml:"msgs"`
	Time  string `json:"time" xml:"time"`
	Hrefs string `json:"hrefs" xml:"hrefs"`
}

func ShowBowen() ([]Bowen, error) {
	mod := make([]Bowen, 0)
	err := DB.Select(&mod, `SELECT * from bowens where id>6 AND id<=10`)
	return mod, err
}

//时间轴
type TimeCar struct {
	ID    int    `json:"id" xml:"id"`
	Time  string `json:"time" xml:"time"`
	Msgs1 string `json:"msgs1" xml:"msgs1"`
	Msgs2 string `json:"msgs2" xml:"msgs2"`
}

func ShowTimecar() ([]TimeCar, error) {
	mod := make([]TimeCar, 0)
	err := DB.Select(&mod, `select * from timecar`)
	return mod, err
}

//自我信息
type Myself struct {
	ID        int    `json:"id" xml:"id"`
	Name      string `json:"name" xml:"name"`
	Address   string `json:"address" xml:"address"`
	Rightname string `json:"rightname" xml:"rightname"`
	Msgs      string `json:"msgs" xml:"msgs"`
	Imgs      string `json:"imgs" form:"imgs"`
}

func ShowMyself() (Myself, error) {
	mod := Myself{}
	err := DB.Get(&mod, `select * from mysqlf`)
	return mod, err
}

//留言信息
type Liuyan struct {
	ID    int    `json:"id" xml:"id"`
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email",form:"email"`
	Msgs  string `json:"msgs" xml:"msgs"`
	Time  string `json:"time" xml:"time"`
}

func SendLY(name, email, msgs, time string) error {
	_, err := DB.Exec("insert into liuyan(`name`, `email`, `msgs`, `time`) values(?,?,?,?)", name, email, msgs, time)
	return err
}

//显示用户留言信息
func ShowUsermsgs() ([]Liuyan, error) {
	mod := make([]Liuyan, 0)
	err := DB.Select(&mod, `select * from liuyan`)
	return mod, err
}
