package main

import (
	"time"
)


type Moment struct {
	id int64 `json:"id,omitempty" gorm:"primaryKey;column:id"`
	User_Nname string `json:"user_name,omitempty" gorm:"column:user_name"`
	User_Age int64 `json:"user_age,omitempty" gorm:"column:user_age"`
	User_Weight float64 `json:"user_weight,omitempty" gorm:"column:user_weight"`
	User_FatRate float64 `json:"user_fatrate,omitempty" gorm:"column:user_fatrate"`
	MomentComment string `json:"momentComment,omitempty" gorm:"column:momentComment"`
	Createtime time.Time `json:"createtime,omitempty" gorm:"column:createtime"`
	MomentStatus int `json:"momentStatus,omitempty" gorm:"column:momentStatus"`
}

