package Week_04

import "fmt"

//贪心, 每次从最大的金币开始找零
func lemonadeChange(bills []int) bool {
	wallet := map[int]int{
		5:0,
		10:0,
		//20:0,
	}
	l := []int{20,10,5}
	for _,i := range bills{
		//找零,从最大开始找
		p := i-5
		if p > 0 {
			for _,j := range l{
				if wallet[j] == 0{
					continue
				}
				if p / j <= wallet[j]{
					wallet[j] -= p / j
					p -= (p / j) * j
				}else{
					wallet[j] = 0
					p -= wallet[j] * j
				}
			}
			fmt.Println(p)
			if p > 0{
				return false
			}
		}

		wallet[i] ++
	}
	return true
}

func main()  {
	fmt.Println(lemonadeChange([]int{5,5,10,10,20}))
}