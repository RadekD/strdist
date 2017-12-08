package strdist

//Levenshtein calculates Levenshtein distance of two strings
// more info https://en.wikipedia.org/wiki/Levenshtein_distance
func Levenshtein(source, target string) int {
	if source == target {
		return 0
	}

	lTarget := len(target) + 1

	v0 := make([]int, lTarget)
	v1 := make([]int, lTarget)

	for i := 0; i < lTarget; i++ {
		v0[i] = i
	}

	for i, s := range source {
		v1[0] = i + 1

		for j, t := range target {
			nv := min(v0[j+1]+1, v1[j]+1)

			if s == t {
				nv = min(nv, v0[j])
			} else {
				nv = min(nv, v0[j]+1)
			}

			v1[j+1] = nv
		}

		v0, v1 = v1, v0
	}
	return v0[lTarget-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
