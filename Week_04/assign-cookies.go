package Week_04

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	num := 0
	st := 0
	for _, i := range g {
		for st < len(s){
			if i <= s[st]{
				num++
				st++
				break
			}
			st++
		}
	}
	return num
}


