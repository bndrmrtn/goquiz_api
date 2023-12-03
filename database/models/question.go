package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Id       uint           `json:"id" gorm:"primaryKey"`
	Question string         `json:"question" gorm:"question"`
	Image    string         `json:"image" gorm:"image"`
	Answers  datatypes.JSON `json:"answers" gorm:"answers"`
	Answer   string         `json:"-" gorm:"answer"`
	QuizId   uint
	Quiz     Quiz `gorm:"foreignKey:QuizId"`
}