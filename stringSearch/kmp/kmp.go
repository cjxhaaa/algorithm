package kmp

func GetNext(P string) []int {
	var (
		j = 0
		k = -1
	)
	next := make([]int, len(P))
	next[j] = -1
	for j < len(P) - 1 {
		if k == -1 || P[j] == P[k] {
			j += 1
			k += 1
			next[j] = k
		} else {
			k = next[k]
		}
	}
	return next
}

func KmpSearchOne(P, T string) int {
	var (
		i = 0
		j = 0
	)
	next := GetNext(P)
	//fmt.Println(next)
	for i < len(T) && j < len(P) {
		//fmt.Println("i:",i,"j:",j)

		if j == -1 || T[i] == P[j] {
			i += 1
			j += 1
		} else {
			j = next[j]
		}
	}
	if j == len(P) {
		return i - j
	} else {
		return -1
	}
}
