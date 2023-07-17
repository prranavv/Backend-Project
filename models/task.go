package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task_Name string `json:"task_name"`
	Priority  string `json:"priority"`
	Status    string `json:"status"`
}
