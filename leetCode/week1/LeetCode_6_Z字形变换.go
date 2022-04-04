package week1

func Convert(s string, n int) string {
	if n == 1 {
		return s
	}

	res := ""
	for i := 0; i < n; i++ {
		if i == 0 || i == n-1 {
			for j := i; j < len(s); j += 2*n - 2 {
				res += string(s[j])
			}
		} else {
			for j, k := i, 2*n-2-i; j < len(s) || k < len(s); j, k = j+2*n-2, k+2*n-2 {
				if j < len(s) { //这里要注意判定
					res += string(s[j])
				}

				if k < len(s) {
					res += string(s[k])
				}
			}
		}
	}

	return res

}
