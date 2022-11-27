package comsoc

func BordaSWF(p Profile) (count Count, err error) {
	err = checkProfile(p)

	if err != nil {
		return
	}

	count = make(Count)

	// initialize
	for _, candidate := range p[0] {
		count[candidate] = 0
	}

	var lenVoter int

	for _, voter := range p {
		lenVoter = len(voter)
		for i, vote := range voter {
			count[vote] += lenVoter - (i + 1)
		}
	}

	return
}
func BordaSCF(p Profile) (bestAlts []Alternative, err error) {
	var count Count
	count, err = BordaSWF(p)

	if err != nil {
		return
	}

	bestAlts = maxCount(count)

	return
}
