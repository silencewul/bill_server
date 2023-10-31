package pagination

import (
	"math"
	"xorm.io/xorm"
)

type PaginatorV3 struct {
	Count     int `json:"count"`
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
	PrevPage  int `json:"prev_page"`
	NextPage  int `json:"next_page"`
	Offset    int `json:"-"`
	Limit     int `json:"-"`
}

func NewV3(session *xorm.Session, page, limit int, result interface{}) (*PaginatorV3, error) {

	//验证page和limit并初始化start

	if page <= 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	start := (page - 1) * limit

	//查询记录和记录总数
	count, err := session.Limit(limit, start).FindAndCount(result)
	if err != nil {
		return nil, err
	}

	//计算总页数
	var totalPage int

	if count == 0 {
		page = 0
		totalPage = 0
	} else {
		totalPage = int(math.Ceil(float64(count) / float64(limit)))
	}

	//计算上一页和下一页
	var prevPage, nextPage int

	if page > 1 {
		prevPage = page - 1
	} else {
		prevPage = page
	}

	if page >= totalPage {
		nextPage = page
	} else {
		nextPage = page + 1
	}

	//返回结果
	return &PaginatorV3{
		Count:     int(count),
		TotalPage: totalPage,
		Offset:    start,
		Limit:     limit,
		Page:      page,
		PrevPage:  prevPage,
		NextPage:  nextPage,
	}, nil
}
