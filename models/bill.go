package models

import (
	"bill/modules/constant"
	"bill/modules/log"
	"time"
)

type Bill struct {
	Id int `json:"id" xorm:"not null pk autoincr comment('用户ID') INT(11)"`
	UId int `json:"u_id" xorm:"INT(11)"`
	Kind int `json:"kind" xorm:"INT(1)"`
	Status int `json:"status" xorm:"INT(1)"`
	Money float32 `json:"money" xorm:"FLOAT(10,2)"`
	CategoryId int `json:"category_id"`
	Date time.Time `json:"date"`
	Note string `json:"note" xorm:"comment('备注') VARCHAR(255)"`
	CreatedAt time.Time `json:"created_at,omitempty" xorm:"created not null default 0 comment('创建时间') INT(10)"`
	UpdatedAt time.Time `json:"updated_at,omitempty" xorm:"updated not null default 0 comment('最后更新时间') index INT(10)"`
}


type BillList []*Bill

// Insert 账单新增
func (bill *Bill) Insert() error {
	affected, err := MasterDB.Insert(bill)
	if err != nil {
		log.GetSugar().Errorf("新增账单错误,sql错误:%s", err.Error())
		return constant.ErrServerInternalError
	}

	if affected == 0 {
		return constant.ErrCreateFail
	}

	return nil
}

func GetUserTeachingCourse(userInfo *User) (BillList, error) {
	ses := MasterDB.NewSession()
	defer ses.Close()

	ses.Where("user_id = ?", userInfo.Id)
	ses.OrderBy("created_at DESC")

	//countSes := ses.Clone()
	//defer countSes.Close()
	//count, err := countSes.Count(new(Course))
	//if err != nil {
	//	log.GetSugar().Errorf("获取用户在教课程计数出错,sql错误:%s", err.Error())
	//	return nil, nil, constant.ErrServerInternalError
	//}



	billList := make(BillList, 0)
	_,err := ses.Get(BillList{})
	//ses.Limit(paginator.Limit, paginator.Offset)

	if err != nil {
		log.GetSugar().Errorf("课程搜索出错,sql错误:%s", err.Error())
		return nil, constant.ErrServerInternalError
	}

	return billList, nil
}