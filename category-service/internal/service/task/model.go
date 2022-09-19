package task

import "time"

type Task struct {
	ID        int64      `db:"id"`
	StartedAt *time.Time `db:"started_at"`
}
