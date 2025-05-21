package main

func SortIntervals(intervals [][2]int) [][2]int {
	l := len(intervals)
	pivot := 0

	for {
		if pivot >= l {
			break
		}

		for i, interval := range intervals {
			if intervals[pivot][0] <= interval[0] {
				temp := append([][2]int{}, intervals[0:i]...)
				temp = append(temp, intervals[pivot])
				temp = append(temp, intervals[i:pivot]...)
				temp = append(temp, intervals[pivot+1:]...)
				intervals = temp
				break
			}
		}
		pivot++
	}

	return intervals
}

func SumOfIntervals(intervals [][2]int) int {
	sortedIntervals := SortIntervals(intervals)

	sum := 0
	start := sortedIntervals[0][0]
	end := sortedIntervals[0][1]

	for _, interval := range sortedIntervals {
		if end < interval[0] {
			sum = sum + (end - start)
			start = interval[0]
			end = interval[1]
		} else if interval[1] > end {
			end = interval[1]
		}
	}

	sum = sum + (end - start)

	return sum
}
