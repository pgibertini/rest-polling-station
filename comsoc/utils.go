package comsoc

import (
	"errors"
	"fmt"
)

type Alternative int
type Profile [][]Alternative
type Count map[Alternative]int

// Renvoie l'indice ou se trouve alt dans prefs
func rank(alt Alternative, prefs []Alternative) (int, error) {
	for i, v := range prefs {
		if v == alt {
			return i, nil
		}
	}
	return 0, errors.New("alt not found in given prefs")
}

// Renvoie vrai ssi alt1 est préférée à alt2
func isPref(alt1, alt2 Alternative, prefs []Alternative) (bool, error) {
	rank1, err := rank(alt1, prefs)
	if err != nil {
		return true, err
	}

	rank2, err := rank(alt2, prefs)
	if err != nil {
		return true, err
	}

	if rank1 < rank2 {
		return true, nil
	}
	return false, nil
}

// Renvoie les meilleures alternatives pour un décompte donné
func maxCount(count Count) (bestAlts []Alternative) {
	max := 0
	for alt, votes := range count {
		if votes == max {
			bestAlts = append(bestAlts, alt)
		} else if votes > max {
			bestAlts = nil
			max = votes
			bestAlts = append(bestAlts, alt)
		}
	}
	return
}

func checkProfileCheckDup(voteur []Alternative) error {
	var j, k int
	var alt Alternative

	for j, alt = range voteur {
		//Compare aux autres valeurs du voteur
		for k = j + 1; k < len(voteur); k++ {
			if alt == voteur[k] {
				return errors.New("same alt found twice in a profile")
			}
		}
	}
	return nil
}

// Vérifie que le profil donné n'a pas de valeur en double dans un voteur
func checkProfile(prefs Profile) error {
	var dupVal error

	//Regarde chaque voteur
	for _, voteur := range prefs {
		//Regarde chaque valeur du voteur
		dupVal = checkProfileCheckDup(voteur)
		if dupVal != nil {
			return dupVal
		}
	}
	return nil
}

func checkProfileAlternativeCompareAlt(voteur []Alternative, alts []Alternative) error {
	var isInAlts bool
	for _, altvoteur := range voteur {
		//Compare les valeurs du voteur aux alternatives
		isInAlts = false
		for _, alt := range alts {
			if altvoteur == alt {
				isInAlts = true
			}
		}
		if !isInAlts {
			return errors.New("unknown alt found in a profile")
		}
	}
	return nil
}

// Vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative de alts apparaît exactement une fois par préférences
func checkProfileAlternative(prefs Profile, alts []Alternative) error {
	//Vérifie que chaque ne contient pas de valeur en double dans un voteur
	errDupe := checkProfile(prefs)
	if errDupe != nil {
		return errDupe
	}

	var valErr error
	//Regarde chaque voteur
	for _, voteur := range prefs {
		if len(voteur) != len(alts) {
			return errors.New("profile with missing or too many values found")
		}
		//Regarde chaque valeur du voteur
		valErr = checkProfileAlternativeCompareAlt(voteur, alts)
		if valErr != nil {
			return valErr
		}
	}
	return nil
}

func TestUtils() {
	prefs := [][]Alternative{
		{1, 2, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	alts := []Alternative{1, 2, 3}

	count := map[Alternative]int{1: 20, 2: 20, 3: 21}

	fmt.Println(rank(4, prefs[0]))

	fmt.Println(isPref(1, 2, prefs[2]))

	fmt.Println(maxCount(count))

	fmt.Println(checkProfile(prefs))

	fmt.Println(checkProfileAlternative(prefs, alts))
}
