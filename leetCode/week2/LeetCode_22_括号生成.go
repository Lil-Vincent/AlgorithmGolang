package week2

var res = []string{}

func generateParenthesis(n int) []string {
	res = []string{}
	dfs(n, 0, 0, "")
	return res
}

func dfs(n int, lc int, rc int, path string) {

	if lc == n && rc == n {
		res = append(res, path)
	} else {
		if lc < n {
			dfs(n, lc+1, rc, path+"(")
		}
		if rc < n && rc < lc {
			dfs(n, lc, rc+1, path+")")
		}
	}
}
