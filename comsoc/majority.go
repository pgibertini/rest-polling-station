package comsoc

func MajoritySWF(p Profile) (count Count, err error) {
	err = checkProfile(p)

	if err != nil {
		return
	}

	count = make(Count)

	// initialize
	for _, candidate := range p[0] {
		count[candidate] = 0
	}

	for _, voter := range p {
		count[voter[0]] += 1
	}

	return
}

func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	var count Count
	count, err = MajoritySWF(p)

	if err != nil {
		return
	}

	bestAlts = maxCount(count)

	return
}
