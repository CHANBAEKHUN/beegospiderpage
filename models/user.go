package models

import (
	"DataCertProject/db_mysql"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Id int      `form:"id"`
	Phone string    `form:"phone"`
	Password string    `form:"password"`
}
func (u User) SaverUser()(int64,error){
//密码脱敏处理
	hashMd5:=md5.New()
	hashMd5.Write([]byte(u.Password))
	bytes:=hashMd5.Sum(nil)
	u.Password=hex.EncodeToString(bytes)

	//连接数据库
	result,err:=db_mysql.Db.Exec("insert into user(phone,password) values(?,?)",u.Phone,u.Password)
	if err!=nil{
		return -1,nil
	}
	id,err:=result.RowsAffected()
	if err!=nil{
		return -1,err
	}
	return id,nil
}

//查询用户信息
func (u User)QueryUser()(*User,error){
	row:=db_mysql.Db.QueryRow("select phone from user where phone =? and password =? ",u.Password,u.Phone)
	err:=row.Scan(&u.Phone)
	if err!=nil{
		return nil,err
	}
	return &u,nil

}

