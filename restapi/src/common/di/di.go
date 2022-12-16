package di

import (
	"restapi/src/domain/service"
	"restapi/src/handler"
	"restapi/src/infrastructure"

	"gorm.io/gorm"
)

func InitTask(db *gorm.DB) handler.ITaskHandler {
	r := infrastructure.NewTaskRepository(db)
	s := service.NewTaskService(r)
	return handler.NewTaskHandler(s)
}

func InitTag(db *gorm.DB) handler.ITagHandler {
	r := infrastructure.NewTagRepository(db)
	s := service.NewTagService(r)
	return handler.NewTagHandler(s)
}
