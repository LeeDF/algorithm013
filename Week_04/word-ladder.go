package Week_04

import (
	"fmt"
)

//"hit"
//"cog"
//["hot","dot","dog","lot","log","cog"]
func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

//bfs 每次找到下一个能到的单词，将下次层放入q中，每次取完q
func ladderLength(beginWord string, endWord string, wordList []string) int {
	mat := map[string][]string{}
	for _, i := range wordList {
		for j := 0; j < len(i); j++ {
			k := i[:j] + "*" + i[j+1:]
			if _, ok := mat[k]; !ok {
				mat[k] = []string{}
			}
			mat[k] = append(mat[k], i)
		}
	}
	q := []string{beginWord,}
	visited := map[string]bool{}
	step := 1

	for len(q) > 0 {
		nextQ := []string{}
		for _, i := range q {
			for j := 0; j < len(i); j++ {
				k := i[:j] + "*" + i[j+1:]
				if mk, ok := mat[k]; ok {
					for _,mk1 := range mk{
						if mk1 == endWord{
							return step+1
						}
						if _, ok := visited[mk1];ok{
							continue
						}
						nextQ = append(nextQ, mk1)
						visited[mk1] = true
					}
				}
			}
		}
		step ++
		q = nextQ
	}
	return 0
}
