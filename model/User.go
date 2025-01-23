package model

import (
	"github.com/ginblog/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(200);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色"` // 1 为管理员
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

func CheckUpUser(id int, name string) int {
	var users *User
	db.Select("id, username").Where("username = ?", name).First(&users)
	if users.ID == uint(id) {
		return errmsg.SUCCESS
	}
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

// 查询单个用户
func GetUser(id int) (User, int) {
	var user User
	err := db.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(username string, pageSize, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	if username == "" {
		db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total)
		return users, total
	}

	db.Where("username LIKE ?", username+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total)

	if err == gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// 编辑用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

// 更新密码
func UpdatePassword(id int, data *User) int {
	var user User
	if data.Password == "" {
		return errmsg.ERROR_PASSWORD_EMPTY
	}
	HashPassword := EncodePassword(data.Password)

	var maps = make(map[string]interface{})
	maps["password"] = HashPassword
	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

// 密码加密
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
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

	passWordError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if passWordError != nil {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NOT_PERMISSION
	}
	return errmsg.SUCCESS
}
