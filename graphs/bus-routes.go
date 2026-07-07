package graphs

// Question:
// LC 815. Bus Routes
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	numBuses := len(routes)

	stopToBus := make(map[int][]int)
	for i := range routes {
		for _, stop := range routes[i] {
			stopToBus[stop] = append(stopToBus[stop], i)
		}
	}

	var busQueue []int
	busTaken := make([]bool, numBuses)
	for _, bus := range stopToBus[source] {
		busQueue = append(busQueue, bus)
		busTaken[bus] = true
	}

	level := 1
	for len(busQueue) > 0 {
		size := len(busQueue)

		for i := range size {
			currBus := busQueue[i]

			stops := routes[currBus]
			for _, stop := range stops {
				if stop == target {
					return level
				}

				buses := stopToBus[stop]
				for _, bus := range buses {
					if !busTaken[bus] {
						busQueue = append(busQueue, bus)
						busTaken[bus] = true
					}
				}
			}
		}

		busQueue = busQueue[size:]
		level++
	}

	return -1
}
