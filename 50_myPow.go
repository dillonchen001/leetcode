package main

func myPow(x float64, n int) float64 {
	INT_MIN := -1 << 31

	if n == INT_MIN {
		if x > 0 {
			return 1 / myPow(x, -1-INT_MIN)
		} else {
			return -1 / myPow(x, -1-INT_MIN)
		}
	} else if n < 0 {
		return 1 / myPow(x, -n)
	} else if n == 0 {
		return 1
	}

	ret := 1.0

	for n > 0 {
		if n&1 > 0 {
			ret *= x
		}

		x *= x
		n >>= 1
	}
	return ret
}
