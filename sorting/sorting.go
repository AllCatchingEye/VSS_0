package sorting

func Quicksort(s []int) {
	if len(s) <= 1 {
		return
	}

	v := s[len(s)-1]
	i, j := 0, len(s)-1

	for {
		for s[i] <= v && i < len(s)-1 {
			i++
		}

		for s[j] > v && j > 0 {
			j--
		}

		if i >= j {
			break
		}

		s[i], s[j] = s[j], s[i]
	}
	//s[i], s[j] = s[len(s)-1], s[i]

	Quicksort(s[:i])
	Quicksort(s[i:])
	return
}

func Mergesort(s []int) []int {
	if len(s) <= 1 {
		return s
	} else {
		m := len(s) / 2
		lhs, rhs := s[:m], s[m:]

		lhs = Mergesort(lhs)
		rhs = Mergesort(rhs)

		var result []int
		i, j := 0, 0
		for i < len(lhs) && j < len(rhs) {
			l, r := lhs[i], rhs[j]
			if l < r {
				result = append(result, l)
				i++
			} else {
				result = append(result, r)
				j++
			}
		}

		if i < len(lhs) {
			result = append(result, lhs[i:]...)
		} else {
			result = append(result, rhs[j:]...)
		}

		return result
	}
}
