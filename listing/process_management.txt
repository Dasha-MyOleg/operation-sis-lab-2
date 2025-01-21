package core

import "fmt"

// TaskHandler ??????????? ?????? ? ???????
type TaskHandler struct {
	PageDirectory *PageDirectory
	RequestCount  int
	ActivePages   *ActivePages
}

// ????????? ?????? ???????
func (s *SystemCore) CreateTask() {
	task := new(TaskHandler)
	addressSpace := GenerateRandom(s.MemoryLimitMin, s.MemoryLimitMax)
	task.PageDirectory = new(PageDirectory)
	task.PageDirectory.Entries = make([]*PageEntry, addressSpace)

	for i := 0; i < addressSpace; i++ {
		entry := &PageEntry{}
		task.PageDirectory.Entries[i] = entry
	}

	task.RequestCount = GenerateRandom(s.ReqPageMin, s.ReqPageMax)
	s.TaskQueue = append(s.TaskQueue, task)
	fmt.Println("Task created with page directory size:", len(task.PageDirectory.Entries))
}

// ?????????? ???????? ?????? ??? ???????
func (s *SystemCore) PrepareActivePages(task *TaskHandler) {
	if task.ActivePages == nil {
		task.ActivePages = &ActivePages{} // ????????????? ???????? ??????
	}

	size := GenerateRandom(s.ReqWorkSetMin, s.ReqWorkSetMax)
	task.ActivePages.PageIndexes = make([]int, size)
	for i := 0; i < size; i++ {
		task.ActivePages.PageIndexes[i] = i
	}
	fmt.Println("Active pages prepared with size:", len(task.ActivePages.PageIndexes))
}

// ????? ???????? ??? ???????
func (t *TaskHandler) SelectPage() int {
	probability := GenerateRandom(0, 100)
	if probability <= 90 {
		randIndex := GenerateRandom(0, len(t.ActivePages.PageIndexes))
		fmt.Println("Selecting page from active set", randIndex)
		return randIndex
	}
	randIndex := GenerateRandom(0, len(t.PageDirectory.Entries))
	fmt.Println("Selecting page from entire table", randIndex)
	return randIndex
}
