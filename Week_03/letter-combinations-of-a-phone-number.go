package Week_03

import "fmt"

func letterCombinations(digits string) []string {
	m := map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	var res []string
	_letter([]byte(digits), m, "", &res)
	return res
}

func _letter(digits []byte, m map[byte][]string, ret string, res *[]string) {
	if len(digits) == 0{
		return
	}
	if len(ret) == len(digits) {
		*res = append(*res, ret)
		return
	}
	for _, i := range m[digits[len(ret)]] {
		_letter(digits, m, ret+i, res)
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
