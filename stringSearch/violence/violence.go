package violence


func SearchOne(P , T string) int {
	var (
		i = 0
		j = 0
	)
	for i < len(T) && j < len(P) {
		if T[i] == P[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j == len(P) {
		return i - j
	} else {
		return -1
	}
}

func SearchAll(P, T string) []int {
	var (
		i = 0
		j = 0
	)

	var match_list []int

	for i < len(T) {
		if T[i] == P[j] {
			i ++
			j ++
		} else {
			i = i - j + 1
			j = 0
		}
		if j == len(P) {
			match_list = append(match_list, i - j)
			j = 0
		}
	}
	return match_list
}
