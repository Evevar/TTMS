package dao

import (
	"TTMS/kitex_gen/studio"
	"context"
	"errors"
)

func AddStudio(ctx context.Context, Name string, row, col int64) error {
	s := studio.Studio{Name: Name, RowsCount: row, ColsCount: col}
	tx := DB.WithContext(ctx).Create(&s)
	if tx.Error != nil {
		return tx.Error
	}
	return AddBatchSeat(ctx, &s)
}
func GetAllStudio(ctx context.Context, Current, PageSize int) ([]*studio.Studio, error) {
	studios := make([]*studio.Studio, PageSize)
	tx := DB.WithContext(ctx).Select("").Offset((Current - 1) * PageSize).Limit(PageSize).Find(&studios)
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
		tx := DB.Begin()
		if StudioInfo.RowsCount > 0 || StudioInfo.ColsCount > 0 { //演出厅规模变小之后删除座位
			var err error
			for i := 1; i <= int(s.RowsCount); i++ {
				for j := 1; j <= int(s.ColsCount); j++ {
					if i > int(StudioInfo.RowsCount) || j > int(StudioInfo.ColsCount) {
						err = RealDeleteSeat(int(StudioInfo.Id), i, j)
					}
				}
			}
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
