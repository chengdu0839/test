package class

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id     string `orm:"pk"`
	Nick   string
	Info   string `orm:"null"`
	Hobby  string `orm:"null"`
	Email  string `orm:"unique"`
	Avator string `orm:"null"`
	Url    string `orm:"null"`

	Posts int

	Followers int
	Following int

	Regtime time.Time `orm:"auto_now_add"`

	Password string
	Private  int //权限控制。二进制数据(0或1)。第一位表示用户是否存在<1表示用户存在>；第二位表示用户是否可以登录<1表示用户可以登录>；第三位表示用户是否有发表文章的权限<1表示可以>
}

const (
	PR_live = iota
	PR_login
	PR_post
)

const (
	DefaultPvt = 1<<3 - 1
)

//	CRUD
//	create
//	read
//	update
//	delete

func (u *User) ReadDB() (err error) {
	o := orm.NewOrm()
	err = o.Read(u)
	return
}

func (u User) Create() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(&u)
	return
}

func (u User) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.Update(&u)
	return
}

// 将用户设置为不存在
func (u User) Delete() (err error) {
	//	xxx1 & 1110 = xxx0
	//	~x ==> ^x (-1 ^ x)
	u.Private &= ^1
	err = u.Update()
	return
}

// 判断用户是否存在。
func (u User) ExistId() bool {
	o := orm.NewOrm()
	// 用户不存在
	if err := o.Read(&u); err == orm.ErrNoRows {
		return false
	}
	return true
}

// 查询邮件是否存在。read()方法只能针对组件查询
func (u User) ExistEmail() bool {
	o := orm.NewOrm()
	return o.QueryTable("user").Filter("Email", u.Email).Exist()
}
