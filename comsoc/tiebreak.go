package comsoc

import "errors"

// MinTieBreak renvoie l'alternative d'indice le plus petit
func MinTieBreak(alts []Alternative) (alt Alternative, err error) {
	if len(alts) == 0 {
		err = errors.New("given alts is an empty slice")
		return
	}

	alt = alts[0]
	for i := range alts {
		if alt > alts[i] {
			alt = alts[i]
		}
	}

	return
}

func TieBreakFactory(order []Alternative) func([]Alternative) (alt Alternative, err error) {
	tieBreak := func(alts []Alternative) (alt Alternative, err error) {
		if len(alts) == 0 {
			err = errors.New("given alts is an empty slice")
			return
		}

		for i := range order {
			for j := range alts {
				if order[i] == alts[j] {
					alt = alts[j]
					return
				}
			}
		}
		err = errors.New("tie break function alternatives does not match given alternatives")
		return
	}

	return tieBreak
}
