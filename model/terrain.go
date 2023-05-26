package model

import (
	"time"
)

// Member [...]
type Member struct {
	ID       uint   `gorm:"primaryKey;column:id" json:"-"`
	UUID     string `gorm:"column:uuid" json:"uuid,omitempty"`
	Nickname string `gorm:"column:nickName" json:"nickName,omitempty"`
}

// TableName get sql table name.获取数据库表名
func (m *Member) TableName() string {
	return "t_member"
}

// MemberColumns get sql column name.获取数据库列名
var MemberColumns = struct {
	ID       string
	UUID     string
	Nickname string
}{
	ID:       "id",
	UUID:     "uuid",
	Nickname: "nickName",
}

// SensorStat [...]
type SensorStat struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"-"`
	Time      time.Time `gorm:"column:time" json:"time"`
	DeletedAt string    `gorm:"column:deleted_at" json:"deletedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *SensorStat) TableName() string {
	return "t_sensor_stat"
}

// SensorStatColumns get sql column name.获取数据库列名
var SensorStatColumns = struct {
	ID        string
	Time      string
	DeletedAt string
}{
	ID:        "id",
	Time:      "time",
	DeletedAt: "deleted_at",
}
