package heaps

import (
	"container/heap"
	"sort"
)

type meeting struct {
	roomNo  int
	endTime int64
}

type meetingHeap []meeting

func (h *meetingHeap) Len() int { return len(*h) }

func (h *meetingHeap) Less(i, j int) bool { return (*h)[i].endTime < (*h)[j].endTime }

func (h *meetingHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *meetingHeap) Push(x any) {
	*h = append(*h, x.(meeting))
}

func (h *meetingHeap) Pop() any {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

type roomsHeap []int

func (h *roomsHeap) Len() int { return len(*h) }

func (h *roomsHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }

func (h *roomsHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *roomsHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *roomsHeap) Pop() any {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func MostBooked(n int, meetings [][]int) int {
	if n == 1 {
		return 0
	}

	availableRooms := &roomsHeap{}
	for i := range n {
		*availableRooms = append(*availableRooms, i)
	}
	heap.Init(availableRooms)

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	ongoingMeetings := &meetingHeap{}
	bookingCount := make([]int, n)

	for _, meet := range meetings {
		for ongoingMeetings.Len() > 0 && (*ongoingMeetings)[0].endTime <= int64(meet[0]) {
			finishedMeeting := heap.Pop(ongoingMeetings).(meeting)
			heap.Push(availableRooms, finishedMeeting.roomNo)
		}

		if availableRooms.Len() > 0 {
			roomNo := heap.Pop(availableRooms).(int)
			heap.Push(ongoingMeetings, meeting{
				roomNo:  roomNo,
				endTime: int64(meet[1]),
			})

			bookingCount[roomNo] += 1
			continue
		}

		ongoingMeeting := heap.Pop(ongoingMeetings).(meeting)
		heap.Push(ongoingMeetings, meeting{
			roomNo:  ongoingMeeting.roomNo,
			endTime: ongoingMeeting.endTime + int64(meet[1]-meet[0]),
		})
		bookingCount[ongoingMeeting.roomNo] += 1
	}

	maxCnt := 0
	roomNo := 0

	for i, v := range bookingCount {
		if v > maxCnt {
			maxCnt = v
			roomNo = i
		}
	}

	return roomNo
}
