package di

import (
	"database/sql"
	"restapi/src/domain/service"
	"restapi/src/handler"
	"restapi/src/infrastructure"
)

func InitTask(db *sql.DB) handler.ITaskHandler {
	r := infrastructure.NewTaskRepository(db)
	s := service.NewTaskService(r)
	return handler.NewTaskHandler(s)
}
