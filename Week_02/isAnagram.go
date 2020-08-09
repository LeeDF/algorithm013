package Week_02

// 1.排序, 看是否相等 nlogn
// 2.map 存每个字母次数

func isAnagram(s string, t string) bool {
	m := map[int32]int32{}
	if len(s) == 0 && len(t) == 0 {
		return false
	}
	for _, i := range s {
		if _, ok := m[i]; !ok {
			m[i] = 0
		}
		m[i]++
	}
	for _, i := range s {
		if _, ok := m[i]; !ok {
			return false
		}
		m[i]--
		if m[i] == 0 {
			delete(m, i)
		}
	}
	if len(m) == 0 {
		return true
	}
	return false
}
