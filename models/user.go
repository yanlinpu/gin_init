package models

import (
	"gin_init/pkg/database"
	"gin_init/pkg/logger"
)

type ObjStruct struct {
}

// get list
func GetListV1(condition map[string]interface{}, page, pageSize int) (list []ObjStruct) {
	start := 0
	if pageSize < 1 {
		pageSize = 20
	}
	if page > 1 {
		start = pageSize * (page - 1)
	}
	db := database.Conn()
	err := db.Where(condition).Order("created_at desc").Offset(start).Limit(pageSize).Find(&list).Error
	if err != nil {
		logger.Error("Get list error :", err.Error())
	}
	return
}

// get count
func GetCountV1(condition map[string]interface{}) (total int64) {
	db := database.Conn()
	err := db.Model(&ObjStruct{}).Where(condition).Count(&total).Error
	if err != nil {
		logger.Error("Get count error :", err.Error())
	}
	return
}

/*
	query := "username = ? "
	args := []interface{}{"cp1"}
	query += "AND province_id in (?) "
	args = append(args, []int{1,2,3})
	GetListV2(params.Page, params.PageSize, "id desc", query, args...)
*/
func GetListV2(page, pageSize int, orderBy string, query interface{}, args ...interface{}) (list []ObjStruct, err error) {
	start := 0
	if pageSize < 1 {
		pageSize = 20
	}
	if page > 1 {
		start = pageSize * (page - 1)
	}
	db := database.Conn()
	if orderBy == "" {
		orderBy = "id asc"
	}
	err = db.Where(query, args...).Order(orderBy).Offset(start).Limit(pageSize).Find(&list).Error
	if err != nil {
		logger.Error("GetList error :", err.Error())
		return
	}
	return
}

func GetCountV2(query interface{}, args ...interface{}) (total int, err error) {
	db := database.Conn()
	err = db.Model(&ObjStruct{}).Where(query, args...).Count(&total).Error
	if err != nil {
		logger.Error("GetCount error :", err.Error())
		return
	}
	return
}

/**
func (o *ObjStruct) Method() (err error) {
	db := database.Conn()
	// err = db.Create(o).Error
	// err = db.Delete(o).Error
	// err = db.Model(&ObjStruct{}).Updates(map[string]interface{}{}).Error
	// err = db.First(&ObjStruct{}, "id = ?", id).Error
	// 预加载       Preload		 db.Preload("Orders").Find(&users)  						SELECT * FROM users; SELECT * FROM orders WHERE user_id IN (1,2,3,4);
	// 关联    	   Association	db.Model(&post).Association("Category").Find(&category1)
	// 关联属于 	Related		 db.Model(&user).Related(&profile)  						SELECT * FROM profiles WHERE id = 111; // 111是user的外键ProfileID
	if err != nil {
		logger.Error("Method error :", err.Error())
	}
	return

	// 事务
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err = tx.Error
	if err != nil {
		return
	}

	...

	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		logger.Error("transaction error :", err.Error())
	}
	return

}
**/
