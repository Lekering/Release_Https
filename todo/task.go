package todo

import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool

	CreatedAt time.Time
	DoneAt    *time.Time
}

func NewTask(title, description string) *Task {
	return &Task{
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		DoneAt:      nil,
	}
}

func (t *Task) Done() {
	t.Completed = true
	t.DoneAt = &time.Time{Time: time.Now()}
}
