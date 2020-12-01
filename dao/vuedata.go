package dao

import (
	m "apidemo/models"
)

func QueryVueDataAll(page int,limit int,sort string) ( []m.VueData,int){
	VueDataAll := []m.VueData{}
	var Count int
	if sort == "+id" {
		DB.Limit(limit).Offset((page - 1) * limit).Preload("Platforms").Order("created_at asc").Find(&VueDataAll)


	}else {
		DB.Limit(limit).Offset((page - 1) * limit).Preload("Platforms").Order("created_at desc").Find(&VueDataAll)
	}
	DB.Model(&VueDataAll).Count(&Count)

	return VueDataAll,Count
}

func UpdateVueData(id uint,model m.VueData){
	mode_VueData := m.VueData{}
	gormobj := DB.Preload("Platforms").Where("id=?",id).First(&mode_VueData)

	gormobj.Save(&model)
}

func CreateVueData(model m.VueData) {
	DB.Create(&model)

}

func DeleteVueData(id int) error{
	err := DB.Where("id=?",id).Delete(&m.VueData{}).Error
	if err != nil{
		return err
	}
	return nil
}