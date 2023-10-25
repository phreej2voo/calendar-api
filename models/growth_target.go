package models

import "time"

// 废弃: 小程序二期上线后删除，每月发育指标使用growth_indicators表
type GrowthTarget struct {
	ID            uint      `json:"-"`
	Name          string    `json:"name"`
	Desc          string    `json:"desc"`
	Color         string    `json:"color"`
	Sequence      int       `json:"-"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
	CalendarDayID uint      `json:"-"`
}
