package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _MemberMgr struct {
	*_BaseMgr
}

// MemberMgr open func
func MemberMgr(db *gorm.DB) *_MemberMgr {
	if db == nil {
		panic(fmt.Errorf("MemberMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MemberMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_member"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MemberMgr) GetTableName() string {
	return "t_member"
}

// Reset 重置gorm会话
func (obj *_MemberMgr) Reset() *_MemberMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MemberMgr) Get() (result Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Member{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MemberMgr) Gets() (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Member{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MemberMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Member{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MemberMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUUID uuid获取
func (obj *_MemberMgr) WithUUID(uuid string) Option {
	return optionFunc(func(o *options) { o.query["uuid"] = uuid })
}

// WithNickname nickName获取
func (obj *_MemberMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickName"] = nickname })
}

// GetByOption 功能选项模式获取
func (obj *_MemberMgr) GetByOption(opts ...Option) (result Member, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Member{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MemberMgr) GetByOptions(opts ...Option) (results []*Member, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Member{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MemberMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Member, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Member{}).Where(options.query)
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
func (obj *_MemberMgr) GetFromID(id uint) (result Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Member{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_MemberMgr) GetBatchFromID(ids []uint) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Member{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUUID 通过uuid获取内容
func (obj *_MemberMgr) GetFromUUID(uuid string) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Member{}).Where("`uuid` = ?", uuid).Find(&results).Error

	return
}

// GetBatchFromUUID 批量查找
func (obj *_MemberMgr) GetBatchFromUUID(uuids []string) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Member{}).Where("`uuid` IN (?)", uuids).Find(&results).Error

	return
}

// GetFromNickname 通过nickName获取内容
func (obj *_MemberMgr) GetFromNickname(nickname string) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Member{}).Where("`nickName` = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量查找
func (obj *_MemberMgr) GetBatchFromNickname(nicknames []string) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Member{}).Where("`nickName` IN (?)", nicknames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MemberMgr) FetchByPrimaryKey(id uint) (result Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Member{}).Where("`id` = ?", id).First(&result).Error

	return
}
