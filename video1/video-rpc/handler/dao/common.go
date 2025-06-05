package dao

import (
	"video1/video-rpc/global"
)

func Create[T any](data *T) bool {

	if err := global.DB.Create(data).Error; err != nil {
		return false
	}

	return true
}

func Delete[T any](data *T) bool {
	if err := global.DB.Delete(data).Error; err != nil {
		return false
	}
	return true
}

func BatchDelete[T any](data *T) bool { //批量删除
	if err := global.DB.Where(data).Delete(data).Error; err != nil {
		return false
	}
	return true
}

// save 结构体全量更新
// updates 有值的更新，不更新

func Update[T any](data *T) bool {
	if err := global.DB.Updates(data).Error; err != nil {
		return false
	}
	return true
}

func BatchUpdate[T any](where *T, data *T) bool { //批量更新
	if err := global.DB.Model(data).Where(where).Updates(data).Error; err != nil {
		return false
	}
	return true
}

func GetOneById[T any](id uint, data *T) bool {

	if err := global.DB.Where("id=?", id).Find(data).Limit(1).Error; err != nil {
		return false
	}
	return true
}

func GetOneByFields[T any](where *T, data *T) bool {

	if err := global.DB.Where(where).Find(data).Limit(1).Error; err != nil {
		return false
	}
	return true

}

func GetListByFields[T any](where *T, data *[]T) bool {

	if err := global.DB.Where(where).Find(data).Error; err != nil {
		return false
	}
	return true

}

type PageDto struct {
	Page     uint   `json:"page" form:"page"`
	PageSize uint   `json:"page_size" form:"page_size"`
	OrderBy  string `json:"order_by" form:"order_by"` // e.g. "created_at desc"
}

type PageResult[T any] struct {
	List     []T  `json:"list"`
	Total    uint `json:"total"`
	Page     uint `json:"page"`
	PageSize uint `json:"page_size"`
}

func Paginate[T any](where *T, req PageDto, out *[]T) (PageResult[T], error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	offset := (req.Page - 1) * req.PageSize

	query := global.DB.Model(out).Where(where)

	// 统计总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return PageResult[T]{}, err
	}

	// 排序
	if req.OrderBy != "" {
		query = query.Order(req.OrderBy)
	}

	// 查询数据
	if err := query.Offset(int(offset)).Limit(int(req.PageSize)).Find(out).Error; err != nil {
		return PageResult[T]{}, err
	}

	return PageResult[T]{
		List:     *out,
		Total:    uint(total),
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
