package comsoc

func ApprovalSWF(p Profile, thresholds []int) (count Count, err error) {
	err = checkProfile(p)

	if err != nil {
		return
	}

	count = make(Count)

	// initialize
	for _, candidate := range p[0] {
		count[candidate] = 0
	}

	for i, voter := range p {
		for j, vote := range voter {
			if j > (thresholds[i] - 1) {
				break
			}
			count[vote] += 1
		}
	}

	return
}

func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	var count Count
	count, err = ApprovalSWF(p, thresholds)

	if err != nil {
		return
	}

	bestAlts = maxCount(count)

	return
}
