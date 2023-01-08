package gasstation

// brute force O(n^2)
// func canCompleteCircuit(gas []int, cost []int) int {
// 	l := len(gas)
// 	i := 0
// 	for i < l {
// 		if gas[i] >= cost[i] && travelCircular(i, gas, cost) {
// 			return i
// 		}
// 		i++
// 	}
// 	return -1
// }

// func travelCircular(start int, gas []int, cost []int) bool {
// 	flag := false
// 	curr := start
// 	tank := gas[curr]

// 	for {
// 		next := curr + 1

// 		if curr == len(gas)-1 {
// 			next = 0
// 			flag = true
// 		}

// 		tank = tank - cost[curr] + gas[next]
// 		if tank < cost[next] {
// 			return false
// 		}

// 		curr++
// 		if next == 0 {
// 			curr = 0
// 		}
// 		if flag && curr == start {
// 			break
// 		}
// 	}
// 	return true
// }

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	tmpTank, globalTank, next := 0, 0, 0

	for i := 0; i < n; i++ {
		tmpTank += gas[i] - cost[i]
		globalTank += gas[i] - cost[i]
		if tmpTank < 0 {
			next = i + 1
			tmpTank = 0
		}
	}

	if globalTank < 0 {
		return -1
	}
	return next
}
