func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	charcount := map[rune]int{}

	for _, char := range s {
		charcount[char]++
	}

	for _, char := range t {
		charcount[char]--
		
		// se for -1 tem letra a mais, se for 0 continua, se for 1 continua, ex: voou
		if charcount[char] < 0 {
			return false
		}
	}

	return true
}
