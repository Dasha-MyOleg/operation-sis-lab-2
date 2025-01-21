package main

import (
	"fmt"
	"operation-sis-lab-2/core"
)

func main() {
	system := &core.SystemCore{}
	memoryUnit := &core.MemoryUnit{}

	system.InitializeSystem(25)
	system.CreateTask()
	system.CreateTask()
	system.CreateTask()

	for _, task := range system.TaskQueue {
		system.PrepareActivePages(task)
	}

	numOfTaskSwitches := 10
	for i := 0; i < numOfTaskSwitches; i++ {
		for taskIndex, _ := range system.TaskQueue {
			task := system.TaskQueue[taskIndex]
			newWorkingSetProb := core.GenerateRandom(0, 100)
			if newWorkingSetProb <= 10 {
				system.PrepareActivePages(task)
				fmt.Println("New working set generated for Task ?", taskIndex+1)
			}

			updateStatProb := core.GenerateRandom(0, 100)
			if updateStatProb <= 40 {
				system.RefreshStatistics()
			}

			for j := 0; j < system.NumMemoryAccesses; j++ {
				pageIndex := task.SelectPage()
				memoryUnit.RequestPageAccess(task.PageDirectory, system, pageIndex)
			}
		}
	}

	// Display final report
	system.DisplayFinalReport(memoryUnit.PageFaults, memoryUnit.TotalAccesses)

	// Write final report to log file in English
	system.LogFinalReport("results.log", memoryUnit.PageFaults, memoryUnit.TotalAccesses)
}
