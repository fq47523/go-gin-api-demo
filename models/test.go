package models

import "github.com/jinzhu/gorm"

type Class struct {
	// 班级，一个学生有一个班级，一个班级有多个学生
	gorm.Model
	ClassName string    `json:"class_name"`
	Students  []Student `json:"students"` // 一个班级有多个学生，一对多关系
}

type Student struct {
	// 学生
	gorm.Model
	StudentName string `json:"student_name"`
	Class       Class  `json:"class"`
	ClassID     uint   `json:"class_id"` // 一个学生有一个班级，对应一个班级有多个学生，学生通过ClassID来知道是属于哪个班级的
	IDCard      IDCard `json:"id_card"`  // 一个学生对应一个身份证，一对一关系
	// 一个学生有多个老师, 与老师建立多对多的关系
	Teachers []Teacher `gorm:"many2many:student_teacher" json:"teachers"` // 会产生一个关联表
}

type IDCard struct {
	// 身份证
	gorm.Model
	Num       int  `json:"num"`
	StudentID uint `json:"student_id"`
}

type Teacher struct {
	// 老师
	gorm.Model
	TeacherName string `json:"teacher_name"`
	// 一个老师有多个学生，与学生建立多对多的关系
	Students []Student `gorm:"many2many:student_teacher" json:"students"` // 会产生一个关联表
}