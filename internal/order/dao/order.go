package dao

import (
	"TTMS/kitex_gen/order"
	"TTMS/kitex_gen/ticket"
	"errors"
)

func AddOrder(UserId, ScheduleId, SeatRow, SeatCol int, Date string) error {
	t := ticket.Ticket{ScheduleId: int64(ScheduleId), SeatRow: int32(SeatRow), SeatCol: int32(SeatCol)}
	DB.Where("schedule_id = ? and seat_row = ? and seat_col = ?", int64(ScheduleId), int32(SeatRow), int32(SeatCol)).Find(&t)
	if t.Id > 0 {
		o := order.Order{UserId: int64(UserId), ScheduleId: int64(ScheduleId), SeatRow: int32(SeatRow), SeatCol: int32(SeatCol), Date: Date, Value: t.Price, Type: 1}
		return DB.Create(&o).Error
	}
	return errors.New("票不存在")
}
