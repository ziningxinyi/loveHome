package models

import (
	_ "time"
)

/* 用户 table_name = user */
type User struct {
	Id            int           `json:"user_id"`                           //用户编号
	Name          string        `orm:"size(32);unique"   json:"name"`      //用户昵称
	Password_hash string        `orm:"size(128)" json:"password"`          //用户密码加密
	Mobile        string        `orm:"size(11);unique" json:"mobile"`      //手机号
	Real_name     string        `orm:"size(32);unique"   json:"real_name"` //真实姓名
	Id_card       string        `orm:"size(20)"   json:"id_card"`          //身份证号
	Avatar_url    string        `orm:"size(256)" json:"avatar_url"`        //用户头像路径
	Houses        []*Hosuse     `orm:"reverse(many)" json:"houses"`        //用户发布的头像信息
	Orders        []*OrderHouse `orm:"reverse(many)" json:"orders"`        //用户下的订单
}

/*房屋信息 tabel_name = house*/
type House struct {
	Id    int            `json:"house_id"`
	User  *User         `orm:`
}