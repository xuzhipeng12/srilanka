package base

import (
	"gorm.io/gorm"
	"math"
)

type Paginator struct {
	TotalCount int         `json:"total_count"`
	TotalPage  int         `json:"total_page"`
	Data       interface{} `json:"data"`
}

func Paging(db *gorm.DB, limit int, page int, result interface{}) (paginator *Paginator, err error) {
	var totalCount int64
	db.Count(&totalCount)
	db.Debug().Limit(limit).Offset((page - 1) * limit).Scan(result)
	paginator = &Paginator{
		TotalCount: int(totalCount),
		TotalPage:  int(math.Ceil(float64(totalCount) / float64(limit))),
		Data:       result,
	}
	return
}
