package model

import (
	"app/shared/database"
	"time"
)

type Task struct {
	Id          int64
	OrgId       int64
	Summary     string
	Description string
	Due         *time.Time
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func TaskCreate(oid int64, summary string, description string, due *time.Time) (int64, error) {
	result, err := database.SQL.Exec(
		"INSERT INTO task (org_id, summary, description, due, done) VALUES (?, ?, ?, ?, ?)",
		oid, summary, description, due, false)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func TaskById(id int64) (*Task, error) {
	row := database.SQL.QueryRow(
		"SELECT id, org_id, summary, description, due, done, created_at, updated_at FROM task WHERE id = ? LIMIT 1",
		id)

	var t Task
	err := row.Scan(&t.Id, &t.OrgId, &t.Summary, &t.Description, &t.Due, &t.Done, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &t, err
}
