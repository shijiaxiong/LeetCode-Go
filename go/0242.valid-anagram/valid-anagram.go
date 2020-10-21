package problem0242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
 
	sr := []rune(s)
	tr := []rune(t)

	rec := make(map[rune]int, len(sr))

	for i := range sr {
		rec[sr[i]]--
		rec[tr[i]]++
	}

	for _, v := range rec {
		if v != 0 {
			return false
		}
	}

	return true
}

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	rec := make([]int,26)
	sb := []byte(s)
	tb := []byte(t)

	for i := range sb {
		rec[sb[i]]--
		rec[tb[i]]++
	}

	for _, v := range rec {
		if v == 0 {
			return false
		}
	}

	return true
}