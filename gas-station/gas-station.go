package main

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	end := 1
	tank := gas[0] - cost[0]
	n := len(gas)
	for traverse := 1; traverse < n; traverse++ {
		// if the car can not go right, then change the position of start
		// until the car has gone n-1 gas stations
		if tank >= 0 && tank+gas[end]-cost[end] >= 0 {
			tank = tank + gas[end] - cost[end]
			end = (end + 1) % n
		} else {
			start = (start + n - 1) % n
			tank = tank + gas[start] - cost[start]
		}
	}
	if tank < 0 {
		return -1
	}
	return start
}
