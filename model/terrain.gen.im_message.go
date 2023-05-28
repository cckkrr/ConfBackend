package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ImMessageMgr struct {
	*_BaseMgr
}

// ImMessageMgr open func
func ImMessageMgr(db *gorm.DB) *_ImMessageMgr {
	if db == nil {
		panic(fmt.Errorf("ImMessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ImMessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_im_message"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ImMessageMgr) GetTableName() string {
	return "t_im_message"
}

// Reset 重置gorm会话
func (obj *_ImMessageMgr) Reset() *_ImMessageMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ImMessageMgr) Get() (result ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ImMessageMgr) Gets() (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ImMessageMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ImMessageMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUUID uuid获取
func (obj *_ImMessageMgr) WithUUID(uuid string) Option {
	return optionFunc(func(o *options) { o.query["uuid"] = uuid })
}

// WithMsgType msg_type获取
func (obj *_ImMessageMgr) WithMsgType(msgType string) Option {
	return optionFunc(func(o *options) { o.query["msg_type"] = msgType })
}

// WithTextTypeText text_type_text获取
func (obj *_ImMessageMgr) WithTextTypeText(textTypeText string) Option {
	return optionFunc(func(o *options) { o.query["text_type_text"] = textTypeText })
}

// WithFileTypeURI file_type_uri获取
func (obj *_ImMessageMgr) WithFileTypeURI(fileTypeURI string) Option {
	return optionFunc(func(o *options) { o.query["file_type_uri"] = fileTypeURI })
}

// WithFromUserUUID from_user_uuid获取
func (obj *_ImMessageMgr) WithFromUserUUID(fromUserUUID string) Option {
	return optionFunc(func(o *options) { o.query["from_user_uuid"] = fromUserUUID })
}

// WithToEntityUUID to_entity_uuid获取
func (obj *_ImMessageMgr) WithToEntityUUID(toEntityUUID string) Option {
	return optionFunc(func(o *options) { o.query["to_entity_uuid"] = toEntityUUID })
}

// WithCreatedAt created_at获取
func (obj *_ImMessageMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_ImMessageMgr) GetByOption(opts ...Option) (result ImMessage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ImMessageMgr) GetByOptions(opts ...Option) (results []*ImMessage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ImMessageMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ImMessage, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where(options.query)
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
func (obj *_ImMessageMgr) GetFromID(id uint64) (result ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ImMessageMgr) GetBatchFromID(ids []uint64) (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUUID 通过uuid获取内容
func (obj *_ImMessageMgr) GetFromUUID(uuid string) (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`uuid` = ?", uuid).Find(&results).Error

	return
}

// GetBatchFromUUID 批量查找
func (obj *_ImMessageMgr) GetBatchFromUUID(uuids []string) (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`uuid` IN (?)", uuids).Find(&results).Error

	return
}

// GetFromMsgType 通过msg_type获取内容
func (obj *_ImMessageMgr) GetFromMsgType(msgType string) (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`msg_type` = ?", msgType).Find(&results).Error

	return
}

// GetBatchFromMsgType 批量查找
func (obj *_ImMessageMgr) GetBatchFromMsgType(msgTypes []string) (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`msg_type` IN (?)", msgTypes).Find(&results).Error

	return
}

// GetFromTextTypeText 通过text_type_text获取内容
func (obj *_ImMessageMgr) GetFromTextTypeText(textTypeText string) (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`text_type_text` = ?", textTypeText).Find(&results).Error

	return
}

// GetBatchFromTextTypeText 批量查找
func (obj *_ImMessageMgr) GetBatchFromTextTypeText(textTypeTexts []string) (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`text_type_text` IN (?)", textTypeTexts).Find(&results).Error

	return
}

// GetFromFileTypeURI 通过file_type_uri获取内容
func (obj *_ImMessageMgr) GetFromFileTypeURI(fileTypeURI string) (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`file_type_uri` = ?", fileTypeURI).Find(&results).Error

	return
}

// GetBatchFromFileTypeURI 批量查找
func (obj *_ImMessageMgr) GetBatchFromFileTypeURI(fileTypeURIs []string) (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`file_type_uri` IN (?)", fileTypeURIs).Find(&results).Error

	return
}

// GetFromFromUserUUID 通过from_user_uuid获取内容
func (obj *_ImMessageMgr) GetFromFromUserUUID(fromUserUUID string) (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`from_user_uuid` = ?", fromUserUUID).Find(&results).Error

	return
}

// GetBatchFromFromUserUUID 批量查找
func (obj *_ImMessageMgr) GetBatchFromFromUserUUID(fromUserUUIDs []string) (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`from_user_uuid` IN (?)", fromUserUUIDs).Find(&results).Error

	return
}

// GetFromToEntityUUID 通过to_entity_uuid获取内容
func (obj *_ImMessageMgr) GetFromToEntityUUID(toEntityUUID string) (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`to_entity_uuid` = ?", toEntityUUID).Find(&results).Error

	return
}

// GetBatchFromToEntityUUID 批量查找
func (obj *_ImMessageMgr) GetBatchFromToEntityUUID(toEntityUUIDs []string) (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`to_entity_uuid` IN (?)", toEntityUUIDs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ImMessageMgr) GetFromCreatedAt(createdAt time.Time) (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_ImMessageMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ImMessageMgr) FetchByPrimaryKey(id uint64) (result ImMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImMessage{}).Where("`id` = ?", id).First(&result).Error

	return
}
