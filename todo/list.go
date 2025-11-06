package todo

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) error {
	if _, ok := l.tasks[task.Title]; ok {
		return ErrTaskAlreadyExists
	}
	l.tasks[task.Title] = task
	return nil
}

func (l *List) ListTasks() map[string]Task {
	timeMap := make(map[string]Task)
	for title, task := range l.tasks {
		timeMap[title] = task
	}
	return timeMap
}

func (l *List) ListNotComleteTasks() map[string]Task {
	timeMapNotComp := make(map[string]Task)
	for title, task := range l.tasks {
		if !task.Completed {
			timeMapNotComp[title] = task
		}
	}
	return timeMapNotComp
}

func (l *List) CompleteTasks(title string) error {
	task, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}
	task.Complete()
	l.tasks[title] = task
	return nil
}

func (l *List) DeleteTask(title string) error {
	_, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}
	delete(l.tasks, title)
	return nil
}
