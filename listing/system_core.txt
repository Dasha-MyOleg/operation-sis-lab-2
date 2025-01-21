package core

import (
	"fmt"
	"os"
)

// SystemCore ?????????? ?? ?????????? ????? ???????? ???'???
type SystemCore struct {
	TaskQueue         []*TaskHandler
	AvailableFrames   []*PageFrame
	OccupiedFrames    []*PageFrame
	MemoryLimitMin    int
	MemoryLimitMax    int
	ReqPageMin        int
	ReqPageMax        int
	ReqWorkSetMin     int
	ReqWorkSetMax     int
	NumMemoryAccesses int
	pageEvictionAlgo  PageEviction
	swapOutCounter    int
}

// RefreshStatistics ??????? ?????????? ???????????? ????????
func (s *SystemCore) RefreshStatistics() {
	for _, frame := range s.OccupiedFrames {
		frame.Entry.Accessed = false
	}
	fmt.Println("Statistics updated.")
}

// ????? ????????? ?????????? ? ????
// LogFinalReport ??????? ???????? ?????????? ? ???? (??????????? ?????)
func (s *SystemCore) LogFinalReport(fileName string, pageFaultCount, accessCount int) {
	logFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer logFile.Close()

	logData := fmt.Sprintf(
		"=== Final Statistics ===\n"+
			"Total accesses: %d\n"+
			"Total page faults: %d\n"+
			"Total page replacements: %d\n"+
			"Replacement to fault ratio: %.2f%%\n",
		accessCount, pageFaultCount, s.swapOutCounter,
		float32(s.swapOutCounter)/float32(pageFaultCount)*100,
	)

	_, writeErr := logFile.WriteString(logData)
	if writeErr != nil {
		fmt.Println("Error writing to log file:", writeErr)
	} else {
		fmt.Println("Final statistics successfully written to", fileName)
	}
}

// ????????????? ???????
func (s *SystemCore) InitializeSystem(n int) {
	fmt.Println("Initializing system...")
	s.TaskQueue = make([]*TaskHandler, 0)
	s.AvailableFrames = make([]*PageFrame, n)

	chooseAlg := GenerateRandom(0, 100)
	if chooseAlg >= 50 {
		s.pageEvictionAlgo = &RandomReplacement{}
		fmt.Println("Using RANDOM replacement algorithm")
	} else {
		s.pageEvictionAlgo = &NRUAlgorithm{}
		fmt.Println("Using NRU replacement algorithm")
	}

	for i := 0; i < n; i++ {
		entry := &PageEntry{}
		frame := &PageFrame{Entry: entry, Index: i, AccessCount: 0}
		s.AvailableFrames[i] = frame
	}

	s.MemoryLimitMin = 15
	s.MemoryLimitMax = 20
	s.NumMemoryAccesses = 10
	s.ReqPageMin = 100
	s.ReqPageMax = 150
	s.ReqWorkSetMin = 10
	s.ReqWorkSetMax = 15
	fmt.Println("System ready to work!")
}

// ????????? ??????? ???????? ? ????????? ??????
func (s *SystemCore) RemoveAvailableFrame(idx int) {
	s.AvailableFrames = append(s.AvailableFrames[:idx], s.AvailableFrames[idx+1:]...)
}

// ????? ?????????? ?????
func (s *SystemCore) DisplayFinalReport(pageFaultCount, accessCount int) {
	fmt.Println("=== Final Statistics ===")
	fmt.Println("Total accesses:", accessCount)
	fmt.Println("Total page faults:", pageFaultCount)
	fmt.Println("Total page replacements:", s.swapOutCounter)

	if pageFaultCount > 0 {
		ratio := float32(s.swapOutCounter) / float32(pageFaultCount) * 100
		fmt.Println("Replacement to fault ratio:", ratio, "%")
	} else {
		fmt.Println("Replacement to fault ratio: 0 % (No page faults occurred)")
	}
}
