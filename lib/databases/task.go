package databases

import (
	"task-management/config"
	"task-management/models"
)

func CreateTask(newTask *models.Task) (interface{}, error) {
	tx := config.DB.Create(&newTask)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return newTask, nil
}
