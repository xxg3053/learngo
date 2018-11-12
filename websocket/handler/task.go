package handler

import "sync"

var (
	Chm = make(map[string]chan Task)
	tx = new(sync.RWMutex)
)

type Task struct {
	Operator string
	Content string
}


func GetOrAddTask(name string) chan Task {
	tx.Lock()
	u, ok := Chm[name]
	if !ok{
		u = make(chan Task, 1000)
		Chm[name] = u
	}
	tx.Unlock()
	return u
}

func Push2TaskCh(task Task) {
	GetOrAddTask(task.Operator) <- task
}