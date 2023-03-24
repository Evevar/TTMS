package dao

import (
	"TTMS/kitex_gen/studio"
	"context"
	"errors"
	"fmt"
)

func AddSeat(ctx context.Context, seatInfo *studio.Seat) error {
	s := studio.Seat{}
	if DB.WithContext(ctx).Where("studio_id = ? and row = ? and col = ? and status > 0", seatInfo.StudioId, seatInfo.Row, seatInfo.Col).Find(&s); s.Id > 0 {
		fmt.Println("不允许的行为：同一位置重复加入座位")
		return errors.New("不允许的行为：同一位置重复加入座位")
	}
	return DB.WithContext(ctx).Create(seatInfo).Error
}
func GetAllSeat(ctx context.Context, studioId, Current, PageSize int) ([]*studio.Seat, error) {
	seats := make([]*studio.Seat, PageSize)
	tx := DB.WithContext(ctx).Offset(Current-1).Limit(PageSize).Where("studio_id = ?", studioId).Find(seats)
	return seats, tx.Error
}
func UpdateSeat(ctx context.Context, seatInfo *studio.Seat) error {
	s := studio.Seat{}
	if DB.WithContext(ctx).Where("studio_id = ? and row = ? and col = ? and status > 0", seatInfo.StudioId, seatInfo.Row, seatInfo.Col).Find(&s); s.Id > 0 {
		return DB.WithContext(ctx).Updates(seatInfo).Error
	}
	return errors.New("该位置上无座位，修改失败")
}
func DeleteSeat(ctx context.Context, seatInfo *studio.Seat) error {
	return DB.WithContext(ctx).Delete(seatInfo).Error
}
