package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ImGroupInfoMgr struct {
	*_BaseMgr
}

// ImGroupInfoMgr open func
func ImGroupInfoMgr(db *gorm.DB) *_ImGroupInfoMgr {
	if db == nil {
		panic(fmt.Errorf("ImGroupInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ImGroupInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_im_group_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ImGroupInfoMgr) GetTableName() string {
	return "t_im_group_info"
}

// Reset 重置gorm会话
func (obj *_ImGroupInfoMgr) Reset() *_ImGroupInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ImGroupInfoMgr) Get() (result ImGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ImGroupInfoMgr) Gets() (results []*ImGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ImGroupInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ImGroupInfoMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGroupUUID group_uuid获取
func (obj *_ImGroupInfoMgr) WithGroupUUID(groupUUID string) Option {
	return optionFunc(func(o *options) { o.query["group_uuid"] = groupUUID })
}

// WithGroupName group_name获取
func (obj *_ImGroupInfoMgr) WithGroupName(groupName string) Option {
	return optionFunc(func(o *options) { o.query["group_name"] = groupName })
}

// WithCreatedAt created_at获取
func (obj *_ImGroupInfoMgr) WithCreatedAt(createdAt string) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithDeletedAt deleted_at获取
func (obj *_ImGroupInfoMgr) WithDeletedAt(deletedAt string) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithMember member获取
func (obj *_ImGroupInfoMgr) WithMember(member string) Option {
	return optionFunc(func(o *options) { o.query["member"] = member })
}

// GetByOption 功能选项模式获取
func (obj *_ImGroupInfoMgr) GetByOption(opts ...Option) (result ImGroupInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ImGroupInfoMgr) GetByOptions(opts ...Option) (results []*ImGroupInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ImGroupInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ImGroupInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where(options.query)
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
func (obj *_ImGroupInfoMgr) GetFromID(id uint) (result ImGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ImGroupInfoMgr) GetBatchFromID(ids []uint) (results []*ImGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGroupUUID 通过group_uuid获取内容
func (obj *_ImGroupInfoMgr) GetFromGroupUUID(groupUUID string) (results []*ImGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where("`group_uuid` = ?", groupUUID).Find(&results).Error

	return
}

// GetBatchFromGroupUUID 批量查找
func (obj *_ImGroupInfoMgr) GetBatchFromGroupUUID(groupUUIDs []string) (results []*ImGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where("`group_uuid` IN (?)", groupUUIDs).Find(&results).Error

	return
}

// GetFromGroupName 通过group_name获取内容
func (obj *_ImGroupInfoMgr) GetFromGroupName(groupName string) (results []*ImGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where("`group_name` = ?", groupName).Find(&results).Error

	return
}

// GetBatchFromGroupName 批量查找
func (obj *_ImGroupInfoMgr) GetBatchFromGroupName(groupNames []string) (results []*ImGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where("`group_name` IN (?)", groupNames).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ImGroupInfoMgr) GetFromCreatedAt(createdAt string) (results []*ImGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_ImGroupInfoMgr) GetBatchFromCreatedAt(createdAts []string) (results []*ImGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容
func (obj *_ImGroupInfoMgr) GetFromDeletedAt(deletedAt string) (results []*ImGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找
func (obj *_ImGroupInfoMgr) GetBatchFromDeletedAt(deletedAts []string) (results []*ImGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromMember 通过member获取内容
func (obj *_ImGroupInfoMgr) GetFromMember(member string) (results []*ImGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where("`member` = ?", member).Find(&results).Error

	return
}

// GetBatchFromMember 批量查找
func (obj *_ImGroupInfoMgr) GetBatchFromMember(members []string) (results []*ImGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where("`member` IN (?)", members).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ImGroupInfoMgr) FetchByPrimaryKey(id uint) (result ImGroupInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ImGroupInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}
