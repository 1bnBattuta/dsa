package array

// Time complexity: O(n), space O(n)
func Rearrange_by_sign(arr []int) {
	n := len(arr)
	pos := make([]int, 0, n)
	neg := make([]int, 0, n)

	for _, i := range arr {
		if i < 0 {
			neg = append(neg, i)
		} else {
			pos = append(pos, i)
		}
	}

	npos := len(pos)
	nneg := len(neg)

	pindx := 0
	nindx := 0
	i := 0

	for pindx < npos && nindx < nneg {
		if i%2 == 0 {
			arr[i] = pos[pindx]
			i++
			pindx++
		} else {
			arr[i] = neg[nindx]
			i++
			nindx++
		}
	}

	for nindx < nneg {
		arr[i] = neg[nindx]
		i++
		nindx++
	}

	for pindx < npos {
		arr[i] = pos[pindx]
		i++
		pindx++
	}
}
