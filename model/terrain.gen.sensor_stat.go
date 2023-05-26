package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SensorStatMgr struct {
	*_BaseMgr
}

// SensorStatMgr open func
func SensorStatMgr(db *gorm.DB) *_SensorStatMgr {
	if db == nil {
		panic(fmt.Errorf("SensorStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SensorStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_sensor_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SensorStatMgr) GetTableName() string {
	return "t_sensor_stat"
}

// Reset 重置gorm会话
func (obj *_SensorStatMgr) Reset() *_SensorStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SensorStatMgr) Get() (result SensorStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SensorStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SensorStatMgr) Gets() (results []*SensorStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SensorStat{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SensorStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SensorStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SensorStatMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTime time获取
func (obj *_SensorStatMgr) WithTime(time time.Time) Option {
	return optionFunc(func(o *options) { o.query["time"] = time })
}

// WithDeletedAt deleted_at获取
func (obj *_SensorStatMgr) WithDeletedAt(deletedAt string) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_SensorStatMgr) GetByOption(opts ...Option) (result SensorStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SensorStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SensorStatMgr) GetByOptions(opts ...Option) (results []*SensorStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SensorStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SensorStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SensorStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SensorStat{}).Where(options.query)
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
func (obj *_SensorStatMgr) GetFromID(id uint) (result SensorStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SensorStat{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SensorStatMgr) GetBatchFromID(ids []uint) (results []*SensorStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SensorStat{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTime 通过time获取内容
func (obj *_SensorStatMgr) GetFromTime(time time.Time) (results []*SensorStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SensorStat{}).Where("`time` = ?", time).Find(&results).Error

	return
}

// GetBatchFromTime 批量查找
func (obj *_SensorStatMgr) GetBatchFromTime(times []time.Time) (results []*SensorStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SensorStat{}).Where("`time` IN (?)", times).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容
func (obj *_SensorStatMgr) GetFromDeletedAt(deletedAt string) (results []*SensorStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SensorStat{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找
func (obj *_SensorStatMgr) GetBatchFromDeletedAt(deletedAts []string) (results []*SensorStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SensorStat{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SensorStatMgr) FetchByPrimaryKey(id uint) (result SensorStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SensorStat{}).Where("`id` = ?", id).First(&result).Error

	return
}
