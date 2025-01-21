package core

import "fmt"

// PageEviction ???????? ????????? ??? ?????????? ?????? ????????
type PageEviction interface {
	EvictPage(system *SystemCore, frame **PageFrame)
}

// RandomReplacement ???????? ???????? ?????????? ?????? ????????
type RandomReplacement struct{}

func (r *RandomReplacement) EvictPage(system *SystemCore, frame **PageFrame) {
	candidateIdx := GenerateRandom(0, len(system.OccupiedFrames))
	*frame = system.OccupiedFrames[candidateIdx]
	system.OccupiedFrames = append(system.OccupiedFrames[:candidateIdx], system.OccupiedFrames[candidateIdx+1:]...)
	fmt.Println("Random Replacement: Page", (*frame).Index, "evicted")
}

// NRUAlgorithm ???????? ???????? Not Recently Used
type NRUAlgorithm struct{}

func (n *NRUAlgorithm) EvictPage(system *SystemCore, frame **PageFrame) {
	var candidateIdx int
	for i, f := range system.OccupiedFrames {
		if !f.Entry.Accessed {
			candidateIdx = i
			break
		}
	}
	*frame = system.OccupiedFrames[candidateIdx]
	system.OccupiedFrames = append(system.OccupiedFrames[:candidateIdx], system.OccupiedFrames[candidateIdx+1:]...)
	fmt.Println("NRU Algorithm: Page", (*frame).Index, "evicted")
}

// RequestPageAccess ???????? ????? ?? ?????? ?? ????????
func (mmu *MemoryUnit) RequestPageAccess(pageTable *PageDirectory, system *SystemCore, idx int) {
	entry := pageTable.Entries[idx]
	mmu.TotalAccesses++ // ????? ????? ?? ???????? ???????? ????????? ????????

	if !entry.Present {
		fmt.Println("Page fault: virtual page", idx, "is not mapped")
		mmu.PageFaults++ // ?????????? ????????? ???????? ????????

		var frame *PageFrame
		if len(system.AvailableFrames) > 0 {
			frame = system.AvailableFrames[len(system.AvailableFrames)-1]
			system.AvailableFrames = system.AvailableFrames[:len(system.AvailableFrames)-1]
		} else {
			system.pageEvictionAlgo.EvictPage(system, &frame)
			system.swapOutCounter++
		}
		frame.Entry = entry
		entry.FrameIndex = frame.Index
		entry.Present = true
		system.OccupiedFrames = append(system.OccupiedFrames, frame)
		fmt.Println("Mapped virtual page", idx, "to physical page", frame.Index)
	} else {
		fmt.Println("Regular access to page:", idx)
	}

	// ??????????, ?? ?? ??????? ??? ?????
	accessType := GenerateRandom(0, 100)
	if accessType < 50 {
		fmt.Println("Read from page", idx)
		entry.Accessed = true
	} else {
		fmt.Println("Write to page", idx)
		entry.Accessed = true
		entry.Modified = true
	}
}
