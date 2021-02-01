package service

import (
	"errors" 

	"github.com/gongxianjin/xcent-common/gorm"
	"github.com/gongxianjin/xcent-common/lib"
	"github.com/gongxianjin/xcent_scaffold/model"
	"github.com/gongxianjin/xcent_scaffold/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error

func CreateApi(api model.SysApi) (err error) {
	if !errors.Is( lib.GORMDefaultPool.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return  lib.GORMDefaultPool.Create(&api).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApi
//@description: 删除基础api
//@param: api model.SysApi
//@return: err error

func DeleteApi(api model.SysApi) (err error) {
	err =  lib.GORMDefaultPool.Delete(api).Error
	ClearCasbin(1, api.Path, api.Method)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAPIInfoList
//@description: 分页获取数据,
//@param: api model.SysApi, info request.PageInfo, order string, desc bool
//@return: err error

func GetAPIInfoList(api model.SysApi, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db :=  lib.GORMDefaultPool.Model(&model.SysApi{})
	var apiList []model.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err, apiList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllApis
//@description: 获取所有的api
//@return: err error, apis []model.SysApi

func GetAllApis() (err error, apis []model.SysApi) {
	err =  lib.GORMDefaultPool.Find(&apis).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: err error, api model.SysApi

func GetApiById(id float64) (err error, api model.SysApi) {
	err =  lib.GORMDefaultPool.Where("id = ?", id).First(&api).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateApi
//@description: 根据id更新api
//@param: api model.SysApi
//@return: err error

func UpdateApi(api model.SysApi) (err error) {
	var oldA model.SysApi
	err =  lib.GORMDefaultPool.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is( lib.GORMDefaultPool.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err =  lib.GORMDefaultPool.Save(&api).Error
		}
	}
	return err
}
