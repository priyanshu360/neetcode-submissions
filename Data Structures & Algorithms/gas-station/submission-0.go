func canCompleteCircuit(gas []int, cost []int) int {
   	// gas = [1,2,3,4], cost = [2,2,4,1] 
	// [1, -2, 2, -2, 3, -4, 4, -1]
	// [1, -1, 1, -1, 2, -2, 2, 1]
	// we have to shift in such way we never get zero

	// [-1, 0, -1, 3]

	si := 0
	tmp := 0
	for i, g := range gas {
		tmp += g - cost[i]
		if tmp < 0 {
			si = i + 1
			tmp = 0
		}
	}

	fmt.Println(si)

	si = si % len(gas)

	tmp = 0
	for i := si; i < len(gas); i++ {
		tmp += gas[i] - cost[i]
		if tmp < 0 {
			return -1
		}
	}

	for i := 0; i < si; i++ {
		tmp += gas[i] - cost[i]
		if tmp < 0 {
			return -1
		}
	}

	return si
}
