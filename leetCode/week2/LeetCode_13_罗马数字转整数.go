package week2

func romanToInt(s string) int {

	myMap := map[byte]int{
		'I': 1, 'V': 5,
		'X': 10, 'L': 50,
		'C': 100, 'D': 500,
		'M': 1000,
	}

	res := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && myMap[s[i]] < myMap[s[i+1]] {
			res -= myMap[s[i]]
		} else {
			res += myMap[s[i]]
		}
	}

	return res
}
