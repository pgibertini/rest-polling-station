package comsoc

func SWF(p Profile) (count Count, err error) {
	err = checkProfile(p)

	if err != nil {
		return
	}

	for _, voter := range p {
		count[voter[0]] += 1
	}
	return
}

func SCF(p Profile) (bestAlts []Alternative, err error) {
	var count Count
	count, err = SWF(p)

	if err != nil {
		return
	}

	bestAlts = maxCount(count)

	return
}
