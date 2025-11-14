package todo

import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool

	CreatedAt    time.Time
	CompleteTime *time.Time
}

func NewTask(title, description string) Task {
	return Task{
		Title:        title,
		Description:  description,
		Completed:    false,
		CreatedAt:    time.Now(),
		CompleteTime: nil,
	}
}

func (t *Task) Complete() {
	now := time.Now()
	t.Completed = true
	t.CompleteTime = &now
}

func (t *Task) Uncomplete() {
	t.Completed = false
	t.CompleteTime = nil
}
