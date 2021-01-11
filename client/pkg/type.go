package pkg

import "time"

//ConfigHandle config handle
type ConfigHandle func(string) error

//UserInfo user
type UserInfo struct {
	Name          string       `json:"name"`
	Password      string       `json:"password"`
	Emil          string       `json:"emil,omitempty"`
	CreateTime    time.Time    `json:"creatTime"`
	Expire        time.Time    `json:"expireTime"`
	ReadInConfig  ConfigHandle `json:"-"`
	WriteInConfig ConfigHandle `json:"-"`
}
