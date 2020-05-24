package main

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	ln1 := len(num1)
	ln2 := len(num2)

	n1 := 0
	n2 := 0
	sum := 0
	res := make(map[int]int)

	for i := ln1 - 1; i >= 0; i-- {
		n1, _ = strconv.Atoi(string(num1[i]))
		for j := ln2 - 1; j >= 0; j-- {
			n2, _ = strconv.Atoi(string(num2[j]))
			sum = (res[i+j+1] + n1*n2)
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}

	var buffer bytes.Buffer
	if res[0] > 0 {
		buffer.WriteString(strconv.Itoa(res[0]))
	}

	for k := 1; k < ln1+ln2; k++ {
		buffer.WriteString(strconv.Itoa(res[k]))
	}
	return buffer.String()
}
