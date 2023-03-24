package dao

import (
	"TTMS/kitex_gen/studio"
	"context"
	"errors"
)

func AddStudio(ctx context.Context, Name string, row, col int64) (*studio.Studio, error) {
	s := studio.Studio{Name: Name, RowsCount: row, ColsCount: col, SeatsCount: row * col}
	tx := DB.WithContext(ctx).Create(&s)
	return &s, tx.Error
}
func GetAllStudio(ctx context.Context, Current, PageSize int) ([]*studio.Studio, error) {
	studios := make([]*studio.Studio, PageSize)
	tx := DB.WithContext(ctx).Offset((Current - 1) * PageSize).Limit(PageSize).Find(&studios)
	return studios, tx.Error
}

func UpdateStudio(ctx context.Context, id, row, col int64, Name string) error {
	s := studio.Studio{}
	DB.WithContext(ctx).Where("id = ?", id).Find(&s)
	if s.Id > 0 {
		if row > s.RowsCount || col > s.ColsCount {
			return errors.New("座位规模不能比原来更大")
		}
		if row != 0 {
			s.RowsCount = row
		}
		if col != 0 {
			s.ColsCount = col
		}
		if Name != "" {
			s.Name = Name
		}
		return DB.WithContext(ctx).Updates(&s).Error
	} else {
		return errors.New("该演出厅不存在")
	}
}
func DeleteStudio(ctx context.Context, id int64) error {
	s := studio.Studio{}
	DB.WithContext(ctx).Where("id  = ?", id).Find(&s)
	if s.Id > 0 {
		return DB.WithContext(ctx).Delete(&s).Error
	} else {
		return errors.New("该演出厅不存在")
	}
}
