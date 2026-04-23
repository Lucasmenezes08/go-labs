package longestcommonprefix

import "math"

func longestCommonPrefix(strs []string) string {
	min := math.MaxInt

	for _, str := range strs {
		if len(str) < min {
			min = len(str)
		}
	}

	var result string

	for i := 0 ; i < min; i++ {
		count := 0
		for j := 1; j < len(strs); j ++{
			if strs[0][i] == strs[j][i]{
				count ++
			}
		}

		if count == len(strs) - 1{
			result += string(strs[0][i])
		} else {
			break
		}
	}
	return result
}