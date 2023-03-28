package dao

import (
	"TTMS/kitex_gen/play"
	"TTMS/kitex_gen/studio"
	"context"
	"errors"
	"fmt"
	"time"
)

func AddPlay(ctx context.Context, PlayInfo *play.Play) error {
	p := play.Play{}
	if DB.WithContext(ctx).Where("name = ?", PlayInfo.Name).Find(&p); p.Id > 0 { //该剧目已经存在
		return errors.New("该剧目已经存在")
	}
	return DB.WithContext(ctx).Create(&PlayInfo).Error
}
func GetAllPlay(ctx context.Context, Current, PageSize int) ([]*play.Play, error) {
	plays := make([]*play.Play, PageSize)
	tx := DB.WithContext(ctx).Offset((Current - 1) * PageSize).Limit(PageSize).Find(&plays)
	return plays, tx.Error
}

func UpdatePlay(ctx context.Context, PlayInfo *play.Play) error {
	s := play.Play{}
	DB.WithContext(ctx).Where("id = ?", PlayInfo.Id).Find(&s)
	if s.Id > 0 {
		return DB.WithContext(ctx).Updates(&PlayInfo).Error
	} else {
		return errors.New("该剧目不存在")
	}
}
func DeletePlay(ctx context.Context, id int64) error {
	s := play.Play{}
	DB.WithContext(ctx).Where("id  = ?", id).Find(&s)
	if s.Id > 0 {
		return DB.WithContext(ctx).Delete(&s).Error
	} else {
		return errors.New("该剧目不存在")
	}
}
func AddSchedule(ctx context.Context, SInfo *play.Schedule) error {
	tx := DB.Begin()
	var s []map[string]string
	var m map[string]string
	tx.WithContext(ctx).Table("schedules").Select("schedules.show_time as start,plays.duration as dur").Joins("join plays on schedules.play_id=plays.id").Where("schedules.studio_id=?", SInfo.StudioId).Find(&s)
	tx.WithContext(ctx).Where("id = ?", SInfo.PlayId).Select("duration").Find(&m)
	m["start"] = SInfo.ShowTime
	if IsConflict(m, s) { //时间有冲突
		tx.Rollback()
		return errors.New("时间有冲突")
	}
	s1 := studio.Studio{}
	tx.WithContext(ctx).Where("id = ?", SInfo.StudioId).Find(&s1)
	tx.WithContext(ctx).Create(&SInfo)
	//生成演出票(还没有写)
	tx.Commit()
	return tx.Error
}
func IsConflict(m map[string]string, S []map[string]string) bool {
	ms, _ := time.Parse("2006-01-02 15:04:05", m["start"])
	duration, _ := time.ParseDuration(m["duration"])
	md := ms.Add(duration)
	for _, v := range S {
		as, _ := time.Parse("2006-01-02 15:04:05", v["start"])
		d, _ := time.ParseDuration(v["dur"])
		ad := as.Add(d)
		if (as.Before(ms) && ad.After(ms)) || (as.Before(md) && ad.After(md)) || (as.Before(ms) && ad.After(md)) || (as.After(ms) && ad.Before(md)) {
			//时间上有冲突
			fmt.Println("my ", ms, md)
			fmt.Println("they ", as, ad)
			return true
		}
	}
	return false
}
func GetAllSchedule(ctx context.Context, Current, PageSize int) ([]*play.Schedule, error) {
	schedules := make([]*play.Schedule, PageSize)
	tx := DB.WithContext(ctx).Offset((Current - 1) * PageSize).Limit(PageSize).Find(&schedules)
	return schedules, tx.Error
}

func UpdateSchedule(ctx context.Context, SInfo *play.Schedule) error {
	tx := DB.Begin()
	err := tx.Where("id = ?", SInfo.Id).Updates(SInfo).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var s []map[string]string
	var m map[string]string
	DB.WithContext(ctx).Table("schedules").Select("schedules.show_time as start,plays.duration as dur").Joins("join plays on schedules.play_id=plays.id").Where("schedules.StudioId=?", SInfo.StudioId).Find(&s)
	DB.WithContext(ctx).Where("id = ?", SInfo.PlayId).Select("duration").Find(&m)
	m["start"] = SInfo.ShowTime
	if IsConflict(m, s) { //时间有冲突
		tx.Rollback()
		return errors.New("时间有冲突")
	}
	tx.Commit()
	return nil
}
func DeleteSchedule(ctx context.Context, id int64) error {
	s := play.Schedule{}
	DB.WithContext(ctx).Where("id  = ?", id).Find(&s)
	if s.Id > 0 {
		return DB.WithContext(ctx).Delete(&s).Error
	} else {
		return errors.New("该演出计划不存在")
	}
}
