package models

type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      int    `json:"status" binding:"oneof=0 1"`
}

type TaskStore struct {
	tasks map[string]*Task
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: make(map[string]*Task),
	}
}

func (s *TaskStore) GetAll() []*Task {
	tasks := make([]*Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (s *TaskStore) Get(id string) (*Task, bool) {
	task, exists := s.tasks[id]
	return task, exists
}

func (s *TaskStore) Create(task *Task) {
	s.tasks[task.ID] = task
}

func (s *TaskStore) Update(task *Task) bool {
	if _, exists := s.tasks[task.ID]; !exists {
		return false
	}
	s.tasks[task.ID] = task
	return true
}

func (s *TaskStore) Delete(id string) bool {
	if _, exists := s.tasks[id]; !exists {
		return false
	}
	delete(s.tasks, id)
	return true
}
