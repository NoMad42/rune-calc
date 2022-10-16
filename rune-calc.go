package main

import (
	"fmt"
	"sort"
)

func RuneCalc(in string) (out string) {
	runeMap := map[rune]int{}

	for _, r := range in {
		runeCount, runeInMapExist := runeMap[r]
		if runeInMapExist {
			runeMap[r] = runeCount + 1
		} else {
			runeMap[r] = 1
		}
	}

	keys := make([]rune, 0, len(runeMap))
	for k := range runeMap {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, k := range keys {
		out += fmt.Sprintf("%s%d", string(k), runeMap[k])
	}

	return out
}

func main() {
	fmt.Println("aaaaaaaabbbccccc", "=>", RuneCalc("aaaaaaaabbbccccc"))
	fmt.Println("zzzzcccUUUuuu", "=>", RuneCalc("zzzzcccUUUuuu"))
	fmt.Println("ЯЯЯБББддд", "=>", RuneCalc("ЯЯЯБББддд"))
	fmt.Println("aaabbbccccc", "=>", RuneCalc("aaabbbccccc"))
}
