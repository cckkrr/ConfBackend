package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ImCacheInboxMgr struct {
	*_BaseMgr
}

// ImCacheInboxMgr open func
func ImCacheInboxMgr(db *gorm.DB) *_ImCacheInboxMgr {
	if db == nil {
		panic(fmt.Errorf("ImCacheInboxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ImCacheInboxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_im_cache_inbox"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ImCacheInboxMgr) GetTableName() string {
	return "t_im_cache_inbox"
}

// Reset 重置gorm会话
func (obj *_ImCacheInboxMgr) Reset() *_ImCacheInboxMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ImCacheInboxMgr) Get() (result ImCacheInbox, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImCacheInbox{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ImCacheInboxMgr) Gets() (results []*ImCacheInbox, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImCacheInbox{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ImCacheInboxMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ImCacheInbox{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ImCacheInboxMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCacheMsgUUID cache_msg_uuid获取
func (obj *_ImCacheInboxMgr) WithCacheMsgUUID(cacheMsgUUID string) Option {
	return optionFunc(func(o *options) { o.query["cache_msg_uuid"] = cacheMsgUUID })
}

// WithInboxUserUUID inbox_user_uuid获取
func (obj *_ImCacheInboxMgr) WithInboxUserUUID(inboxUserUUID string) Option {
	return optionFunc(func(o *options) { o.query["inbox_user_uuid"] = inboxUserUUID })
}

// GetByOption 功能选项模式获取
func (obj *_ImCacheInboxMgr) GetByOption(opts ...Option) (result ImCacheInbox, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ImCacheInbox{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ImCacheInboxMgr) GetByOptions(opts ...Option) (results []*ImCacheInbox, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ImCacheInbox{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ImCacheInboxMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ImCacheInbox, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ImCacheInbox{}).Where(options.query)
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
func (obj *_ImCacheInboxMgr) GetFromID(id uint64) (result ImCacheInbox, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImCacheInbox{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ImCacheInboxMgr) GetBatchFromID(ids []uint64) (results []*ImCacheInbox, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImCacheInbox{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCacheMsgUUID 通过cache_msg_uuid获取内容
func (obj *_ImCacheInboxMgr) GetFromCacheMsgUUID(cacheMsgUUID string) (results []*ImCacheInbox, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImCacheInbox{}).Where("`cache_msg_uuid` = ?", cacheMsgUUID).Find(&results).Error

	return
}

// GetBatchFromCacheMsgUUID 批量查找
func (obj *_ImCacheInboxMgr) GetBatchFromCacheMsgUUID(cacheMsgUUIDs []string) (results []*ImCacheInbox, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImCacheInbox{}).Where("`cache_msg_uuid` IN (?)", cacheMsgUUIDs).Find(&results).Error

	return
}

// GetFromInboxUserUUID 通过inbox_user_uuid获取内容
func (obj *_ImCacheInboxMgr) GetFromInboxUserUUID(inboxUserUUID string) (results []*ImCacheInbox, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImCacheInbox{}).Where("`inbox_user_uuid` = ?", inboxUserUUID).Find(&results).Error

	return
}

// GetBatchFromInboxUserUUID 批量查找
func (obj *_ImCacheInboxMgr) GetBatchFromInboxUserUUID(inboxUserUUIDs []string) (results []*ImCacheInbox, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImCacheInbox{}).Where("`inbox_user_uuid` IN (?)", inboxUserUUIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ImCacheInboxMgr) FetchByPrimaryKey(id uint64) (result ImCacheInbox, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImCacheInbox{}).Where("`id` = ?", id).First(&result).Error

	return
}
