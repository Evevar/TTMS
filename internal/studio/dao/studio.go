package dao

import (
	"TTMS/kitex_gen/studio"
	"context"
	"errors"
	"fmt"
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

func UpdateStudio(ctx context.Context, StudioInfo *studio.Studio) error {
	s := studio.Studio{}
	DB.WithContext(ctx).Where("id = ?", StudioInfo.Id).Find(&s)
	if s.Id > 0 {
		if StudioInfo.RowsCount > s.RowsCount || StudioInfo.ColsCount > s.ColsCount {
			return errors.New("座位规模不能比原来更大")
		}
		if StudioInfo.RowsCount == 0 {
			StudioInfo.RowsCount = s.RowsCount
		}
		if StudioInfo.ColsCount == 0 {
			StudioInfo.ColsCount = s.ColsCount
		}
		StudioInfo.SeatsCount = StudioInfo.RowsCount * StudioInfo.ColsCount
		//if StudioInfo.Name == "" {
		//	s.Name = StudioInfo.Name
		//}
		tx := DB.Begin()
		fmt.Println(s.SeatsCount, StudioInfo.SeatsCount)
		if s.SeatsCount > StudioInfo.SeatsCount {
			err := tx.WithContext(ctx).Where("studio_id=? and (row>? or col>?)", StudioInfo.Id, StudioInfo.RowsCount, StudioInfo.ColsCount).Delete(&studio.Seat{}).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		err := DB.WithContext(ctx).Model(&studio.Studio{}).Where("id = ?", StudioInfo.Id).Updates(&StudioInfo).Error
		if err != nil { //修改成功
			tx.Rollback()
			return err
		}
		tx.Commit()
	} else {
		return errors.New("该演出厅不存在")
	}
	return nil
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
