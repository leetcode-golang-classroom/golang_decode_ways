package sol

func numDecodings(s string) int {
	digitMap := map[byte]struct{}{
		'0': {},
		'1': {},
		'2': {},
		'3': {},
		'4': {},
		'5': {},
		'6': {},
	}
	sLen := len(s)
	prevTwo := 1
	prevOne := 0
	if s[sLen-1] != '0' {
		prevOne = prevTwo
	}
	ways := prevOne
	for start := sLen - 2; start >= 0; start-- {
		ways = 0
		if s[start] != '0' {
			ways += prevOne
		}
		if start+1 < sLen {
			if _, ok := digitMap[s[start+1]]; s[start] == '1' || (s[start] == '2' && ok) {
				ways += prevTwo
			}
		}
		prevTwo = prevOne
		prevOne = ways
	}
	return ways
}
