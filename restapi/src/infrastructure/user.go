package infrastructure

import (
	"database/sql"

	"restapi/src/domain/repository"
	"restapi/src/infrastructure/entity"
)

type TaskRepository struct {
	*sql.DB
}

func NewTaskRepository(db *sql.DB) repository.ITaskRepository {
	return &TaskRepository{db}
}

func (r *TaskRepository) GetTasks() ([]*entity.Task, error) {
	rows, err := r.Query(`SELECT id, title, tag, tag_number, created_at, updated_at FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*entity.Task
	for rows.Next() {
		task := entity.Task{}
		if err := rows.Scan(&task.ID, &task.Title, &task.Tag, &task.TagNumber, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (r *TaskRepository) FindByID(id string) (*entity.Task, error) {
	stmt, err := r.Prepare(`SELECT id, title, tag, tag_number, created_at, updated_at FROM tasks WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	task := &entity.Task{}
	if err = stmt.QueryRow(id).Scan(&task); err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) Create(task *entity.Task) (*entity.Task, error) {
	stmt, err := r.Prepare(`INSERT INTO tasks(id, title, tag_name, tag, created_at) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(task.ID, task.Title, task.TagNumber, task.Tag, task.CreatedAt); err != nil {
		return nil, err
	}
	return task, nil
}

// func (r *TaskRepository) Update(task *entity.Task) (*entity.Task, error) {
// 	stmt, err := r.Prepare(`UPDATE tasks SET name = ?, updated_at = ? WHERE id = ?`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	if _, err := stmt.Exec(task.Name, task.UpdatedAt, task.ID); err != nil {
// 		return nil, err
// 	}

// 	return task, nil
// }

// func (r *TaskRepository) Delete(id string) error {
// 	stmt, err := r.Prepare(`DELETE FROM tasks WHERE id = ?`)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	if _, err := stmt.Exec(id); err != nil {
// 		return err
// 	}

// 	return nil

// }
