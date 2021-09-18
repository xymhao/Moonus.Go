package sort

func bubbleSort(a []int) {
	if len(a) <= 1 {
		return
	}

	for i := 0; i < len(a); i++ {
		flag := false
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
				flag = true
			}
		}

		if !flag {
			return
		}
	}
}
