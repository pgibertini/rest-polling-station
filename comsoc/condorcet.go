package comsoc

func CondorcetWinner(p Profile) (alt []Alternative) {
	err := checkProfile(p)

	if err != nil {
		return
	}

	lenPrefs := len(p)
	tabCount := make([][]int, lenPrefs)
	for i := 0; i < lenPrefs; i++ {
		tabCount[i] = make([]int, lenPrefs)
	}

	//Pour chaque candidat, regarde lequel il préfère
	for _, voter := range p {
		for _, candidate := range voter {
			for _, opposition := range voter {
				pref, errPref := isPref(candidate, opposition, voter)
				if errPref != nil {
					return
				}
				if pref {
					tabCount[candidate-1][opposition-1]++
				}
			}
		}
	}

	var isWinner bool

	//Si un candidat est préféré à tous les autres, il est gagnant de Condorcet
	for i, vote := range tabCount {
		isWinner = true
		for j, against := range vote {
			if against < lenPrefs/2+1 && i != j {
				isWinner = false
			}
		}
		if isWinner {
			alt = append(alt, Alternative(i+1))
			return
		}
	}

	return
}
