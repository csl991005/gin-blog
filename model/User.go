package model

import (
	"github.com/ginblog/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(200);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// 查询用户是否存在
func CheckUser(name string) int {
	var users *User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED // 1001
	}
	return errmsg.SUCCESS // 200
}

// 新增用户
func CreateUser(data *User) int {
	//data.Password = EncodePassword(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

// 查询用户列表
func GetUsers(pageSize, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 编辑用户
func EditUser(id int, data *User) int {
	var user User
	var updatemap = make(map[string]interface{})
	updatemap["username"] = data.Username
	updatemap["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Updates(updatemap).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

//删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

// 密码加密
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = EncodePassword(u.Password)
	return nil
}

func EncodePassword(password string) string {
	const DefaultCost int = 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

// 登录验证
func CheckLogin(username, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	passWordError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if passWordError != nil {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 0 {
		return errmsg.ERROR_USER_NOT_PERMISSION
	}
	return errmsg.SUCCESS
}
