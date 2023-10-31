package models

import "time"

type Category struct {
	Id        int       `json:"id" xorm:"not null pk autoincr comment('用户ID') INT(11)"`
	Title     string    `json:"title" xorm:"not null comment('名称') VARCHAR(255)"`
	Pid       int       `json:"pid" xorm:"INT(11)"`
	Kind      int       `json:"kind" xorm:"INT(11)"`
	Icon      string    `json:"icon" xorm:"VARCHAR(255)"`
	Status    int       `json:"status" xorm:"INT(11)"`
	CreatedAt time.Time `json:"created_at,omitempty" xorm:"created comment('创建时间')"`
	UpdatedAt time.Time `json:"updated_at,omitempty" xorm:"updated comment('最后更新时间')"`
}
