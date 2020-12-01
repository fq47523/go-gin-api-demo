package dao

import (
	m "apidemo/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
)



func QueryStudentAll(page int,limit int,sort string)  ([]m.Student,int) {
	var Count int
	Stu := []m.Student{}
	if sort == "+id" {
		DB.Limit(limit).Offset((page - 1) * limit).Preload("IDCard").Preload("Class").Preload("Teachers").Order("created_at asc").Find(&Stu)

	}else {
		DB.Limit(limit).Offset((page - 1) * limit).Preload("IDCard").Preload("Class").Preload("Teachers").Order("created_at desc").Find(&Stu)
	}

	DB.Model(&Stu).Count(&Count)

	return Stu,Count
}
/*
func QueryStudentSingle(id int)  m.Student {
	stu := m.Student{}
	//DB.Preload("Teachers").Preload("IDCard").Preload("Class").First(&stu,"id = ?", id)
	DB.Where("id = ?",id).First(&stu)
	DB.Model(&stu).Related(&stu.IDCard,"IDCard").Related(&stu.Class,"Class").Association("Teachers").Find(&stu.Teachers)
	return stu
}
*/
func CreateStudent(c *gin.Context) error{
	/*
	{
		"student_name": "liuliu",
		"id_card":{
			"num": 652
		},
		"class_id": 1,
		"teachers": [{
			"id":1
		}]
	}
	*/
	stu := m.Student{}

	if err := c.ShouldBindJSON(&stu);err != nil {
		log.Errorf(err,"CreateStudentInfo")
		return err
	}

	stuu := m.Student{
		StudentName: stu.StudentName,
		ClassID:     stu.ClassID,
		IDCard:      stu.IDCard,
	}

	DB.Create(&stuu)
	DB.Model(&stuu).Association("Teachers").Append(stu.Teachers)

	return nil
}

func UpdateStudent(c *gin.Context,id uint) error{
	/*
	{
		"id": 1,
		"student_name": "nihap",
		"id_card":{
			"num": 2233
		},
		"class_id": 1,
		"teachers": [{
			"id":1
		},
		{
			"id":2
		}
		]
	}
	*/
	stu := m.Student{}
	upobj := DB.Preload("IDCard").First(&stu,"id = ?",id)

	if err := c.ShouldBindJSON(&stu);err != nil {
		log.Errorf(err,"UpdateStudentInfo")
		return err
	}

	upobj.Omit("Teachers").Updates(&stu)
	DB.Model(&stu).Association("Teachers").Replace(stu.Teachers)
	return nil
}

func DeleteStudent(stuobj *m.Student) error {
	return  DB.Transaction(func(tx *gorm.DB) error {
		obj := m.Student{}
		// 在事务中做一些数据库操作 (这里应该使用 'tx' ，而不是 'db')
		if err := tx.Preload("Teachers").First(&obj,"id = ?",stuobj.ID).Error; err != nil {
			// 返回任意 err ，整个事务都会 rollback
			log.Errorf(err,"DeleteStudentInfo")
			return err
		}

		if err := tx.Model(&obj).Association("Teachers").Delete(obj.Teachers).Error; err != nil {
			log.Errorf(err,"DeleteStudentInfo")
			return err
		}

		if err := tx.Unscoped().Delete(m.IDCard{},"student_id = ?",stuobj.ID).Error; err != nil{
			log.Errorf(err,"DeleteStudentInfo")
			return err
		}

		if err := tx.Unscoped().Delete(&stuobj).Error; err != nil {
			log.Errorf(err,"DeleteStudentInfo")
			return err
		}
		// 返回 nil ，事务会 commit
		return nil
	})
}


func QueryClassList() []m.Class{
	class := []m.Class{}
	DB.Preload("Students").Find(&class)
	return class
}

func QueryTeacherList() []m.Teacher{
	teacher := []m.Teacher{}
	DB.Find(&teacher)

	return teacher
}