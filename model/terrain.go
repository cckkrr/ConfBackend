package model

import (
	"time"
)

// ImCacheInbox [...]
type ImCacheInbox struct {
	ID            uint64 `gorm:"primaryKey;column:id" json:"-"`
	CacheMsgUUID  string `gorm:"column:cache_msg_uuid" json:"cacheMsgUuid"`
	InboxUserUUID string `gorm:"column:inbox_user_uuid" json:"inboxUserUuid"`
}

// TableName get sql table name.获取数据库表名
func (m *ImCacheInbox) TableName() string {
	return "t_im_cache_inbox"
}

// ImCacheInboxColumns get sql column name.获取数据库列名
var ImCacheInboxColumns = struct {
	ID            string
	CacheMsgUUID  string
	InboxUserUUID string
}{
	ID:            "id",
	CacheMsgUUID:  "cache_msg_uuid",
	InboxUserUUID: "inbox_user_uuid",
}

// ImGroupInfo [...]
type ImGroupInfo struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"-"`
	GroupUUID string `gorm:"column:group_uuid" json:"groupUuid"`
	GroupName string `gorm:"column:group_name" json:"groupName"`
	CreatedAt string `gorm:"column:created_at" json:"createdAt"`
	DeletedAt string `gorm:"column:deleted_at" json:"deletedAt"`
	Member    string `gorm:"column:member" json:"member"`
}

// TableName get sql table name.获取数据库表名
func (m *ImGroupInfo) TableName() string {
	return "t_im_group_info"
}

// ImGroupInfoColumns get sql column name.获取数据库列名
var ImGroupInfoColumns = struct {
	ID        string
	GroupUUID string
	GroupName string
	CreatedAt string
	DeletedAt string
	Member    string
}{
	ID:        "id",
	GroupUUID: "group_uuid",
	GroupName: "group_name",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
	Member:    "member",
}

// ImGroupMember [...]
type ImGroupMember struct {
	ID         uint   `gorm:"primaryKey;column:id" json:"-"`
	GroupUUID  string `gorm:"column:group_uuid" json:"groupUuid"`
	MemberUUID string `gorm:"column:member_uuid" json:"memberUuid"`
}

// TableName get sql table name.获取数据库表名
func (m *ImGroupMember) TableName() string {
	return "t_im_group_member"
}

// ImGroupMemberColumns get sql column name.获取数据库列名
var ImGroupMemberColumns = struct {
	ID         string
	GroupUUID  string
	MemberUUID string
}{
	ID:         "id",
	GroupUUID:  "group_uuid",
	MemberUUID: "member_uuid",
}

// ImMessage [...]
type ImMessage struct {
	ID           uint64    `gorm:"primaryKey;column:id" json:"-"`
	UUID         string    `gorm:"column:uuid" json:"uuid"`
	MsgType      string    `gorm:"column:msg_type" json:"msgType"`
	TextTypeText string    `gorm:"column:text_type_text" json:"textTypeText"`
	FileTypeURI  string    `gorm:"column:file_type_uri" json:"fileTypeUri"`
	FromUserUUID string    `gorm:"column:from_user_uuid" json:"fromUserUuid"`
	ToEntityUUID string    `gorm:"column:to_entity_uuid" json:"toEntityUuid"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`
}

// TableName get sql table name.获取数据库表名
func (m *ImMessage) TableName() string {
	return "t_im_message"
}

// ImMessageColumns get sql column name.获取数据库列名
var ImMessageColumns = struct {
	ID           string
	UUID         string
	MsgType      string
	TextTypeText string
	FileTypeURI  string
	FromUserUUID string
	ToEntityUUID string
	CreatedAt    string
}{
	ID:           "id",
	UUID:         "uuid",
	MsgType:      "msg_type",
	TextTypeText: "text_type_text",
	FileTypeURI:  "file_type_uri",
	FromUserUUID: "from_user_uuid",
	ToEntityUUID: "to_entity_uuid",
	CreatedAt:    "created_at",
}

// Member [...]
type Member struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"-"`
	UUID      string    `gorm:"column:uuid" json:"uuid"`
	Nickname  string    `gorm:"column:nickname" json:"nickname"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	LoginID   string    `gorm:"column:login_id" json:"loginId"`
	Password  string    `gorm:"column:password" json:"password"`
}

// TableName get sql table name.获取数据库表名
func (m *Member) TableName() string {
	return "t_member"
}

// MemberColumns get sql column name.获取数据库列名
var MemberColumns = struct {
	ID        string
	UUID      string
	Nickname  string
	CreatedAt string
	LoginID   string
	Password  string
}{
	ID:        "id",
	UUID:      "uuid",
	Nickname:  "nickname",
	CreatedAt: "created_at",
	LoginID:   "login_id",
	Password:  "password",
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
