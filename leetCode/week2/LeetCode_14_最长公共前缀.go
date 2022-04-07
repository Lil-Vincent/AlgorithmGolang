package week2

func LongestCommonPrefix(strs []string) string {
	res := ""
	if len(strs) == 0 {
		return res
	}

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for _, str := range strs {
			if i > len(str)-1 {
				return res
			}
			if str[i] != c {
				return res
			}
		}
		res += string(c)
	}
	return res
}
