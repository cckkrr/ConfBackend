package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ImGroupMemberMgr struct {
	*_BaseMgr
}

// ImGroupMemberMgr open func
func ImGroupMemberMgr(db *gorm.DB) *_ImGroupMemberMgr {
	if db == nil {
		panic(fmt.Errorf("ImGroupMemberMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ImGroupMemberMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_im_group_member"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ImGroupMemberMgr) GetTableName() string {
	return "t_im_group_member"
}

// Reset 重置gorm会话
func (obj *_ImGroupMemberMgr) Reset() *_ImGroupMemberMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ImGroupMemberMgr) Get() (result ImGroupMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupMember{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ImGroupMemberMgr) Gets() (results []*ImGroupMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupMember{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ImGroupMemberMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ImGroupMember{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ImGroupMemberMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGroupUUID group_uuid获取
func (obj *_ImGroupMemberMgr) WithGroupUUID(groupUUID string) Option {
	return optionFunc(func(o *options) { o.query["group_uuid"] = groupUUID })
}

// WithMemberUUID member_uuid获取
func (obj *_ImGroupMemberMgr) WithMemberUUID(memberUUID string) Option {
	return optionFunc(func(o *options) { o.query["member_uuid"] = memberUUID })
}

// GetByOption 功能选项模式获取
func (obj *_ImGroupMemberMgr) GetByOption(opts ...Option) (result ImGroupMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ImGroupMember{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ImGroupMemberMgr) GetByOptions(opts ...Option) (results []*ImGroupMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ImGroupMember{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ImGroupMemberMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ImGroupMember, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ImGroupMember{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_ImGroupMemberMgr) GetFromID(id uint) (result ImGroupMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupMember{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ImGroupMemberMgr) GetBatchFromID(ids []uint) (results []*ImGroupMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupMember{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGroupUUID 通过group_uuid获取内容
func (obj *_ImGroupMemberMgr) GetFromGroupUUID(groupUUID string) (results []*ImGroupMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupMember{}).Where("`group_uuid` = ?", groupUUID).Find(&results).Error

	return
}

// GetBatchFromGroupUUID 批量查找
func (obj *_ImGroupMemberMgr) GetBatchFromGroupUUID(groupUUIDs []string) (results []*ImGroupMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupMember{}).Where("`group_uuid` IN (?)", groupUUIDs).Find(&results).Error

	return
}

// GetFromMemberUUID 通过member_uuid获取内容
func (obj *_ImGroupMemberMgr) GetFromMemberUUID(memberUUID string) (results []*ImGroupMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupMember{}).Where("`member_uuid` = ?", memberUUID).Find(&results).Error

	return
}

// GetBatchFromMemberUUID 批量查找
func (obj *_ImGroupMemberMgr) GetBatchFromMemberUUID(memberUUIDs []string) (results []*ImGroupMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupMember{}).Where("`member_uuid` IN (?)", memberUUIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ImGroupMemberMgr) FetchByPrimaryKey(id uint) (result ImGroupMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupMember{}).Where("`id` = ?", id).First(&result).Error

	return
}
