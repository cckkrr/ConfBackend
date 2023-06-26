package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _HeroPcdUoloadMgr struct {
	*_BaseMgr
}

// HeroPcdUoloadMgr open func
func HeroPcdUoloadMgr(db *gorm.DB) *_HeroPcdUoloadMgr {
	if db == nil {
		panic(fmt.Errorf("HeroPcdUoloadMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HeroPcdUoloadMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_hero_pcd_uoload"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_HeroPcdUoloadMgr) GetTableName() string {
	return "t_hero_pcd_uoload"
}

// Reset 重置gorm会话
func (obj *_HeroPcdUoloadMgr) Reset() *_HeroPcdUoloadMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_HeroPcdUoloadMgr) Get() (result HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_HeroPcdUoloadMgr) Gets() (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_HeroPcdUoloadMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_HeroPcdUoloadMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithFileUUID file_uuid获取
func (obj *_HeroPcdUoloadMgr) WithFileUUID(fileUUID string) Option {
	return optionFunc(func(o *options) { o.query["file_uuid"] = fileUUID })
}

// WithCreatedAt created_at获取
func (obj *_HeroPcdUoloadMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithOriginalUploadedFilename original_uploaded_filename获取
func (obj *_HeroPcdUoloadMgr) WithOriginalUploadedFilename(originalUploadedFilename string) Option {
	return optionFunc(func(o *options) { o.query["original_uploaded_filename"] = originalUploadedFilename })
}

// WithSavedFilename saved_filename获取
func (obj *_HeroPcdUoloadMgr) WithSavedFilename(savedFilename string) Option {
	return optionFunc(func(o *options) { o.query["saved_filename"] = savedFilename })
}

// WithFileSize file_size获取
func (obj *_HeroPcdUoloadMgr) WithFileSize(fileSize int64) Option {
	return optionFunc(func(o *options) { o.query["file_size"] = fileSize })
}

// WithSaveDuration save_duration获取
func (obj *_HeroPcdUoloadMgr) WithSaveDuration(saveDuration int) Option {
	return optionFunc(func(o *options) { o.query["save_duration"] = saveDuration })
}

// WithPcdFileType pcd_file_type获取 2d or 3d
func (obj *_HeroPcdUoloadMgr) WithPcdFileType(pcdFileType string) Option {
	return optionFunc(func(o *options) { o.query["pcd_file_type"] = pcdFileType })
}

// GetByOption 功能选项模式获取
func (obj *_HeroPcdUoloadMgr) GetByOption(opts ...Option) (result HeroPcdUoload, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_HeroPcdUoloadMgr) GetByOptions(opts ...Option) (results []*HeroPcdUoload, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_HeroPcdUoloadMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]HeroPcdUoload, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where(options.query)
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
func (obj *_HeroPcdUoloadMgr) GetFromID(id uint) (result HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_HeroPcdUoloadMgr) GetBatchFromID(ids []uint) (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromFileUUID 通过file_uuid获取内容
func (obj *_HeroPcdUoloadMgr) GetFromFileUUID(fileUUID string) (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`file_uuid` = ?", fileUUID).Find(&results).Error

	return
}

// GetBatchFromFileUUID 批量查找
func (obj *_HeroPcdUoloadMgr) GetBatchFromFileUUID(fileUUIDs []string) (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`file_uuid` IN (?)", fileUUIDs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_HeroPcdUoloadMgr) GetFromCreatedAt(createdAt time.Time) (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_HeroPcdUoloadMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromOriginalUploadedFilename 通过original_uploaded_filename获取内容
func (obj *_HeroPcdUoloadMgr) GetFromOriginalUploadedFilename(originalUploadedFilename string) (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`original_uploaded_filename` = ?", originalUploadedFilename).Find(&results).Error

	return
}

// GetBatchFromOriginalUploadedFilename 批量查找
func (obj *_HeroPcdUoloadMgr) GetBatchFromOriginalUploadedFilename(originalUploadedFilenames []string) (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`original_uploaded_filename` IN (?)", originalUploadedFilenames).Find(&results).Error

	return
}

// GetFromSavedFilename 通过saved_filename获取内容
func (obj *_HeroPcdUoloadMgr) GetFromSavedFilename(savedFilename string) (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`saved_filename` = ?", savedFilename).Find(&results).Error

	return
}

// GetBatchFromSavedFilename 批量查找
func (obj *_HeroPcdUoloadMgr) GetBatchFromSavedFilename(savedFilenames []string) (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`saved_filename` IN (?)", savedFilenames).Find(&results).Error

	return
}

// GetFromFileSize 通过file_size获取内容
func (obj *_HeroPcdUoloadMgr) GetFromFileSize(fileSize int64) (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`file_size` = ?", fileSize).Find(&results).Error

	return
}

// GetBatchFromFileSize 批量查找
func (obj *_HeroPcdUoloadMgr) GetBatchFromFileSize(fileSizes []int64) (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`file_size` IN (?)", fileSizes).Find(&results).Error

	return
}

// GetFromSaveDuration 通过save_duration获取内容
func (obj *_HeroPcdUoloadMgr) GetFromSaveDuration(saveDuration int) (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`save_duration` = ?", saveDuration).Find(&results).Error

	return
}

// GetBatchFromSaveDuration 批量查找
func (obj *_HeroPcdUoloadMgr) GetBatchFromSaveDuration(saveDurations []int) (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`save_duration` IN (?)", saveDurations).Find(&results).Error

	return
}

// GetFromPcdFileType 通过pcd_file_type获取内容 2d or 3d
func (obj *_HeroPcdUoloadMgr) GetFromPcdFileType(pcdFileType string) (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`pcd_file_type` = ?", pcdFileType).Find(&results).Error

	return
}

// GetBatchFromPcdFileType 批量查找 2d or 3d
func (obj *_HeroPcdUoloadMgr) GetBatchFromPcdFileType(pcdFileTypes []string) (results []*HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`pcd_file_type` IN (?)", pcdFileTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_HeroPcdUoloadMgr) FetchByPrimaryKey(id uint) (result HeroPcdUoload, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HeroPcdUoload{}).Where("`id` = ?", id).First(&result).Error

	return
}
