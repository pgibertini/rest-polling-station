package main

import (
	"fmt"
	"github.com/pgibertini/rest-polling-station/comsoc"
)

func main() {
	//comsoc.TestUtils()

	// testing vote methods
	prefs := [][]comsoc.Alternative{
		{1, 2, 3},
		{3, 2, 1},
		{3, 2, 1},
	}

	println("Majority")
	res0, _ := comsoc.MajoritySWF(prefs)
	println(res0[1], res0[2], res0[3])

	res1, _ := comsoc.MajoritySCF(prefs)
	println(res1[0])

	println("Borda")
	res2, _ := comsoc.BordaSWF(prefs)
	println(res2[1], res2[2], res2[3])

	res3, _ := comsoc.BordaSCF(prefs)
	println(res3[0])

	thresholds := []int{2, 1, 2}

	println("Approval")
	res4, _ := comsoc.ApprovalSWF(prefs, thresholds)
	println(res4[1], res4[2], res4[3])
	println("Ranking: ")
	ranking := comsoc.Ranking(res4)
	for _, k := range ranking {
		print(k)
	}

	println()
	println()

	prefs = [][]comsoc.Alternative{
		{1, 3, 2},
		{1, 2, 3},
		{2, 1, 3},
	}

	res5, _ := comsoc.ApprovalSCF(prefs, thresholds)
	println(res5[0])
	println(len(res5))

	// tie break
	fmt.Println("\nTIE BREAK")

	order := []comsoc.Alternative{4, 3, 1, 2}
	alts := []comsoc.Alternative{2, 4}

	tb := comsoc.TieBreakFactory(order)

	best, _ := tb(alts)

	fmt.Println(best)

}
